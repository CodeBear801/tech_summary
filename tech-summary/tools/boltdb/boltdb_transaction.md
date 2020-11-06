# Transaction

BoltDB provide serializable transaction support.

## flock
BoltDB support multiple read transaction and max one read transaction.  
It has been controlled by file level read/write lock

```go
func flock(db *DB, mode os.FileMode, exclusive bool, timeout time.Duration) error {
	// ...
  flag := syscall.LOCK_SH   // read only
  if exclusive {
    flag = syscall.LOCK_EX  // read-write, exclusive
  }

  err := syscall.Flock(int(db.file.Fd()), flag|syscall.LOCK_NB)
  // ..
```

more info:
- https://man7.org/linux/man-pages/man2/flock.2.html
- [被遗忘的桃源——flock 文件锁](https://zhuanlan.zhihu.com/p/25134841)

## Read & Write

### Read transaction

example
```go
func main() {
  // init db
  err = db.View(func(tx *bolt.Tx) error {
    bucket := tx.Bucket([]byte("b1"))
    if bucket != nil {
      v := bucket.Get([]byte("k1"))
      fmt.Printf("%s\n", v)
    }
    return nil
  })
}
```
In the implementation of [`db.View`](https://github.com/boltdb/bolt/blob/fd01fc79c553a8e99d512a07e8e0c63d4a3ccfc5/db.go#L612), it controls init, execute and close read-only transaction

```golang
func (db *DB) View(fn func(*Tx) error) error {
  t := db.Begin(false)
  // ...
  t.managed = true
  fn(t)
  t.managed = false
  // ...
  t.Rollback()
}
```

### Read/Write transaction
```go
func main() {
  // init db
  err = db.Update(func(tx *bolt.Tx) error {
    bucket, err := tx.CreateBucketIfNotExists([]byte("b1"))
    if err != nil {
      return err
    }
    return bucket.Put([]byte("k1"), []byte("v1"))
  })
}
```

[`db.Update`](https://github.com/boltdb/bolt/blob/fd01fc79c553a8e99d512a07e8e0c63d4a3ccfc5/db.go#L581) will init, execute, commit/rollback, modifications
```go
func (db *DB) Update(fn func(*Tx) error) error {
  t := db.Begin(true)
  // ...
  t.managed = true
  err := fn(t)
  t.managed = false
  if err != nil {
    t.Rollback()
    return err
  }
  return t.Commit()
}
```


## MVCC

<img src="https://user-images.githubusercontent.com/16873751/98324274-b19cbe00-1fa0-11eb-91a9-89c4f9cd4468.png" alt="botdb_read" width="400"/>

The relation of transaction, database version and time is:
- t < t1, no change to DB, data is v1
- t1 <= t <= t2, `RW-Tx-1` is finished and persisted, data is v2
- t2 <= t <= t3, `RW-Tx-1` and `RW-Tx-2` is finished and persisted, data is v3
- t >= t3, `RW-Tx-1`, `RW-Tx-2` and `RW-Tx-3` is finished and persisted, data is v4

When a transaction start before t1 and last after t3, which version of data should return
- Two choice, return version before t1 or after t3, which is V1 or V4.  V1 is better
- If return V4, to guarantee consistency between different reads, this transaction need to on hold until t3 then excute, it might bring additional load for DB

<img src="https://user-images.githubusercontent.com/16873751/98324292-b8c3cc00-1fa0-11eb-84cc-86b466f48439.png" alt="botdb_read" width="400"/>

- Between t3 -> t4, read only transaction `RO-Tx-1` requires V1 in db
- `RO-Tx-2 ` requires V2
- `RO-Tx-3` requires V3

## More info
- https://github.com/boltdb/bolt/blob/fd01fc79c553a8e99d512a07e8e0c63d4a3ccfc5/tx.go#L144
- https://github.com/boltdb/bolt/blob/fd01fc79c553a8e99d512a07e8e0c63d4a3ccfc5/tx_test.go#L15
- https://github.com/ZhengHe-MD/learn-bolt/blob/master/TX.md
- https://github.com/boltdb/bolt/issues/392