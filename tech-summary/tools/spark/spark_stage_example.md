# Spark stage examples

## Word count

<img src="https://user-images.githubusercontent.com/16873751/108278106-63669180-712f-11eb-8dbd-9870418b0498.png" alt="spark_arch" width="600"/>   

(from: https://juejin.cn/post/6844904047011430407#heading-16)
<br/>

Job由saveAsTextFile触发，该Job由RDD-3和saveAsTextFile方法组成，根据RDD之间的依赖关系从RDD-3开始回溯搜索，直到没有依赖的RDD-0，在回溯搜索过程中，RDD-3依赖RDD-2，并且是宽依赖，所以在RDD-2和RDD-3之间划分Stage，RDD-3被划到最后一个Stage，即ResultStage中，RDD-2依赖RDD-1，RDD-1依赖RDD-0，这些依赖都是窄依赖，所以将RDD-0、RDD-1和RDD-2划分到同一个Stage，即ShuffleMapStage中，实际执行的时候，数据记录会一气呵成地执行RDD-0到RDD-2的转化。不难看出，其本质上是一个深度优先搜索算法。


**一个Stage是否被提交，需要判断它的父Stage是否执行，只有在父Stage执行完毕才能提交当前Stage，如果一个Stage没有父Stage，那么从该Stage开始提交**。Stage提交时会将Task信息（分区信息以及方法等）序列化并被打包成TaskSet交给TaskScheduler，一个Partition对应一个Task，另一方面TaskScheduler会监控Stage的运行状态，只有Executor丢失或者Task由于Fetch失败才需要重新提交失败的Stage以调度运行失败的任务，其他类型的Task失败会在TaskScheduler的调度过程中重试。


<img src="https://user-images.githubusercontent.com/16873751/108278236-9c9f0180-712f-11eb-9d4c-e6216a5434c5.png" alt="spark_arch" width="800"/>   

(from: https://www.jianshu.com/p/162f82d93ff4)
<br/>

## Top N number

Req: there are duplicate integers in a file, find top-10 integer which has most counts

```java
scala> val sourceRdd = sc.textFile("/tmp/hive/hive/result",10).repartition(5)
sourceRdd: org.apache.spark.rdd.RDD[String] = MapPartitionsRDD[5] at repartition at <console>:27

scala> val allTopNs = sourceRdd.flatMap(line => line.split(" ")).map(word => (word, 1)).reduceByKey(_+_).repartition(10).sortByKey(ascending = true, 100).map(tup => (tup._2, tup._1)).mapPartitions(
| iter => {
| iter.toList.sortBy(tup => tup._1).takeRight(100).iterator
| }
| ).collect()

scala> val finalTopN = scala.collection.SortedMap.empty[Int, String].++(allTopNs)

scala> finalTopN.takeRight(10).foreach(tup => {println(tup._2 + " occurs times : " + tup._1)})
```

result

```shell
53 occurs times : 1070
147 occurs times : 1072
567 occurs times : 1073
931 occurs times : 1075
267 occurs times : 1077
768 occurs times : 1080
612 occurs times : 1081
877 occurs times : 1082
459 occurs times : 1084
514 occurs times : 1087
```

<img src="https://user-images.githubusercontent.com/16873751/108278647-2f3fa080-7130-11eb-90eb-c827fbe12c40.png" alt="spark_arch" width="800"/>   
<br/>

<img src="https://user-images.githubusercontent.com/16873751/108278789-61e99900-7130-11eb-887b-6a12e533eaad.png" alt="spark_arch" width="800"/>   
<br/>

<img src="https://user-images.githubusercontent.com/16873751/108278814-6e6df180-7130-11eb-9b39-46fa35f61c1b.png" alt="spark_arch" width="800"/>   
<br/>



From stackover flow: https://stackoverflow.com/questions/29849413/spark-rdd-sortbykey-triggering-a-new-job
```
As Sean pointed out in https://www.mail-archive.com/user@spark.apache.org/msg27005.html, "[...]sortByKey actually runs a job to assess the distribution of the data (see JIRA https://issues.apache.org/jira/browse/SPARK-1021)". I hope that this help others when debugging the number of jobs and stages of an aplications.
```

<img src="https://user-images.githubusercontent.com/16873751/108280394-f6ed9180-7132-11eb-981c-32559adb6912.png" alt="spark_arch" width="800"/>   
<br/>

<img src="https://user-images.githubusercontent.com/16873751/108280432-066cda80-7133-11eb-9647-70a3de781a85.png" alt="spark_arch" width="800"/>   

(from: TRANSFORMATIONS AND ACTIONS https://training.databricks.com/visualapi.pdf)
<br/>




([more info](https://www.cnblogs.com/johnny666888/p/11233982.html))