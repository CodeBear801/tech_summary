# LevelDB read

## Big picture
```
Memtabel -> immutable memtable -> sstable level 0 -> sstable level 1 …
```
## Internal

### Example

[code](https://github.com/facebook/rocksdb/blob/00751e4292e55c1604b28b7b93fe7a538fa05f29/examples/simple_example.cc#L39)
```c++
Status s = DB::Open(options, kDBPath, &db);
s = db->Get(ReadOptions(), "key1", &value);
```
### Flow

Steps of Get()(based on leveldb, which is simpler)
<img src="https://user-images.githubusercontent.com/16873751/96516475-18d61680-121c-11eb-9eeb-c8e38e2ef13f.png" alt="rocksdb_get" width="1000"/>

1. check memtable
2. if not find, check immutable-memtable
3. if not find, based on `current version`, iterate `FileMetaData` for all `sstable` from level 0 -> level max
   - there could be multiple `sstable` in level 0, iterate them based on freshness
4. Identify which `block` belong to for target key, based on `data index block` of `sstable` in `Table cache`
5. If there is cache miss, load `data index block` from `sstable` on disk
6. When `block` has been identified, first check from `blockcache`
7. If `blockcache miss`, load `datablock` from `sstable`
8. Still not find, iterate step 3 ~ 7


### Notes for leveldb
为啥是如此顺序
```
Memtabel -> immutable memtable -> sstable level 0 -> sstable level 1 …
```
从信息的更新时间来说，很明显Memtable存储的是最新鲜的KV对；Immutable Memtable中存储的KV数据对的新鲜程度次之；而所有SSTable文件中的KV数据新鲜程度一定不如内存中的Memtable和Immutable Memtable的。对于SSTable文件来说，如果同时在level L和Level L+1找到同一个key，level L的信息一定比level L+1的要新。也就是说，上面列出的查找路径就是按照数据新鲜程度排列出来的，越新鲜的越先查找。  

Level L + 1 是 level L 经过compaction 生成的，所以level L 一定比L+1要新鲜  

如何快速地找到key对应的value值？在LevelDb中，level 0一直都爱搞特殊化，在level 0和其它level中查找某个key的过程是不一样的。因为level 0下的不同文件可能key的范围有重叠，某个要查询的key有可能多个文件都包含，这样的话LevelDb的策略是先找出level 0中哪些文件包含这个key（manifest文件中记载了level和对应的文件及文件里key的范围信息，LevelDb在内存中保留这种映射表）， 之后按照文件的新鲜程度排序，新的文件排在前面，之后依次查找，读出key对应的value。而如果是非level 0的话，因为这个level的文件之间key是不重叠的，所以只从一个文件就可以找到key对应的value。  


如果给定一个要查询的key和某个key range包含这个key的SSTable文件，那么levelDb是如何进行具体查找过程的呢？levelDb一般会先在内存中的Cache中查找是否包含这个文件的缓存记录，如果包含，则从缓存中读取；如果不包含，则打开SSTable文件，同时将这个文件的索引部分加载到内存中并放入Cache中。 这样Cache里面就有了这个SSTable的缓存项，但是只有索引部分在内存中，之后levelDb根据索引可以定位到哪个内容Block会包含这条key，从文件中读出这个Block的内容，在根据记录一一比较，如果找到则返回结果，如果没有找到，那么说明这个level的SSTable文件并不包含这个key，所以到下一级别的SSTable中去查找。 

### Notes for rocksdb

<img src="https://user-images.githubusercontent.com/16873751/96518047-dcf08080-121e-11eb-86bd-123db83d22c7.png" alt="rocksdb_get2" width="600"/>

- Find key in mem table, directly return
- If not, go down
  + Level 0 is special, cover all level of key
  + Level1, 2, 3, are partitions, a ~ c, d ~ f
  + If you want to find the key, you need to go to all files in level 0
  + If you don't find the key in level 0, then need go down, level 1, level 2, (just one file)

#### Point query(single key)
- bloom filter avoid disk IO.  `Bloomfilter` could tell you, this key certainly not here, or the key might be here.  
- at most 1 physical read


#### Range query
- bloom filters don't help
- short scans at most 1 or 2 physical reads
  - memory would helps, cache everything above level 4, if so only one IO
  - [How do range scan in rocksdb](https://github.com/facebook/rocksdb/issues/204)
  - [DeleteRange: A New Native RocksDB Operation](https://rocksdb.org/blog/2018/11/21/delete-range.html)

### Code
