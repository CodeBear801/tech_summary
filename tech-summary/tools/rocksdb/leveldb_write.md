# LevelDB write

## Key Points
- a single put operation just result one disk write and one memory write(一次写入操作只涉及一次磁盘顺序写和一次内存写入)

## Big picture
```
Memtable -> Immutable Memtable -> SSTable
(skiplist)                     (bloomfilter)
```


## Internal

### Example
[code](https://github.com/facebook/rocksdb/blob/00751e4292e55c1604b28b7b93fe7a538fa05f29/examples/simple_example.cc#L35)
```c++
Status s = DB::Open(options, kDBPath, &db);
s = db->Put(WriteOptions(), "key1", "value");
```

1. Update WAL(more info: WAL)
2. Update Memory table(more info: skiplist)

3. Switch to immutable memtable
4. Compaction(more info: compaction)


### Notes for rocksdb

<img src="https://user-images.githubusercontent.com/16873751/96521492-4922b280-1226-11eb-9803-a1d0768713f4.png" alt="rocksdb_read" width="600"/>

<br/><br/><br/>

### Code 
[DBImpl::Write in db/db_impl.cc](https://github.com/google/leveldb/blob/b7d302326961fb809d92a95ce813e2d26fe2e16e/db/db_impl.cc#L1196)

