# Spark blocks

Overview

<img src="https://user-images.githubusercontent.com/16873751/108654110-d5ffa600-747c-11eb-921c-2e27e676701b.png" alt="spark_arch" width="600"/> 
(from: https://masterwangzx.com/2020/09/07/store-overview/)
<br/>


what is block

```java
  /**
   * Gets or computes an RDD partition. Used by RDD.iterator() when an RDD is cached.
   */
  private[spark] def getOrCompute(partition: Partition, context: TaskContext): Iterator[T] = {
    val blockId = RDDBlockId(id, partition.index)
    var readCachedBlock = true
    // This method is called on executors, so we need call SparkEnv.get instead of sc.env.
    SparkEnv.get.blockManager.getOrElseUpdate(blockId, storageLevel, elementClassTag, () => {
 // ...
```
