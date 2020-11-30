

## Keywords

- why (Motivation, why not berkley db)
  - Turning complexity
    + data comes through 3 separate layers of caches
    + Each layer has different size/speed traits
    + Balancing 3 layers against each other can be a difficult juggling act
  - Cache inefficiency
  - lock mgr and deadlocks
    + tow level of locking is needed, due to BDB locks are too slow
    + deadlocks occurring routinely 
  - Logging
- key/Value store using B+ tree
- single writer + N readers
  - writers don't block readers, readers don't block writers
  - Reads scale perfectly linearly with available CPUs
  - No dead lock
- MVCC
  - MVCC is implemented by copy-on-write
  - "Full" MVCC can be extremely resource intensive !!! -> compaction, slow
  - LMDB nomally maintains only **two** version of the DB
  - free list for unused pages
- mmap files
  - Reads are satisfied directly from mmap, no malloc or memory overhead
  - Writes can be performed directly to the mmap, no writer buffers
  - relies on OS/filesystem cache
  -  Batched writes
- B+ tree
- Transaction
  - A single mutex serializes all write transactions
  - Transactions are stamped with an ID which is a monotonically increasing integer(write, when actually modify data)
  - The currently valid meta page is chosen based on the greatest transaction ID in each meta page
  - During commit, 2 flush -> data pages written then meta page is updated


(from:[The Lightning Memory-mapped Database](https://www.infoq.com/presentations/lmdb-lighting-memory-mapped-database/#downloadPdf/))