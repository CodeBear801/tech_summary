# RocksDB

## Why
`LSM + Compaction + BloomFilter + WAL` vs `B+ tree` + `LOG`
```
Very good in write-heavy workload as well as low-latency read workload
```

`Embedded`

<img src="https://user-images.githubusercontent.com/16873751/96497419-80ca3400-11ff-11eb-9a79-62ca5212d408.png" alt="embedded_engine" width="600"/>


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
