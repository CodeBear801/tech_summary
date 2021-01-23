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
GFS chunkserver replaced by `D`.  `D Server`, act as a GFS chunkserver running on Borg.  
Its the lowest level of storage layer, and is designed as the only app to direct read/write data on physical disk, all of physical disks have been connected to server running D.   

## Arch
<img src="https://user-images.githubusercontent.com/16873751/105611815-7ffbed80-5d6c-11eb-820c-5ac31153f1e4.png" alt="colossus_arch" width="1000"/>
<br/>

- `Colossus client`: probably the most complex part of the system
  - lots of functions go directly in the client, such as
    + software RAID
    + application encoding chosen
- `Curators`: foundation of Colossus, its scalable metadata service
  - can scale out horizontally
  - built on top of a NoSQL database like `BigTable`
  - allow Colossus to scale up by over a 100x over the largest GFS
- `D servers`: simple network attached disks
- `Custodians`: background storage managers, handle such as disk space balancing, and RAID construction
  - ensures the durability and availability
  - ensures the system is working efficiently
- `Data`: there are hot data (e.g. newly written data) and cold data
- Mixing flash and spinning disks
  - really efficient storage organization
    + just enough flash to push the I/O density per gigabyte of data
    + just enough disks to fill them all up
  - use flash to serve really hot data, and lower latency
  - regarding to disks
    + equal amounts of hot data across disks
        * each disk has roughly same bandwidth
        * spreads new writes evenly across all the disks so disk spindles are busy
  - rest of disks filled with cold data
    + moves older cold data to bigger drives so disks are full


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
- [Google and evolution of big-data](https://usefulstuff.io/2013/05/google-and-evolution-of-big-data/)
- [Evolution of Google FS Talk](http://sghose.me/talks/storage%20systems/2015/11/23/GFS-Talk/)
- [A peek behind the VM at the Google Storage infrastructure](https://www.youtube.com/watch?v=q4WC_6SzBz4&feature=youtu.be)
  - [discussions](http://tab.d-thinker.org/showthread.php?tid=1&pid=332#pid332)
  - [A discussion between Kirk McKusick and Sean Quinlan about the origin and evolution of the Google File System.](https://queue.acm.org/detail.cfm?id=1594206)
- [Storage Architecture and Challenges by Andrew Fikes](https://cloud.google.com/files/storage_architecture_and_challenges.pdf)
- [Colossus: Successor to the Google File System (GFS)](https://www.systutorials.com/colossus-successor-to-google-file-system-gfs/)
- [Google File System II: Dawn of the Multiplying Master Nodes](https://www.theregister.com/2009/08/12/google_file_system_part_deux/)