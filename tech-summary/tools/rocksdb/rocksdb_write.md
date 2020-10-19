# RocksDB write

## Key Points
- 一次写入操作只涉及一次磁盘顺序写和一次内存写入

## Big picture
```
Memtable -> Immutable Memtable -> SSTable
(skiplist)                     (bloomfilter)
```

## Example
[code](https://github.com/facebook/rocksdb/blob/00751e4292e55c1604b28b7b93fe7a538fa05f29/examples/simple_example.cc#L35)
```c++
Status s = DB::Open(options, kDBPath, &db);
s = db->Put(WriteOptions(), "key1", "value");
```

1. Update WAL(more info: WAL)
2. Update Memory table(more info: skiplist)

3. Switch to immutable memtable
4. Compaction(more info: compaction)

### Code in leveldb


