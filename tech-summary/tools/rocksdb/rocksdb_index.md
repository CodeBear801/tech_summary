# RocksDB

## Why
`LSM + Compaction + BloomFilter + WAL` **vs** `B+ tree + LOG`
```
Very good in write-heavy workload as well as low-latency read workload

minimize random writes
```

`Embedded`

<img src="https://user-images.githubusercontent.com/16873751/96497419-80ca3400-11ff-11eb-9a79-62ca5212d408.png" alt="embedded_engine" width="600"/>


## What
- Key-Value persistent store
- Point/Range lookup
- Optimized for flash
- Embedded storage of application, storage engine of database

## Arch


<img src="https://user-images.githubusercontent.com/16873751/96756920-979c9200-1389-11eb-984c-34957c8248a5.png" alt="arch1" width="600"/>  


## Interface
- Keys and Values are byte arrays
- Keys have total order
- Update : Put/Delete/Merge
- Queries : Get/Iterator
- Consistency: atomic multi-put, multi-get, iterator, snapshot read, transactions

## Key Techniques 
- [get()](./leveldb_read.md)
- [put()](./leveldb_write.md)
- [skiplist](./skiplist.md)
- [wal](./write_ahead_log.md)
- [sstable](./sstable.md)
- [compaction](./compaction.md)
- [bloomfilter](./bloomfilter.md)