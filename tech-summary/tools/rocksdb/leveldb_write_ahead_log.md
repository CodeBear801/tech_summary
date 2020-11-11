# LevelDB Write ahead log

<img src="https://user-images.githubusercontent.com/16873751/97507573-b0351b00-193a-11eb-8e08-df65d43cb89e.png" alt="write_ahead_log" width="600"/>
<br/>
Image from: https://xixymyciqecu.schmidt-grafikdesign.com/write-ahead-log-postgresql-database-32348ng.html    
<br/>

Log is written record by record, and each fixed sized block contains 1 or many records  

<img src="https://user-images.githubusercontent.com/16873751/98849390-fb096500-2407-11eb-91a8-781d0610b6ca.png" alt="write_ahead_log" width="600"/>
<br/>

[log_fotmat.h](https://github.com/google/leveldb/blob/b7d302326961fb809d92a95ce813e2d26fe2e16e/db/log_format.h#L14)
```C++
enum RecordType {
  // Zero is reserved for preallocated files
  kZeroType = 0,

  kFullType = 1,

  // For fragments
  kFirstType = 2,
  kMiddleType = 3,
  kLastType = 4
};

// Header is checksum (4 bytes), length (2 bytes), type (1 byte).
static const int kHeaderSize = 4 + 2 + 1;
```
more info: https://github.com/google/leveldb/blob/master/doc/log_format.md  

WAL represent information of current memtable, when memtable change to immutable-memtable, WAL will also be persist.
