# RocksDB read

## Big picture
```
Memtabel -> immutable memtable -> sstable level 0 -> sstable level 1 …
```

## Example

[code](https://github.com/facebook/rocksdb/blob/00751e4292e55c1604b28b7b93fe7a538fa05f29/examples/simple_example.cc#L39)
```c++
Status s = DB::Open(options, kDBPath, &db);
s = db->Get(ReadOptions(), "key1", &value);
```

Steps of Get()
<img src="https://user-images.githubusercontent.com/16873751/96516475-18d61680-121c-11eb-9eeb-c8e38e2ef13f.png" alt="rocksdb_get" width="1000"/>

1. check memtable
2. if not find, check immutable-memtable
3. if not find, based on `current version`, iterate `FileMetaData` for all `sstable` from level 0 -> level max
   - there could be multiple `sstable` in level 0, iterate them based on freshness
4. Identify which `block` belong to for target key, based on `data index block` of `sstable` in `Table cache`
5. If there is cache miss, load `data index block` from `sstable` on disk
6. When `block` has been identified, first check from `blockcache`
7. If `blockcache miss`, load `datablock` from `sstable`
8. Still not find, iterate step 3 ~ 7