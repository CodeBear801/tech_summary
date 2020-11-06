
- [Example](#example)
- [db](#db)
- [Transaction](#transaction)
- [Bucket](#bucket)
	- [Cursor](#cursor)
- [Put](#put)
- [commit](#commit)
- [More info](#more-info)

## Example

Start by an example of how to use boltdb to write

```go
db, err := bolt.Open("test.db", 0600, nil)
if err != nil {
    log.Fatal(err)
}
defer db.Close()

// Start a writable transaction.
tx, err := db.Begin(true)
if err != nil {
    return err
}
defer tx.Rollback()

// Use the transaction...
bucket, err := tx.CreateBucket([]byte("testBucket"))
if err != nil {
    return err
}
bucket.Put([]byte("testKey"), []byte("testValue"))

// Commit the transaction and check for error.
if err := tx.Commit(); err != nil {
    return err
}
```

## db

The `test.db` is recorded on a **single** persistent file on disk, [`DB`](https://github.com/boltdb/bolt/blob/fd01fc79c553a8e99d512a07e8e0c63d4a3ccfc5/db.go#L45) is the abstraction for this.  As its comments said, `DB` represents a collection of buckets.  

When calling [`db.Open()`](https://github.com/boltdb/bolt/blob/fd01fc79c553a8e99d512a07e8e0c63d4a3ccfc5/db.go#L150)

```go
// Open creates and opens a database at the given path.
// If the file does not exist then it will be created automatically.
// Passing in nil options will cause Bolt to open the database with the default options.
func Open(path string, mode os.FileMode, options *Options) (*DB, error) {

    // Open data file and separate sync handler for metadata writes.
    if db.file, err = os.OpenFile(db.path, flag|os.O_CREATE, mode); err != nil {
        // ...

    // Lock file so that other processes using Bolt in read-write mode cannot
	// use the database  at the same time. This would cause corruption since
	// the two processes would write meta pages and free pages separately.
	// The database file is locked exclusively (only one process can grab the lock)
	// if !options.ReadOnly.
	// The database file is locked using the shared lock (more than one process may
    // hold a lock at the same time) otherwise (options.ReadOnly is set).
    if err := flock(db, mode, !db.readOnly, options.Timeout); err != nil {
        // ...

    // load page size
    // ...

    // Initialize page pool.
    db.pagePool = sync.Pool{
        // ...

    // Memory map the data file.
    if err := db.mmap(options.InitialMmapSize); err != nil {
        // ... 

    // Read in the freelist.
	db.freelist = newFreelist()
    db.freelist.read(db.page(db.meta().freelist))

    return

```
The implementation of `db.mmap` is platform dependently, take `bolt_unix` as an example:
```go
// mmap memory maps a DB's data file.
func mmap(db *DB, sz int) error {
	// Map the data file to memory.
	b, err := syscall.Mmap(int(db.file.Fd()), 0, sz, syscall.PROT_READ, syscall.MAP_SHARED|db.MmapFlags)
	if err != nil {
		return err
	}

	// Advise the kernel that the mmap is accessed randomly.
	if err := madvise(b, syscall.MADV_RANDOM); err != nil {
		return fmt.Errorf("madvise: %s", err)
	}

	// Save the original byte slice and convert to a byte array pointer.
	db.dataref = b
	db.data = (*[maxMapSize]byte)(unsafe.Pointer(&b[0]))
	db.datasz = sz
	return nil
}
// [Perry] there is no page cache in boltdb, who depend on the OS to manage pages
// all operations of read/write on file will be apply to db.data directly 

```

When create a db file, will write 4 pages by default:
- 2 meta page, which records where is freelist page and data page
- 1 freelist page
- 1 data page


<img src="https://user-images.githubusercontent.com/16873751/98424592-f844f380-2046-11eb-99f4-021003288f93.png" alt="botdb_bucket" width="400"/>
<br/>

```go
	db.meta0 = db.page(0).meta()
    db.meta1 = db.page(1).meta()
    
    m.freelist = 2
    m.root = bucket{root: 3}
    
// [Perry] the real data inside db is a b+ tree and represented as bucket
```

More info:
- mmap
- flock

## Transaction

[`db.Begin(true)`](https://github.com/boltdb/bolt/blob/fd01fc79c553a8e99d512a07e8e0c63d4a3ccfc5/db.go#L442) creates a writeable transaction
```go
// Begin starts a new transaction.
// Multiple read-only transactions can be used concurrently but only one
// write transaction can be used at a time. Starting multiple write transactions
// will cause the calls to block and be serialized until the current write
// transaction finishes.
//
// Transactions should not be dependent on one another. Opening a read
// transaction and a write transaction in the same goroutine can cause the
// writer to deadlock because the database periodically needs to re-mmap itself
// as it grows and it cannot do that while a read transaction is open.
//
// If a long running read transaction (for example, a snapshot transaction) is
// needed, you might want to set DB.InitialMmapSize to a large enough value
// to avoid potential blocking of write transaction.
//
// IMPORTANT: You must close read-only transactions after you are finished or
// else the database will not reclaim old pages.
func (db *DB) Begin(writable bool) (*Tx, error) {
	if writable {
		return db.beginRWTx()
	}
	return db.beginTx()
}
```
`beginRWTx()` will create a [Tx](https://github.com/boltdb/bolt/blob/fd01fc79c553a8e99d512a07e8e0c63d4a3ccfc5/tx.go#L24) which manages read-only or read-write operation on db


## Bucket
`tx.CreateBucket()` try to create a [`Bucket`](https://github.com/boltdb/bolt/blob/fd01fc79c553a8e99d512a07e8e0c63d4a3ccfc5/bucket.go#L36) object, you can think which as logic operator on DB's B+ tree, and its a `namespace` which is similar as `table` in database

```go
// Bucket represents a collection of key/value pairs inside the database.
type Bucket struct {
	*bucket
	tx       *Tx                // the associated transaction
    buckets  map[string]*Bucket // subbucket cache
    // [Perry] during transactions' exploration, cache all buckets visited

	page     *page              // inline page reference
    rootNode *node              // materialized node for the root page.
    // [Perry] this is the key 
    // rootNode is the one to index all KV inside bucket

    nodes    map[pgid]*node     // node cache
    // [Perry] during transactions' exploration, cache all nodes visited

	// Sets the threshold for filling nodes when they split. By default,
	// the bucket will fill to 50% but it can be useful to increase this
	// amount if you know that your write workloads are mostly append-only.
	//
	// This is non-persisted across transactions so it must be set in every Tx.
	FillPercent float64
}

// bucket represents the on-file representation of a bucket.
// This is stored as the "value" of a bucket key. If the bucket is small enough,
// then its root page can be stored inline in the "value", after the bucket
// header. In the case of inline buckets, the "root" will be 0.
type bucket struct {
	root     pgid   // page id of the bucket's root-level page
	sequence uint64 // monotonically incrementing, used by NextSequence()
}
```

<img src="https://user-images.githubusercontent.com/16873751/98407817-244f7d00-2025-11eb-800f-ad0356149725.png" alt="botdb_bucket" width="400"/>
<br/>

The implementation of bucket CreateBucket is [here](https://github.com/boltdb/bolt/blob/fd01fc79c553a8e99d512a07e8e0c63d4a3ccfc5/bucket.go#L158)

```go
// CreateBucket creates a new bucket at the given key and returns the new bucket.Ben Johnson, 7 years ago: • Return bucket from CreateBucket() functions.
// Returns an error if the key already exists, if the bucket name is blank, or if the bucket name is too long.
// The bucket instance is only valid for the lifetime of the transaction.
func (b *Bucket) CreateBucket(key []byte) (*Bucket, error) {

	// Move cursor to correct position.
	c := b.Cursor()
	k, _, flags := c.seek(key)
	
	// Return an error if there is an existing key.
	if bytes.Equal(key, k) {
		if (flags & bucketLeafFlag) != 0 {
			return nil, ErrBucketExists
		}
		return nil, ErrIncompatibleValue
	}

	// Insert into node.
	key = cloneBytes(key)
	c.node().put(key, key, value, 0, bucketLeafFlag)

// [perry] CreateBucket's logic is very similar to Bucket->Put()
// Because both of them is maintaining the B+ tree
```


more info:
- [boltdb transaction](./boltdb_transaction.md)

### Cursor

```go
// Cursor creates a cursor associated with the bucket.
// The cursor is only valid as long as the transaction is open.
// Do not use a cursor after the transaction is closed.
func (b *Bucket) Cursor() *Cursor {
	// Update transaction statistics.
	b.tx.stats.CursorCount++

	// Allocate and return a cursor.
	return &Cursor{
		bucket: b,
		stack:  make([]elemRef, 0),
	}
}
type Cursor struct {
	bucket *Bucket
	stack  []elemRef
}
type elemRef struct {
	page  *page
    node  *node
    // [Perry] An elemRef is either page or node, only one of them have value
	index int
}
```


## Put
[`bucket.Put()`](https://github.com/boltdb/bolt/blob/fd01fc79c553a8e99d512a07e8e0c63d4a3ccfc5/bucket.go#L285)

```go
// Put sets the value for a key in the bucket.
// If the key exist then its previous value will be overwritten.
// Supplied value must remain valid for the life of the transaction.
// Returns an error if the bucket was created from a read-only transaction, if the key is blank, if the key is too large, or if the value is too large.
func (b *Bucket) Put(key []byte, value []byte) error {

    // Move cursor to correct position.
	c := b.Cursor()
	k, _, flags := c.seek(key)

	// Return an error if there is an existing key with a bucket value.
	if bytes.Equal(key, k) && (flags&bucketLeafFlag) != 0 {
		return ErrIncompatibleValue
	}

	// Insert into node.
	key = cloneBytes(key)
	c.node().put(key, key, value, 0, 0)

```

Calling [node's put()](https://github.com/boltdb/bolt/blob/fd01fc79c553a8e99d512a07e8e0c63d4a3ccfc5/node.go#L116)
```go
// put inserts a key/value.
func (n *node) put(oldKey, newKey, value []byte, pgid pgid, flags uint32) {
    // Find insertion index.
	index := sort.Search(len(n.inodes), func(i int) bool { return bytes.Compare(n.inodes[i].key, oldKey) != -1 })

	// Add capacity and shift nodes if we don't have an exact match and need to insert.
	exact := (len(n.inodes) > 0 && index < len(n.inodes) && bytes.Equal(n.inodes[index].key, oldKey))
	if !exact {
		n.inodes = append(n.inodes, inode{})
		copy(n.inodes[index+1:], n.inodes[index:])
	}

```



## commit

## More info
- [How BoltDB Write its Data?](https://medium.com/@abserari/how-boltdb-write-its-data-61f64a3c0e06)