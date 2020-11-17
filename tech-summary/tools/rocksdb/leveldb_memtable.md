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
//    kTypeDeletion varstring    [Perry] No value size and value for delete
// varstring :=
//    len: varint32
//    data: uint8[len]


```

Let's say we want to write `<key= TestKey, Value=TestValue>`, here is the sample of `WriteBatch`:  
<img src="https://user-images.githubusercontent.com/16873751/98858230-1f1f7300-2415-11eb-981d-9d3c741e91d1.png" alt="write_ahead_log" width="600"/>
<br/>

It will be wrapped as a [`Writer`](https://github.com/google/leveldb/blob/b7d302326961fb809d92a95ce813e2d26fe2e16e/db/db_impl.cc#L1196) and then insert into a [`deque`](https://github.com/google/leveldb/blob/b7d302326961fb809d92a95ce813e2d26fe2e16e/db/db_impl.h#L186) 
```C++
  Writer w(&mutex_);
  w.batch = updates;
  w.sync = options.sync;
  w.done = false;
``` 
<img src="https://user-images.githubusercontent.com/16873751/98878638-4f2c3d80-2438-11eb-812b-fa2d6611af26.png" alt="write_ahead_log" width="1200" height = "400"/>
<br/>

`Writer` need to write two part of information, one is `WAL` on disk, in case of failure
```C++
      status = log_->AddRecord(WriteBatchInternal::Contents(write_batch));
```

Another part is add <key,value> pair into memory table([code](https://github.com/google/leveldb/blob/b7d302326961fb809d92a95ce813e2d26fe2e16e/db/db_impl.cc#L1234))
```C++

-> WriteBatchInternal::InsertInto(write_batch, mem_);
```
[`WriteBatchInternal::InsertInto`](https://github.com/google/leveldb/blob/b7d302326961fb809d92a95ce813e2d26fe2e16e/db/write_batch.cc#L132:1)
```C++
Status WriteBatchInternal::InsertInto(const WriteBatch* b, MemTable* memtable) {
  MemTableInserter inserter;
  inserter.sequence_ = WriteBatchInternal::Sequence(b);
  inserter.mem_ = memtable;
  return b->Iterate(&inserter);
}

// Perry: Inside Iterate, it will call inserter's put to put key value
Status WriteBatch::Iterate(Handler* handler) const {
}

// Calling memtable's add
class MemTableInserter : public WriteBatch::Handler {
  void Put(const Slice& key, const Slice& value) override {
    mem_->Add(sequence_, kTypeValue, key, value);
    sequence_++;
  }
}

```
Finally, come to call [`memtable's Add()`](https://github.com/google/leveldb/blob/b7d302326961fb809d92a95ce813e2d26fe2e16e/db/memtable.cc#L76) and then [insert <key, value> into skiplist](https://github.com/google/leveldb/blob/b7d302326961fb809d92a95ce813e2d26fe2e16e/db/skiplist.h#L340)

```C++
void MemTable::Add(SequenceNumber s, ValueType type, const Slice& key,
{
    // typedef SkipList<const char*, KeyComparator> Table;
    table_.Insert(buf);

}

```
For more information about skiplist, please go to [here](./leveldb_skiplist.md), a great example about how to test current reading for skiplist can be found [here](https://github.com/google/leveldb/blob/b7d302326961fb809d92a95ce813e2d26fe2e16e/db/skiplist_test.cc#L342)


