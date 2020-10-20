# sstable

`sstable` means `sorted string table`, which is a simple abstraction to efficiently store large numbers of key-value pairs while optimizing for high throughput, sequential read/write workloads.

## Why SSTable

Sorted string table, random reads and writes are not an option, instead we will want to stream the data in and once we're done, flush it back to disk as a streaming operation. This way, we can amortize the disk I/O costs. Nothing revolutionary, moving right along.  

A "Sorted String Table" then is exactly what it sounds like, it is a file which contains a set of arbitrary, sorted key-value pairs inside. Duplicate keys are fine, there is no need for "padding" for keys or values, and keys and values are arbitrary blobs. Read in the entire file sequentially and you have a sorted index. Optionally, if the file is very large, we can also prepend, or create a standalone key:offset index for fast access. That's all an SSTable is: very simple, but also a very useful way to exchange large, sorted data segments.  

## SSTable and BigTable

How to achieve fast read and write?  
For read, once sstable on disk it is effectively immutable and become static index  
How about write: update and delete?  

## SSTables and Log Structured Merge Trees

<img src="https://user-images.githubusercontent.com/16873751/96522374-1da0c780-1228-11eb-9157-7ee4c569cf65.png
" alt="sstable" width="600"/>  

### Write & Read
writes are always fast regardless of the size of dataset (append-only), and random reads are either served from memory or require a quick disk seek.  

### Update & Delete
Once the SSTable is on disk, it is immutable, hence updates and deletes can't touch the data. Instead, a more recent value is simply stored in MemTable in case of update, and a "tombstone" record is appended for deletes. Because we check the indexes in sequence, future reads will find the updated or the tombstone record without ever reaching the older values! Finally, having hundreds of on-disk SSTables is also not a great idea, hence periodically we will run a process to merge the on-disk SSTables, at which time the update and delete records will overwrite and remove the older data.  

## SSTable format

<img src="https://user-images.githubusercontent.com/16873751/96522495-5476dd80-1228-11eb-8498-7fd74248433c.png
" alt="sstable" width="600"/>  
(image from internet)

- meta block
  - 文件编码
  - 大小
  - 最大key值
  - 最小key值
  
