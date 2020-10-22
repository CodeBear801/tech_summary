# RocksDB vs MySQL

## MySQL

As internal storage engine of `MySQL`, `InnoDB` is great for performance and reliability, but there are some inefficiencies on space and write amplification with flash.

### MySQL Arch

<img src="https://user-images.githubusercontent.com/16873751/96908037-f0832d80-1450-11eb-96f7-57602d1b9162.png" alt="innodb" width="400"/>  
Image from: https://time.geekbang.org/column/article/68633  

<br/>

### Get/Put

<img src="https://user-images.githubusercontent.com/16873751/96909914-85872600-1453-11eb-9ae1-b9612e57461a.png" alt="embedded_engine" width="400"/>  
Image from: https://time.geekbang.org/column/article/68633   

`redo log` belongs to `InnoDB`, records physical log, such as "modification on a data page"  
`binlog` belongs to `MySQL`, records all logic operation, such as "Add 1 to c where ID = 2", WAL  

<br/>

### Notes

- `InnoDB` write amplification
   + index values are stored in leaf nodes and sorted by keys, the database working set doesn’t fit in memory and keys are updated in a random pattern.  Updating one row requires a number of page reads, makes several pages dirty, and forces many dirty pages to be written back to storage. 
   + In addition, `InnoDB` writes dirty pages to storage twice to support recovery from partial page writes on unplanned machine failure. 

<img src="https://user-images.githubusercontent.com/16873751/96910616-9c7a4800-1454-11eb-8574-df16883033bc.png" alt="embedded_engine" width="400"/><br/>
Image from: https://engineering.fb.com/core-data/myrocks-a-space-and-write-optimized-mysql-database/


- `InnoDB` compression alignment results in extra unused space 

<img src="https://user-images.githubusercontent.com/16873751/96910593-92f0e000-1454-11eb-9c17-8e1a45dd558a.png" alt="embedded_engine" width="400"/><br/>
Image from: https://engineering.fb.com/core-data/myrocks-a-space-and-write-optimized-mysql-database/

## RocksDB
An alternative space and write-optimized database technology.

<img src="https://user-images.githubusercontent.com/16873751/96756920-979c9200-1389-11eb-984c-34957c8248a5.png" alt="arch1" width="600"/>  

### Write

Avoid random write  

More info: [notes about get()](./leveldb_write.md)

```
Memtable -> Immutable Memtable -> SSTable
(skiplist)                     (bloomfilter)
```
Let's write several value into database
```
db.put("bill", "1");
db.put("david", "2");
db.put("lily", "3");
```
Those data will first write into WAL, and then into memory table, which is in supported by skip-list.

When in-memory table reach pre-defined sized limit, it will change current memory table as immutable memory table    

Compaction engine will try to merge immutable memory table with `sstable` in level 0.  Level 0 table always record all range of keys.  
<img src="https://user-images.githubusercontent.com/16873751/96912453-44911080-1457-11eb-9786-
312880c54beb.png" alt="rocksdb_write" width="300"/><br/>

When level 0 reach its size limit, compaction engine will be triggered to merge with level 1.  Let's say Level 1 record <lily, 4> in previous record, after compaction, only latest data be recorded.  
<img src="https://user-images.githubusercontent.com/16873751/96912482-4fe43c00-1457-11eb-873b-3344897c5369.png" alt="rocksdb_write" width="300"/><br/>


Then level 1 merge with level 2, etc...  
<img src="https://user-images.githubusercontent.com/16873751/96912516-57a3e080-1457-11eb-8f4e-c66293c9a777.png" alt="rocksdb_write" width="300"/><br/>



### Read
More info: [notes about get()](./leveldb_read.md)

```
Memtabel -> immutable memtable -> sstable level 0 -> sstable level 1 …
```


