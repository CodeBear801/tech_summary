# LevelDB arena

Arena is a grain level memory pool behind `levelDB`'s `skiplist` and `memtable`.  `Arena` is not for general purpose memory pool and highly customized for leveldb, very good for small memory allocation.


<img src="https://user-images.githubusercontent.com/16873751/98845225-415bc580-2402-11eb-9447-1f1d77667540.png" alt="botdb_bucket" width="300"/>
<br/>



Here is the interface [util/arena.h](https://github.com/google/leveldb/blob/9bd23c767601a2420478eec158927882b879bada/util/arena.h#L16)
```C++
  // Return a pointer to a newly allocated memory block of "bytes" bytes.
  char* Allocate(size_t bytes);

  // Allocate memory with the normal alignment guarantees provided by malloc.
  char* AllocateAligned(size_t bytes);

```

Allocate is called by [`memtable::add()`](https://github.com/google/leveldb/blob/b7d302326961fb809d92a95ce813e2d26fe2e16e/db/memtable.cc#L89) to allocate bytes of space, the implementation is [here](https://github.com/google/leveldb/blob/9bd23c767601a2420478eec158927882b879bada/util/arena.h#L55)

```C++
inline char* Arena::Allocate(size_t bytes) {
    //[perry] try to allocate space in current block if has enough space, 
    // otherwise
    return AllocateFallback(bytes);
}

char* Arena::AllocateFallback(size_t bytes) {

  if (bytes > kBlockSize / 4) {
    // Object is more than a quarter of our block size.  Allocate it separately
    // to avoid wasting too much space in leftover bytes.
    char* result = AllocateNewBlock(bytes);
    return result;
  }

  // We waste the remaining space in the current block.
  alloc_ptr_ = AllocateNewBlock(kBlockSize);
  alloc_bytes_remaining_ = kBlockSize;

  char* result = alloc_ptr_;
  alloc_ptr_ += bytes;
  alloc_bytes_remaining_ -= bytes;
  return result;
}

```

`AllocateAligned` is called by `skiplist` to allocate node object.([code](https://github.com/google/leveldb/blob/b7d302326961fb809d92a95ce813e2d26fe2e16e/db/skiplist.h#L184))

All release happened in destructor
```C++
Arena::~Arena() {
  for (size_t i = 0; i < blocks_.size(); i++) {
    delete[] blocks_[i];
  }
}
```



