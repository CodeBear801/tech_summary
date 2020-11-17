# Compaction


## Code
[DBImpl::MaybeScheduleCompaction() in db/db_impl.cc](https://github.com/google/leveldb/blob/b7d302326961fb809d92a95ce813e2d26fe2e16e/db/db_impl.cc#L658)
```C++
void DBImpl::MaybeScheduleCompaction() 
```
When to call this function  
- Each write, if `memtable` is full, will convert which to `immutable-memtable`, then will call this function
- Each time re-start db, after cover from WAL
- Each read

### Minor compaction
Main logic:  
- dump from `immutable-memtable` to `sstable` on disk
- if level-0's key range have no overlap with current level when try to do more compaction until `config::kMaxMemCompactLevel`(default=2)

Here is an example from when will trigger [`MinorCompaction`](https://github.com/google/leveldb/blob/a6b3a2012e9c598258a295aef74d88b796c47a2b/db/db_test.cc#L1031)  


### Major compaction
Main logic:
- compact level-n sstable with level-(n+1) with overlapped keyrange, multi-path compaction, and generate new level-(n+1) sstable
- if compact from level-0, due to sstables in level 0 have overlapped key range, so there might be more than one sstable join the compaction

When to trigger:  
- size([code](https://github.com/google/leveldb/blob/a6b3a2012e9c598258a295aef74d88b796c47a2b/db/version_set.cc#L650))  
- seek
- By user(operations on leveldb)

[code](https://github.com/google/leveldb/blob/b7d302326961fb809d92a95ce813e2d26fe2e16e/db/db_impl.cc#L697)
```C++
// The trigger of compaction
void DBImpl::BackgroundCompaction() {
    if (is_manual) {
        ManualCompaction* m = manual_compaction_;
        // [Perry]
        // VersionSet* const versions_ GUARDED_BY(mutex_);
        // VersionSet is a double-linked-list which manages current version and all versions which are servicing
        c = versions_->CompactRange(m->level, m->begin, m->end);
    }else {
    c = versions_->PickCompaction();
    }

    if (c == nullptr) {
      // Nothing to do
    } else if (!is_manual && c->IsTrivialMove()) {
         status = versions_->LogAndApply(c->edit(), &mutex_);
    } else {
        CompactionState* compact = new CompactionState(c);
        status = DoCompactionWork(compact);
         CleanupCompaction(compact);
        c->ReleaseInputs();
        RemoveObsoleteFiles();
  }


```
Why there is an abstraction of [`version`](https://github.com/google/leveldb/blob/b7d302326961fb809d92a95ce813e2d26fe2e16e/db/version_set.h#L60)?  
`Version` is db's state after each compaction, it contains meta data of db and a collection of sstable which contains latest state in each level.  When compaction happening, there is sstable addition and deletion, why if they are be read?  To handle such race situation, there is ref count for each version, to represent the situation of read and unread.  So there are multiple version exists for current db, when a version's ref count is 0 and not the latest version, it can be removed from list.  


[`VersionSet::CompactRange`](https://github.com/google/leveldb/blob/9bd23c767601a2420478eec158927882b879bada/db/version_set.cc#L1464)
```C++
// 1. acquire sstables in level-n which are inside key-range[startkey, endkey]
Version::GetOverlappingInputs()

// 2. avoid to compact too many sstable in single round

// 3. collect other sstable needed
VersionSet::SetupOtherInputs
//     for key range from level n, find overlap sstable from level n+1
//     add level-n's sstable if they won't enlarge current key-range
//     get  grandparents_
//     update compact_pointer_ for next round

```


[`DBImpl::DoCompactionWork()`](https://github.com/google/leveldb/blob/a6b3a2012e9c598258a295aef74d88b796c47a2b/db/db_impl.cc#L887)
```C++
// merge sstable which is in sorted order, drop same key and deleted key
DBImpl::DoCompactionWork() {
   // pick sstable from Compaction and construct MergingIterator
   Iterator* input = versions_->MakeInputIterator(compact->compaction);

   
}

/*
DBImpl::DoCompactionWork() (db/db_impl.cc)
实际的 compact 过程就是对多个已经排序的 sstable 做一次 merge 排序，丢弃掉相同 key 以及删
除的数据。
a． 将选出的 Compaction 中的 sstable，构造成
MergingIterator(VersionSet::MakeInputIterator())
a) 对 level-0 的每个 sstable，构造出对应的 iterator：TwoLevelIterator
（TableCache::NewIterator()）。
b) 对非 level-0 的 sstable 构造出 sstable 集合的 iterator：TwoLevelIterator
(NewTwoLevelIterator())
c) 将这些 iterator 作为 children iterator 构造出 MergingIterator
（NewMergingIterator()）。
b． iterator->SeekToFirst()
c． 遍历 Next()
d． 检查并优先 compact 存在的 immutable memtable。
e． 如果当前与 grandparent 层产生 overlap 的 size 超过阈值
29
（Compaction::ShouldStopBefore()），立即结束当前写入的 sstable
（DBImpl::FinishCompactionOutputFile（）），停止遍历。
f. 确定当前 key 的数据是否丢弃。
a) key 是与前面的 key 重复，丢弃。
b) key 是 删 除 （ 检 查 ValueType ） 并且该 key 不 位 于 指 定 的 Snapshot 内（检查
SequnceNumber）并且 key 在 level-n+1 以上的的 level 中不存在（Compaction：：
IsBaseLevelForKey（）），则丢弃。
g. 如 果 当 前 要 写 入 的 sstable 未生成，生成新的 sstable （ DBImpl::
OpenCompactionOutputFile（））。将不丢弃的 key 数据写入（TableBulider::add()）。
h. 如果当前输出的 sstable size 达到阈值（ Compaction::MaxOutputFileSize() 即
MaxFileSizeForLevel（）,当前统一为 kTargeFileSize）,结束输出的 sstable（DBImpl：：
FinishCompactionOutputFile（））。
i. 循环 c-h，直至遍历完成或主动停止。
j. 结束最后一个输出的 sstable（ DBImpl::FinishCompactionOutputFile（）。
k. 更新 compact 的统计信息。
l. 生效 compact 之后的状态。(DBImpl:: InstallCompactionResults())。


*/


```



```C++
  // https://github.com/google/leveldb/blob/a6b3a2012e9c598258a295aef74d88b796c47a2b/db/version_set.cc#L650
  // Apply all of the edits in *edit to the current state.
  void Apply(VersionEdit* edit) {

      // We arrange to automatically compact this file after
      // a certain number of seeks.  Let's assume:
      //   (1) One seek costs 10ms
      //   (2) Writing or reading 1MB costs 10ms (100MB/s)
      //   (3) A compaction of 1MB does 25MB of IO:
      //         1MB read from this level
      //         10-12MB read from next level (boundaries may be misaligned)
      //         10-12MB written to next level
      // This implies that 25 seeks cost the same as the compaction
      // of 1MB of data.  I.e., one seek costs approximately the
      // same as the compaction of 40KB of data.  We are a little
      // conservative and allow approximately one seek for every 16KB
      // of data before triggering a compaction.
      f->allowed_seeks = static_cast<int>((f->file_size / 16384U));
      if (f->allowed_seeks < 100) f->allowed_seeks = 100;
  }
```






## Example

The following example comes from RocksDB, but LevelDB should be very similar.  
RocksDB guarantees efficient disk usage, the size of persistent store is similar to user data size, only 10% is used for extra data.
<img src="https://user-images.githubusercontent.com/16873751/96749219-63bc6f00-137f-11eb-9198-ffbe7854e21c.png" alt="rocksdb_write" width="600"/>

Minor compaction : memory to sstable


<img src="https://user-images.githubusercontent.com/16873751/96751705-5fde1c00-1382-11eb-8c0e-2a5bf9c41b94.png" alt="rocksdb_write" width="300"/>

Major compaction: sstable merge between different level

做major compaction的时候，对于大于level 0的层级，选择其中一个文件就行，但是对于level 0来说，指定某个文件后，本level中很可能有其他SSTable文件的key范围和这个文件有重叠，这种情况下，要找出所有有重叠的文件和level 1的文件进行合并，即level 0在进行文件选择的时候，可能会有多个文件参与major compaction。


<img src="https://user-images.githubusercontent.com/16873751/96751743-6c627480-1382-11eb-98fe-ae1671c437ae.png" alt="rocksdb_write" width="600"/>

<img src="https://user-images.githubusercontent.com/16873751/96752101-de3abe00-1382-11eb-8fed-f5695b90b1c2.png" alt="rocksdb_write" width="600"/>

levelDb在选定某个level进行compaction后，还要选择是具体哪个文件要进行compaction，levelDb在这里有个小技巧， 就是说轮流来，比如这次是文件A进行compaction，那么下次就是在key range上紧挨着文件A的文件B进行compaction，这样每个文件都会有机会轮流和高层的level 文件进行合并。
如果选好了level L的文件A和level L+1层的文件进行合并，那么问题又来了，应该选择level L+1哪些文件进行合并？levelDb选择L+1层中和文件A在key range上有重叠的所有文件来和文件A进行合并。
 也就是说，选定了level L的文件A,之后在level L+1中找到了所有需要合并的文件B,C,D…..等等。剩下的问题就是具体是如何进行major 合并的？就是说给定了一系列文件，每个文件内部是key有序的，如何对这些文件进行合并，使得新生成的文件仍然Key有序，同时抛掉哪些不再有价值的KV 数据。


<img src="https://user-images.githubusercontent.com/16873751/96751773-73898280-1382-11eb-9dc0-4f5f0b74dc51.png" alt="rocksdb_write" width="600"/>

Major compaction的过程如下：对多个文件采用多路归并排序的方式，依次找出其中最小的Key记录，也就是对多个文件中的所有记录重新进行排序。之后采取一定的标准判断这个Key是否还需要保存，如果判断没有保存价值，那么直接抛掉，如果觉得还需要继续保存，那么就将其写入level L+1层中新生成的一个SSTable文件中。就这样对KV数据一一处理，形成了一系列新的L+1层数据文件，之前的L层文件和L+1层参与compaction 的文件数据此时已经没有意义了，所以全部删除。这样就完成了L层和L+1层文件记录的合并过程。
那么在major compaction过程中，判断一个KV记录是否抛弃的标准是什么呢？其中一个标准是:对于某个key来说，如果在小于L层中存在这个Key，那么这个KV在major compaction过程中可以抛掉。因为我们前面分析过，对于层级低于L的文件中如果存在同一Key的记录，那么说明对于Key来说，有更新鲜的Value存在，那么过去的Value就等于没有意义了，所以可以删除。

