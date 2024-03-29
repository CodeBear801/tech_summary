- [Spark shuffle](#spark-shuffle)
	- [What is Shuffle](#what-is-shuffle)
	- [How](#how)
		- [How to generate shuffle file](#how-to-generate-shuffle-file)
		- [How shuffle file is read](#how-shuffle-file-is-read)
		- [Shuffle Manager](#shuffle-manager)
		- [External shuffle service](#external-shuffle-service)
	- [Chanllege](#chanllege)
		- [Cosco from Facebook](#cosco-from-facebook)
		- [Magnet  from LinkedIn](#magnet--from-linkedin)

# Spark shuffle

## What is Shuffle

<img src="https://user-images.githubusercontent.com/16873751/108574287-d278ed00-72cb-11eb-961e-f98f62030db7.png" alt="spark_arch" width="600"/> 

<br/>

*** 

## How

### How to generate shuffle file

```java
/**
DAGScheduler -> create job, stage, task

getShuffleDependencies(RDD) -> will retrieve all parent shuffle dependencies for given RDD, via shuffleDependency

shuffleDependency is created by ShuffleExchangeExec's  as ShuffleDependency[Int, InternalRow, InternalRow]
         - Int is the partition number get from RoundRobinPartitioning, HashPartitioning, RangePartitioning or SinglePartition
         - InternalRow is  the corresponding row
         - InternalRow is combined rows after the shuffle.

shuffle writer retrieves the partition stored in the ShuffleDependency and applies a write method. 

*/


```java
// how hash is generated
     case h: HashPartitioning =>
        val projection = UnsafeProjection.create(h.partitionIdExpression :: Nil, outputAttributes)
        row => projection(row).getInt(0)
def partitionIdExpression: Expression = Pmod(new Murmur3Hash(expressions), Literal(numPartitions))


// how to write shuffle
val rddAndDep = ser.deserialize[(RDD[_], ShuffleDependency[_, _, _])]( // ...
    val rdd = rddAndDep._1
    val dep = rddAndDep._2
    dep.shuffleWriterProcessor.write(rdd, dep, mapId, context, partition)
}

// ShuffleWriteProcessor
  def write(
      rdd: RDD[_],
      dep: ShuffleDependency[_, _, _],
      mapId: Long,
      context: TaskContext,
      partition: Partition): MapStatus = {
}

```

### How shuffle file is read

Result generated from map stage:
```
|-- 09
|-- 0a
|-- 0b
|-- 0c
|   `-- shuffle_0_0_0.data
|-- 0d
|   `-- shuffle_0_3_0.index
|-- 0e
|-- 0f
|   `-- shuffle_0_1_0.index
|-- 11
|-- 15
|   `-- shuffle_0_1_0.data
```

```
the DAGScheduler executes the ShuffledRDD holding the ShuffleDependency introduced in 
the previous section.

the compute(split: Partition, context: TaskContext) method will return all records that 
should be returned for the Partition from the signature.

The compute method will create a ShuffleReader instance that will be responsible, 
through its read() method, to return an iterator storing all rows that are set for 
the specific reducer's

MapOutputTracker: will tell reader which file it should fetch, the tracker is called 
to retrieve all shuffle locations for the given shuffle id before creating the shuffle 
files reader
```

```java

// https://github.com/apache/spark/blob/60c71c6d2d38163468c0f428fd1f33015b58c32c/core/src/main/scala/org/apache/spark/shuffle/sort/SortShuffleManager.scala#L124

override def getReader[K, C](	
	handle: ShuffleHandle,
	startMapIndex: Int,
	endMapIndex: Int,
	startPartition: Int,
	endPartition: Int,
	context: TaskContext,
	metrics: ShuffleReadMetricsReporter): ShuffleReader[K, C] = {
	val blocksByAddress = SparkEnv.get.mapOutputTracker.getMapSizesByExecutorId(
	handle.shuffleId, startMapIndex, endMapIndex, startPartition, endPartition)
	new BlockStoreShuffleReader(
	handle.asInstanceOf[BaseShuffleHandle[K, _, C]], blocksByAddress, context, metrics,
	shouldBatchFetch = canUseBatchFetch(startPartition, endPartition, context))
	}


// https://github.com/apache/spark/blob/c6994354f70061b2a15445dbd298a2db926b548c/core/src/main/scala/org/apache/spark/storage/ShuffleBlockFetcherIterator.scala#L576

/**	
	* Fetches the next (BlockId, InputStream). If a task fails, the ManagedBuffers
	* underlying each InputStream will be freed by the cleanup() method registered with the
	* TaskCompletionListener. However, callers should close() these InputStreams
	* as soon as they are no longer needed, in order to release memory as early as possible.
	*
	* Throws a FetchFailedException if the next block could not be fetched.
	*/
	override def next(): (BlockId, InputStream) = {


// Send fetch requests up to maxBytesInFlight	
	fetchUpToMaxBytes()
	}

```



### Shuffle Manager

Old strategy: hash based shuffle manager  
Too many random writes and intermidiate files: Each mapper was creating 1 file for each reducer. For example, for 5 mappers and 5 reducers, the hash-based manager was operating on 25 files   


Sort based manager: Mapper puts all partition records to a single file.  
At the beginning, mapper accumulates all records in memory within PartitionedAppendOnlyMap. The records are grouped together by partition. When there are no more space in the memory, records are saved to the disk. In Spark's nomenclature this action is often called spilling.  
Once all records are treated, Spark saves them on disk. It generates 2 files: .data holding records and .index. Data file contains records ordered by their partition. The index file contains the beginning and the end of each stored partition in data file. It defines where given partition starts and ends.  
During reading, reducers use index file to see where records they need are located. Once they know that, they fetch the data and iterate over it to construct expected output.  If files weren't merged during mapping phase, they're merged before iterating in the reading step.  



<img src="https://user-images.githubusercontent.com/16873751/108574529-85e1e180-72cc-11eb-8c36-72114814d3b8.png" alt="spark_arch" width="600"/> 

(from: https://www.waitingforcode.com/apache-spark/shuffle-apache-spark-back-basics/read#shuffle-writing_side)
<br/>


### External shuffle service

<img src="https://user-images.githubusercontent.com/16873751/108574651-d2c5b800-72cc-11eb-94c0-216cffa61e93.png" alt="spark_arch" width="600"/> 

(from: https://www.waitingforcode.com/apache-spark/shuffle-apache-spark-back-basics/read)
<br/>

<img src="https://user-images.githubusercontent.com/16873751/108574657-d6593f00-72cc-11eb-90b7-fef6242db9ca.png" alt="spark_arch" width="600"/> 

(from: https://www.waitingforcode.com/apache-spark/shuffle-apache-spark-back-basics/read)
<br/>


## Chanllege
- Small writes(spindle)
- Connectivity issues between executors and shuffle service:number of connection will be the multiplication of the number of executors (E) by the number of shuffle services (S)
- Scale down


### Cosco from Facebook

[Cosco An Efficient Facebook Scale Shuffle ServiceBrian Cho Facebook,Dmitry Borovsky Facebook](https://www.youtube.com/watch?v=O3z9jDsfnf4)

<img src="https://user-images.githubusercontent.com/16873751/108574045-19b2ae00-72cb-11eb-8aa5-67ab6ca875be.png" alt="cosco" width="600"/> 
<br/>

<img src="https://user-images.githubusercontent.com/16873751/108574097-3fd84e00-72cb-11eb-987b-b89787f8b91f.png" alt="cosco" width="600"/> 
<br/>

<img src="https://user-images.githubusercontent.com/16873751/108574133-5ed6e000-72cb-11eb-82fd-eb801c63483c.png" alt="cosco" width="600"/> 
<br/>


API  

<img src="https://user-images.githubusercontent.com/16873751/108574157-6eeebf80-72cb-11eb-9ead-345353fe96ca.png" alt="cosco" width="600"/> 
<br/>

<img src="https://user-images.githubusercontent.com/16873751/108574183-7ca44500-72cb-11eb-8e4a-120790e0f731.png" alt="cosco" width="600"/> 
<br/>


### Magnet  from LinkedIn

https://engineering.linkedin.com/blog/2020/introducing-magnet

[SFBigAnalytics_20200908: Magnet Shuffle Service: Push-based Shuffle at LinkedIn](https://www.youtube.com/watch?v=D6dzuoAldRw)


Why

<img src="https://user-images.githubusercontent.com/16873751/108653455-532a1b80-747b-11eb-9452-62abb5ef13d6.png" alt="magnet" width="600"/> 
<br/>

- Shuffle service could crash
- Shuffle service is impact by small Ios
- Scale shuffle service(scale up and scale down)


Magnet Overview

- Merge small random reads in shuffle to large sequential reads
- 2 replicated shuffle data
- Locality aware scheduling of reducers

<img src="https://user-images.githubusercontent.com/16873751/108653467-5cb38380-747b-11eb-8e86-8fc67f38c7ce.png" alt="magnet" width="600"/> 
<br/>

- Shuffle service on mapper node will try to **push shuffle blocks** to remote shuffle services
- Push service will also trigger a replica generation

<img src="https://user-images.githubusercontent.com/16873751/108653474-62a96480-747b-11eb-9829-e2643d6a7b4f.png" alt="magnet" width="600"/> 
<br/>


- Reducer has possibility to load data from merged shuffle file
- Reducer can still load unmerged block from original shuffle blocks

<img src="https://user-images.githubusercontent.com/16873751/108653483-689f4580-747b-11eb-8ec7-97adc312b89a.png" alt="magnet" width="600"/> 
<br/>

Magnet benefits
- Improve disk IO
- Cut down RPC needs during shuffle
- Improve shuffle reliability(merged shuffle file be sharded -> original shuffle blocks -> re-run)
