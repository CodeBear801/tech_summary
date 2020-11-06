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


## More info
- https://github.com/ZhengHe-MD/learn-bolt/blob/master/TX.md