# DAG scheduler and Task Execution

For components in spark, please go to [Spark architecture page](./spark_arch.md) especially [here](https://github.com/CodeBear801/tech_summary/blob/master/tech-summary/tools/spark/spark_arch.md#job-stage-task)

For spark stage examples, please go to [Spark stage examples](./spark_stage_example.md)

## DAG scheduler


<img src="https://user-images.githubusercontent.com/16873751/108280936-b8a4a200-7133-11eb-9d17-f224ca7e749f.png" alt="spark_arch" width="600"/>   

(from: https://juejin.cn/post/6844904047011430407#heading-16)
<br/>

Important comments in [DAGScheduler.scala](https://github.com/apache/spark/blob/23d4f6b3935bb6ca3ecb8ce43bd53788d5e16e74/core/src/main/scala/org/apache/spark/scheduler/DAGScheduler.scala#L49)

```java
/**
 * The high-level scheduling layer that implements stage-oriented scheduling. It computes a DAG ofMatei Zaharia, 8 years ago: • Added some comments on threading in scheduler…
 * stages for each job, keeps track of which RDDs and stage outputs are materialized, and finds a
 * minimal schedule to run the job. It then submits stages as TaskSets to an underlying
 * TaskScheduler implementation that runs them on the cluster. A TaskSet contains fully independent
 * tasks that can run right away based on the data that's already on the cluster (e.g. map output
 * files from previous stages), though it may fail if this data becomes unavailable.
 *
 * Spark stages are created by breaking the RDD graph at shuffle boundaries. RDD operations with
 * "narrow" dependencies, like map() and filter(), are pipelined together into one set of tasks
 * in each stage, but operations with shuffle dependencies require multiple stages (one to write a
 * set of map output files, and another to read those files after a barrier). In the end, every
 * stage will have only shuffle dependencies on other stages, and may compute multiple operations
 * inside it. The actual pipelining of these operations happens in the RDD.compute() functions of
 * various RDDs
 *
 * In addition to coming up with a DAG of stages, the DAGScheduler also determines the preferred
 * locations to run each task on, based on the current cache status, and passes these to the
 * low-level TaskScheduler. Furthermore, it handles failures due to shuffle output files being
 * lost, in which case old stages may need to be resubmitted. Failures *within* a stage that are
 * not caused by shuffle file loss are handled by the TaskScheduler, which will retry each task
 * a small number of times before cancelling the whole stage.
 */
```

org.apache.spark.SparkContext#runJob（dagScheduler.runJob） —>
org.apache.spark.scheduler.DAGScheduler#runJob（submitJob） —>
org.apache.spark.scheduler.JobSubmitted$#onReceive（JobSubmitted） —> 
org.apache.spark.scheduler.DAGScheduler#[handleJobSubmitted](https://github.com/apache/spark/blob/23d4f6b3935bb6ca3ecb8ce43bd53788d5e16e74/core/src/main/scala/org/apache/spark/scheduler/DAGScheduler.scala#L1107)
