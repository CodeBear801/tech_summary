# Colossus

## Why

`GFS`'s problems
- GFS master
  - One machine not large enough for large FS
  - Single bottleneck for metadata operations
  - Fault tolerant, not HA
- Predictable performance
  - no guarantees of latency

`Colossus`'s goal
- Bigger
- Faster
- More predictable tail latency

GFS master replaced by `Colossus`  
GFS chunkserver replaced by `D`.  `D Server`, act as a GFS chunkserver running on Borg.  Its the lowest level of storage layer, and is designed as the only app to direct read/write data on physical disk, all of physical disks have been connected to server running D.   

## How

<img src="https://user-images.githubusercontent.com/16873751/99613036-a481c500-29cb-11eb-8ab9-6e3901faa825.png" alt="colossus_arch" width="600"/>
<br/>

`Colossus` choose `Bigtable` to record meta data due to BigTable solves many of the hard problems:
- Automatically shards data across tablets
- Locates tablets via metadata lookups
- Easy to use semantics
- Efficient point lookups and scans
- **File system metadata kept in an 0n-memory locality group**

The idea of **Use colossus to store Colossus's meta data**
- metadata is ~ 1/10000 the size of data
- If we host a Colossus on Colossus, 100PB data -> 10TB meta data, 10TB meta data -> 1GB metadata, 1GB metadata -> 100KB data
- The data is smaller enough to put into Chubby
- LSM tree minimize random write by LSM tree.  For GFS/Colossus, it will trigger communicate with Master/CFS only when creating new data block, most of other time just communication with ChunkServer/D Server.  Meanwhile, the compaction/merge also decrease the frequency of creating new data blocks. 

Here is a picture represent this:  
<img src="https://user-images.githubusercontent.com/16873751/99704210-eacb3880-2a4c-11eb-929d-7e333df6a1b4.png" alt="colossus_arch" width="600"/>
<br/>
(picture from: https://levy.at/blog/22)

Metadata in `bigtable`  

<img src="https://user-images.githubusercontent.com/16873751/99617107-70120700-29d3-11eb-8bbc-e4e2ceaffb0e.png" alt="colossus_arch" width="600"/>
<br/>

GFS master -> `CFS`  

<img src="https://user-images.githubusercontent.com/16873751/99617132-7e602300-29d3-11eb-9d46-1fecec29133c.png" alt="colossus_arch" width="600"/>
<br/>

`Colossus` cluster example  
 
<img src="https://user-images.githubusercontent.com/16873751/99616691-853a6600-29d2-11eb-8990-e718ae5d9654.png" alt="colossus_arch" width="600"/>
<br/>





## More info
- [Storage Architecture and Challenges](https://cloud.google.com/files/storage_architecture_and_challenges.pdf)
- [Google File System及其继任者Colossus](https://levy.at/blog/22)
- [谷歌Colossus文件系统的设计经验](https://www.sohu.com/a/413895492_673711)