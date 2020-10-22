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

