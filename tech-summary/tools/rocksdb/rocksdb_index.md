# RocksDB

## Why
`LSM + Compaction + BloomFilter + WAL` vs `B+ tree` + `LOG`
```
Very good in write-heavy workload as well as low-latency read workload
```

`Embedded`




## What
- Key-Value persistent store
- Point/Range lookup
- Optimized for flash
- Embedded storage of application, storage engine of database

## Interface
- Keys and Values are byte arrays
- Keys have total order
- Update : Put/Delete/Merge
- Queries : Get/Iterator
- Consistency: atomic multi-put, multi-get, iterator, snapshot read, transactions

 ## Key Techniques 

