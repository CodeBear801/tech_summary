# RocksDB

## Why
1. `LSM + Compaction + BloomFilter + WAL` **vs** `B+ tree + LOG`
```
Very good in write-heavy workload as well as low-latency read workload

minimize random writes

an alternative space and write-optimized database technology
```
<br/>

2. `Embedded`
<br/>
<img src="https://user-images.githubusercontent.com/16873751/96497419-80ca3400-11ff-11eb-9a79-62ca5212d408.png" alt="embedded_engine" width="600"/>

<br/>

3. [MySQL vs RocksDB](./rocksdb_vs_mysql.md)

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

## Internal  
- [get()](./leveldb_read.md)
- [put()](./leveldb_write.md)
- [skiplist](./leveldb_skiplist.md)
- [memtable](./leveldb_memtable.md)
- [sstable](./leveldb_sstable.md)
- [compaction](./leveldb_compaction.md)
- [wal](./leveldb_write_ahead_log.md)
- [bloomfilter](./leveldb_bloomfilter.md)
- [arena](./leveldb_arena.md)
- [iterator](./leveldb_iterator.md)

## More info
- [Knowledge share on 11162020](https://www.slideshare.net/ssuser4c810e/rocksdb-vs-blotdb)