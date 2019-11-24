# Google File System

## Challenge for distribute file system

### Local vs distribute

Local file system | Distribute file system
---|---
Logic unit called files <br/> absolute path and logic path + name | Remote access <br/>Must support concurrency <br/>Replication & local cache<br/>Scale

### GFS's challenge
- With many machines failures are common
- Many concurrent readers and writers, high performance is required
   + MR jobs read and store **final** result in GFS
- Use network efficiently

### GFS's motivation
- Redundant storage of massive amount of data on cheap and unreliable devices
- Good for read.  Good for write

### GFS's Assumption
- High component failure rates(commodity components)
- Modest number of huge files
- Files are write once, mostly append to
- Large streaming reads
- High sustained throughput favored over low latency

## GFS's architecture

<img src="resources/pictures/gfs_arch.png" alt="gfs_arch" width="600"/>
 <br/>

- master must avoid single point of failure
- master has operation log
- shadow masters that lag a little behind master, can be promoted to master
- For each file, master just record the mapping of trunk index <-> trunk server, let trunk server handle how the trunk be recorded

<img src="resources/pictures/gfs_mutations.png" alt="gfs_mutations" width="600"/>
 <br/>
## Master's responsibility
- Metadata store
- Namespace management locking
- Periodic communication with chunkservers
- chunk creation, re-replication, rebalancing
- GC
- master keeps state in memory(64 bytes of metadata per each chunk)

## Operations

### Read

<img src="resources/pictures/gfs_read.png" alt="gfs_read" width="600"/>
 <br/>




## Fault tolerance
- fault tolerance of data (3 copies)
- High availability (fast recovery, chunk replication, shadow master)
- Data integrity(checksum)

## Summary
(based on original paper)
Great | Less well
---|---
huge sequential reads and writes<br/>appends<br/>huge throughput (3 copies, striping)<br/>fault tolerance of data (3 copies) |     fault-tolerance of master<br/> small files (master a bottleneck)<br/>clients may see stale data<br/>appends maybe duplicated

## More info
- [The Google File System](https://static.googleusercontent.com/media/research.google.com/en//archive/gfs-sosp2003.pdf)
- [Distributed computing seminar lecture 3 - distributed file systems - Aaron kimball 2007](https://www.youtube.com/watch?v=5Eib_H_zCEY&t=2394s) [slides](Distributed computing seminar   lecture 3 - distributed file systems)

