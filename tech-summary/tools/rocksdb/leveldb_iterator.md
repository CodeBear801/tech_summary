# LevelDB Iterator

<img src="https://user-images.githubusercontent.com/16873751/99838197-9643be00-2b1d-11eb-860e-2c0e47561ae6.png" alt="rocksdb_iterator" width="600"/>


## TwoLevelIterator

[`TwoLevelIterator`](https://github.com/google/leveldb/blob/28e6d238be73e743c963fc0a26395b783a7565e2/table/two_level_iterator.cc#L61) is used for iterating data inside `sstable`


<img src="https://user-images.githubusercontent.com/16873751/99840548-96de5380-2b21-11eb-94f1-b20aa8217176.png" alt="rocksdb_iterator" width="600"/>



```C++
class TwoLevelIterator: public Iterator {
BlockFunction block_function_; block内部迭代器的生成函数
  void* arg_;  // Most time its TableCahe
  const ReadOptions options_;
  Status status_;
  IteratorWrapper index_iter_; //index_iter_ -> data_iter_
  IteratorWrapper data_iter_; 
  std::string data_block_handle_;
}；

void TwoLevelIterator::SeekToFirst() {
  index_iter_.SeekToFirst(); //Jump to the first block
  InitDataBlock();                 //set data_iter_ based on current block
  if (data_iter_.iter() != NULL) data_iter_.SeekToFirst();
  SkipEmptyDataBlocksForward(); //skip empty block
}
```

Please notice that: There is difference between `TwoLevelIterator` for level0 and for level1~n.  


## MergingIterator

[`MergingIterator`](https://github.com/google/leveldb/blob/28e6d238be73e743c963fc0a26395b783a7565e2/table/merger.cc#L16) is been used during `DBImpl::DoCompactionWork`,  whose target is merging `sstable` between different levels, drop deleted data and duplicate key.  It act similar as [`std::merge`](https://en.cppreference.com/w/cpp/algorithm/merge).  

Construct `MergingIterator(VersionSet::MakeInputIterator()` and construct `TwoLevelIterator` for each `sstable`, they will all be used as `children iterator` to construct `MergingIterator`


<img src="https://user-images.githubusercontent.com/16873751/99840976-5206ec80-2b22-11eb-87e1-f4f67d720e3d.png" alt="rocksdb_iterator" width="600"/>


```C++
  MergingIterator(const Comparator* comparator, Iterator** children, int n)
      : comparator_(comparator),
        children_(new IteratorWrapper[n]),
        n_(n),
        current_(nullptr),
        direction_(kForward) {
    for (int i = 0; i < n; i++) {
      children_[i].Set(children[i]);
    }
  }

  virtual void Seek(const Slice& target) {
    for (int i = 0; i < n_; i++) {
      children_[i].Seek(target);
    }
    FindSmallest();
    direction_ = kForward;
  }
  // All Iterator will seek for target, then will use FindSmallest to find minimum one, record which to current

  void MergingIterator::FindSmallest() {
  IteratorWrapper* smallest = nullptr;
  for (int i = 0; i < n_; i++) {
    IteratorWrapper* child = &children_[i];
    if (child->Valid()) {
      if (smallest == nullptr) {
        smallest = child;
      } else if (comparator_->Compare(child->key(), smallest->key()) < 0) {
        smallest = child;
      }
    }
  }
  current_ = smallest;
}

// Then each Next is    current_++;FindSmallest(), which is operating pointer for merge sort

```