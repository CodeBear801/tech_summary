# Spark shuffle

## What is Shuffle

<img src="
https://user-images.githubusercontent.com/16873751/108574287-d278ed00-72cb-11eb-961e-f98f62030db7.png" alt="spark_arch" width="600"/> 

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
