# BoltDB

```
BoltDB started out as a port of LMDB to Go but has somewhat diverged since then. While LMDB heavily focuses on raw performance, Bolt has focused on simplicity and ease of use. However, they both share the same design

```
(from: https://dgraph.io/blog/post/badger-lmdb-boltdb/)

```
“Bolt was originally a port of LMDB so it is architecturally similar. Both use a B+tree, have ACID semantics with fully serializable transactions, and support lock-free MVCC using a single writer and multiple readers.”

“Bolt is a relatively small code base (<3KLOC) for an embedded, serializable, transactional key/value database so it can be a good starting point for people interested in how databases work.”
```
(from: github/boltdb)


## Summary
- Sigle file store and mmap loading
- single write and multiple read
  - MVCC is target for append B+ tree, nt directly change memory representation of specific page but use new free page
  - Thus read transaction is isolate with write transaction, and before commit of write transaction read will still get old data
  - copy on write
  - two meta data: A/B switch
- Support ACID

## SubPages
- [overview of lmdb](./lmdb.md)
- [boltdb internal](./boltdb_internal.md)
- [boltdb physical file's logic representation](./boltdb_file.md)
- [boltdb transaction](./boltdb_transaction.md)
