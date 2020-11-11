# LevelDB memtable


- `Put` operation is wrapped with WriteBatch([code](https://github.com/google/leveldb/blob/b7d302326961fb809d92a95ce813e2d26fe2e16e/db/db_impl.cc#L1464))

```C++
Status DB::Put(const WriteOptions& opt, const Slice& key, const Slice& value) {
  WriteBatch batch;
  batch.Put(key, value);
  return Write(opt, &batch);
}
```

- Here is the definition of [`WriteBatch`](https://github.com/google/leveldb/blob/b7d302326961fb809d92a95ce813e2d26fe2e16e/db/write_batch.cc#L5)

```C++
// WriteBatch::rep_ :=
//    sequence: fixed64
//    count: fixed32        [Perry] WriteBatch could contains multiple records
//    data: record[count]
// record :=
//    kTypeValue varstring varstring         |
//    kTypeDeletion varstring    [Perry] No value size and value
// varstring :=
//    len: varint32
//    data: uint8[len]


```

Let's say we want to write `<key= TestKey, Value=TestValue>`  
<img src="https://user-images.githubusercontent.com/16873751/98858230-1f1f7300-2415-11eb-981d-9d3c741e91d1.png" alt="write_ahead_log" width="600"/>
<br/>



