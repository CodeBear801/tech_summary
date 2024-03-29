# DAGScheduler and TaskScheduler

For components in spark, please go to [Spark architecture page](./spark_arch.md) especially [here](https://github.com/CodeBear801/tech_summary/blob/master/tech-summary/tools/spark/spark_arch.md#job-stage-task)

For spark stage examples, please go to [Spark stage examples](./spark_stage_example.md)


<img src="https://user-images.githubusercontent.com/16873751/108236879-3baa0600-70fc-11eb-8fc0-0870fa38c613.png" alt="spark_task_commit" width="600"/> 

(from: 深入理解Spark核心思想与源码分析 Chapter 5)
<br/>


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


```java
/**
Summary
- Go reversely from final stage
- create stages based on shuffle stage
- use recursion and submit parent stage with higher priority

*/

 private[scheduler] def handleJobSubmitted(jobId: Int,
      finalRDD: RDD[_],
      func: (TaskContext, Iterator[_]) => _,
      partitions: Array[Int],
      callSite: CallSite,
      listener: JobListener,
      properties: Properties): Unit = {
        
      // New stage creation may throw an exception if, for example, jobs are run on a
      // HadoopRDD whose underlying HDFS files have been deleted.
      // [Perry] Step 1. create finalStage based on finalRDD, add which into DAGScheduler's memory cache
      finalStage = createResultStage(finalRDD, func, partitions, jobId, callSite)

      // [Perry] Step 2. create a job based on final stage rdd
      val job = new ActiveJob(jobId, finalStage, callSite, listener, properties)


      val jobSubmissionTime = clock.getTimeMillis()
      // [Perry] Step 3. add job into memory cache, submit final stage
      //         which will trigger other stage be put into waitingStages queue
      jobIdToActiveJob(jobId) = job
      activeJobs += job
      finalStage.setActiveJob(job)
      val stageIds = jobIdToStageIds(jobId).toArray
      val stageInfos = stageIds.flatMap(id => stageIdToStage.get(id).map(_.latestInfo))
      listenerBus.post(
      SparkListenerJobStart(job.jobId, jobSubmissionTime, stageInfos, properties))
      submitStage(finalStage)

      }


  /** Submits stage, but first recursively submits any missing parents. */
  private def submitStage(stage: Stage): Unit = {
    val jobId = activeJobForStage(stage)
    if (jobId.isDefined) {
      logDebug(s"submitStage($stage (name=${stage.name};" +
        s"jobs=${stage.jobIds.toSeq.sorted.mkString(",")}))")
      if (!waitingStages(stage) && !runningStages(stage) && !failedStages(stage)) {
        val missing = getMissingParentStages(stage).sortBy(_.id)
        logDebug("missing: " + missing)
        if (missing.isEmpty) {
          logInfo("Submitting " + stage + " (" + stage.rdd + "), which has no missing parents")
          submitMissingTasks(stage, jobId.get)
        } else {
          //[Perry] Key point is here: recursion, add all missing parent stages
          //        stop condition is in the upper if: when missing.isEmpty, like stage0
          //        submitMissingTasks
          for (parent <- missing) {
            submitStage(parent)
          }
          waitingStages += stage
        }
      }
    } else {
      abortStage(stage, "No active job for stage " + stage.id, None)
    }
  }

  }


  /** Called when stage's parents are available and we can now do its task. */
  private def submitMissingTasks(stage: Stage, jobId: Int): Unit = {

// [Perry] Partition amount is related with task amount
    // Figure out the indexes of partition ids to compute.
    val partitionsToCompute: Seq[Int] = stage.findMissingPartitions()

// [Perry] find best location to apply a task, very important
    val taskIdToLocations: Map[Int, Seq[TaskLocation]] = try {
      stage match {
        case s: ShuffleMapStage =>
          partitionsToCompute.map { id => (id, getPreferredLocs(stage.rdd, id))}.toMap
        case s: ResultStage =>
          partitionsToCompute.map { id =>
            val p = s.partitions(id)
            (id, getPreferredLocs(stage.rdd, p))
          }.toMap
      }
    } catch {
    }
  }

  val tasks: Seq[Task[_]] = try {
      val serializedTaskMetrics = closureSerializer.serialize(stage.latestInfo.taskMetrics).array()
      stage match {
        case stage: ShuffleMapStage =>
          stage.pendingPartitions.clear()
          partitionsToCompute.map { id =>
            val locs = taskIdToLocations(id)
            val part = partitions(id)
            stage.pendingPartitions += id
            new ShuffleMapTask(stage.id, stage.latestInfo.attemptNumber,
              taskBinary, part, locs, properties, serializedTaskMetrics, Option(jobId),
              Option(sc.applicationId), sc.applicationAttemptId, stage.rdd.isBarrier())
          }

        case stage: ResultStage =>
          partitionsToCompute.map { id =>
            val p: Int = stage.partitions(id)
            val part = partitions(p)
            val locs = taskIdToLocations(id)
            new ResultTask(stage.id, stage.latestInfo.attemptNumber,
              taskBinary, part, locs, id, properties, serializedTaskMetrics,
              Option(jobId), Option(sc.applicationId), sc.applicationAttemptId,
              stage.rdd.isBarrier())
          }
      }
    }


      if (tasks.nonEmpty) {
      logInfo(s"Submitting ${tasks.size} missing tasks from $stage (${stage.rdd}) (first 15 " +
        s"tasks are for partitions ${tasks.take(15).map(_.partitionId)})")
//[Perry] Create TaskSet object, calling submitTasks() from TaskScheduler to submit taskSet
      taskScheduler.submitTasks(new TaskSet(
        tasks.toArray, stage.id, stage.latestInfo.attemptNumber, jobId, properties,
        stage.resourceProfileId))
    } 


  /**
   * Recursive implementation for getPreferredLocs.
   *
   * This method is thread-safe because it only accesses DAGScheduler state through thread-safe
   * methods (getCacheLocs()); please be careful when modifying this method, because any new
   * DAGScheduler state accessed by it may require additional synchronization.
   */
//[Perry] From last RDD of the stage to find rdd's partition, check whether its be cached or marked 
//        with checkpoint
//        If found, the best location for that task is the partition cached with checkpoint, because
//        if we applied calculation there we don't need to recalculate previous rdd
  private def getPreferredLocsInternal(
      rdd: RDD[_],
      partition: Int,
      visited: HashSet[(RDD[_], Int)]): Seq[TaskLocation] = {
    // If the partition has already been visited, no need to re-visit.
    // This avoids exponential path exploration.  SPARK-695
    if (!visited.add((rdd, partition))) {
      // Nil has already been returned for previously visited partitions.
      return Nil
    }
// [Perry] is cached
    // If the partition is cached, return the cache locations
    val cached = getCacheLocs(rdd)(partition)
    if (cached.nonEmpty) {
      return cached
    }

//[Perry] check wether current RDD's partition is checkpoint
    // If the RDD has some placement preferences (as is the case for input RDDs), get those
    val rddPrefs = rdd.preferredLocations(rdd.partitions(partition)).toList
    if (rddPrefs.nonEmpty) {
      return rddPrefs.map(TaskLocation(_))
    }

//[Perry] recursion: check parent rdd whether its cached or checkpoint
    // If the RDD has narrow dependencies, pick the first partition of the first narrow dependency
    // that has any placement preferences. Ideally we would choose based on transfer sizes,
    // but this will do for now.
    rdd.dependencies.foreach {
      case n: NarrowDependency[_] =>
        for (inPart <- n.getParents(partition)) {
          val locs = getPreferredLocsInternal(n.rdd, inPart, visited)
          if (locs != Nil) {
            return locs
          }
        }

      case _ =>
    }

//[Perry] for this stage, from last rdd to first rdd, there is no cache or checkpoint for this partition
// then mark PreferredLocs with NIL
    Nil
  }
```

More important code for `stage`

```java

private[scheduler] def handleJobSubmitted(
){
    logInfo("Missing parents: " + getMissingParentStages(finalStage))
}


  private def getMissingParentStages(stage: Stage): List[Stage] = {
    val missing = new HashSet[Stage]
    val visited = new HashSet[RDD[_]]
    // We are manually maintaining a stack here to prevent StackOverflowError
    // caused by recursively visiting
    val waitingForVisit = new ListBuffer[RDD[_]]
    waitingForVisit += stage.rdd
    def visit(rdd: RDD[_]): Unit = {
      if (!visited(rdd)) {
        visited += rdd
        val rddHasUncachedPartitions = getCacheLocs(rdd).contains(Nil)
        if (rddHasUncachedPartitions) {
          for (dep <- rdd.dependencies) {
            dep match {
              case shufDep: ShuffleDependency[_, _, _] =>
                val mapStage = getOrCreateShuffleMapStage(shufDep, stage.firstJobId)
                if (!mapStage.isAvailable) {
                  missing += mapStage
                }
              case narrowDep: NarrowDependency[_] =>
                waitingForVisit.prepend(narrowDep.rdd)
            }
          }
        }
      }
    }


    while (waitingForVisit.nonEmpty) {
      visit(waitingForVisit.remove(0))
    }
    missing.toList
  }

  /**
   * Gets a shuffle map stage if one exists in shuffleIdToMapStage. Otherwise, if the
   * shuffle map stage doesn't already exist, this method will create the shuffle map stage in
   * addition to any missing ancestor shuffle map stages.
[Perry] When there is no wide transform or shuffle stage, there is no new stage need to be created
        otherwise, new stage need to be created
   */
  private def getOrCreateShuffleMapStage(
      shuffleDep: ShuffleDependency[_, _, _],
      firstJobId: Int): ShuffleMapStage = {
    shuffleIdToMapStage.get(shuffleDep.shuffleId) match {
      case Some(stage) =>
        stage

      case None =>
        // Create stages for all missing ancestor shuffle dependencies.
        getMissingAncestorShuffleDependencies(shuffleDep.rdd).foreach { dep =>
          // Even though getMissingAncestorShuffleDependencies only returns shuffle dependencies
          // that were not already in shuffleIdToMapStage, it's possible that by the time we
          // get to a particular dependency in the foreach loop, it's been added to
          // shuffleIdToMapStage by the stage creation process for an earlier dependency. See
          // SPARK-13902 for more information.
          if (!shuffleIdToMapStage.contains(dep.shuffleId)) {
            createShuffleMapStage(dep, firstJobId)
          }
        }
        // Finally, create a stage for the given shuffle dependency.
        createShuffleMapStage(shuffleDep, firstJobId)
    }
  }
```

## Task Scheduler

<img src="https://user-images.githubusercontent.com/16873751/108285957-1093d680-713d-11eb-96a2-16d9e4d0e43e.png" alt="spark_arch" width="600"/>   

(from: https://juejin.cn/post/6844904047011430407#heading-16)
<br/>

<img src="https://user-images.githubusercontent.com/16873751/108397874-4bdde600-71cd-11eb-9d8f-d12298c4e732.png" alt="spark_arch" width="600"/>   

(from: https://juejin.cn/post/6844904047011430407#heading-16)
<br/>


org.apache.spark.scheduler.TaskSchedulerImpl#submitTasks ->  
submitTasks ->backend.reviveOffers() ->   
org.apache.spark.scheduler.cluster.CoarseGrainedClusterMessages.ReviveOffers ->  
makeOffers()->  
[launchTasks()](https://github.com/apache/spark/blob/23d4f6b3935bb6ca3ecb8ce43bd53788d5e16e74/core/src/main/scala/org/apache/spark/scheduler/TaskSchedulerImpl.scala#L234)


```java
// https://github.com/apache/spark/blob/23d4f6b3935bb6ca3ecb8ce43bd53788d5e16e74/core/src/main/scala/org/apache/spark/scheduler/TaskSchedulerImpl.scala#L234

  override def submitTasks(taskSet: TaskSet): Unit = {
    val tasks = taskSet.tasks
    logInfo("Adding task set " + taskSet.id + " with " + tasks.length + " tasks "
      + "resource profile " + taskSet.resourceProfileId)
    this.synchronized {
// [Perry] Each task has a manager
//         when task fails will restart task until limits
//         and will use delay-schedule to handle task's local scheduling
      val manager = createTaskSetManager(taskSet, maxTaskFailures)
      val stage = taskSet.stageId
      val stageTaskSets =
        taskSetsByStageIdAndAttempt.getOrElseUpdate(stage, new HashMap[Int, TaskSetManager])

      stageTaskSets.foreach { case (_, ts) =>
        ts.isZombie = true
      }
      stageTaskSets(taskSet.stageAttemptId) = manager
      schedulableBuilder.addTaskSetManager(manager, manager.taskSet.properties)

    backend.reviveOffers()
  }



  /**
  https://github.com/apache/spark/blob/23d4f6b3935bb6ca3ecb8ce43bd53788d5e16e74/core/src/main/scala/org/apache/spark/scheduler/TaskSchedulerImpl.scala#L492
[Perry] Distribute tasks to executors
   * Called by cluster manager to offer resources on workers. We respond by asking our active task
   * sets for tasks in order of priority. We fill each node with tasks in a round-robin manner so
   * that tasks are balanced across the cluster.
   */
  def resourceOffers(
      offers: IndexedSeq[WorkerOffer],
      isAllFreeResources: Boolean = true): Seq[Seq[TaskDescription]] = synchronized {
      }


/**
https://github.com/apache/spark/blob/5248ecb5ab0c9fdbd5f333e751d1f5fb3c514d14/core/src/main/scala/org/apache/spark/scheduler/cluster/CoarseGrainedSchedulerBackend.scala#L365


*/

    // Launch tasks returned by a set of resource offers
    private def launchTasks(tasks: Seq[Seq[TaskDescription]]): Unit = {
      for (task <- tasks.flatten) {
// [Perry] Serialize the task
        val serializedTask = TaskDescription.encode(task)
        if (serializedTask.limit() >= maxRpcMessageSize) {
          Option(scheduler.taskIdToTaskSetManager.get(task.taskId)).foreach { taskSetMgr =>
            try {
              var msg = "Serialized task %s:%d was %d bytes, which exceeds max allowed: " +
                s"${RPC_MESSAGE_MAX_SIZE.key} (%d bytes). Consider increasing " +
                s"${RPC_MESSAGE_MAX_SIZE.key} or using broadcast variables for large values."
              msg = msg.format(task.taskId, task.index, serializedTask.limit(), maxRpcMessageSize)
              taskSetMgr.abort(msg)
            } catch {
              case e: Exception => logError("Exception in error callback", e)
            }
          }
        }
        else {
// [Perry] Find related executor
          val executorData = executorDataMap(task.executorId)
          // Do resources allocation here. The allocated resources will get released after the task
          // finishes.
          val rpId = executorData.resourceProfileId
          val prof = scheduler.sc.resourceProfileManager.resourceProfileFromId(rpId)
          val taskCpus = ResourceProfile.getTaskCpusOrDefaultForProfile(prof, conf)
// [Perry] Adjust resources in executor
          executorData.freeCores -= taskCpus
          task.resources.foreach { case (rName, rInfo) =>
            assert(executorData.resourcesInfo.contains(rName))
            executorData.resourcesInfo(rName).acquire(rInfo.addresses)
          }

          logDebug(s"Launching task ${task.taskId} on executor id: ${task.executorId} hostname: " +
            s"${executorData.executorHost}.")

//[Perry] Send launchTask to executor
          executorData.executorEndpoint.send(LaunchTask(new SerializableBuffer(serializedTask)))
        }
      }
    }


/**
https://github.com/apache/spark/blob/23d4f6b3935bb6ca3ecb8ce43bd53788d5e16e74/core/src/main/scala/org/apache/spark/scheduler/TaskSchedulerImpl.scala#L492
*/
  /**
   * Called by cluster manager to offer resources on workers. We respond by asking our active task
   * sets for tasks in order of priority. We fill each node with tasks in a round-robin manner so
   * that tasks are balanced across the cluster.
   */
  def resourceOffers(
      offers: IndexedSeq[WorkerOffer],
      isAllFreeResources: Boolean = true): Seq[Seq[TaskDescription]] = synchronized {


    // Take each TaskSet in our scheduling order, and then offer it to each node in increasing order
    // of locality levels so that it gets a chance to launch local tasks on all of them.
    // NOTE: the preferredLocality order: PROCESS_LOCAL, NODE_LOCAL, NO_PREF, RACK_LOCAL, ANY
    for (taskSet <- sortedTaskSets) {

      // Skip the barrier taskSet if the available slots are less than the number of pending tasks.
      if (taskSet.isBarrier && numBarrierSlotsAvailable < taskSet.numTasks) {
        // Skip the launch process.

      } else {
        var launchedAnyTask = false
        var noDelaySchedulingRejects = true
        var globalMinLocality: Option[TaskLocality] = None
        // Record all the executor IDs assigned barrier tasks on.
        val addressesWithDescs = ArrayBuffer[(String, TaskDescription)]()
        for (currentMaxLocality <- taskSet.myLocalityLevels) {
          var launchedTaskAtCurrentMaxLocality = false
          do {
// [Perry] Most important part: 
//         Try to launch task with CurrentMaxLocality
//         if not, then try with next level
            val (noDelayScheduleReject, minLocality) = resourceOfferSingleTaskSet(
              taskSet, currentMaxLocality, shuffledOffers, availableCpus,
              availableResources, tasks, addressesWithDescs)
            launchedTaskAtCurrentMaxLocality = minLocality.isDefined
            launchedAnyTask |= launchedTaskAtCurrentMaxLocality
            noDelaySchedulingRejects &= noDelayScheduleReject
            globalMinLocality = minTaskLocality(globalMinLocality, minLocality)
          } while (launchedTaskAtCurrentMaxLocality)
        }



/**
https://github.com/apache/spark/blob/23d4f6b3935bb6ca3ecb8ce43bd53788d5e16e74/core/src/main/scala/org/apache/spark/scheduler/TaskSchedulerImpl.scala#L370
*/

  /**
   * Offers resources to a single [[TaskSetManager]] at a given max allowed [[TaskLocality]].
   *
   * @param taskSet task set manager to offer resources to
   * @param maxLocality max locality to allow when scheduling
   * @param shuffledOffers shuffled resource offers to use for scheduling,
   *                       remaining resources are tracked by below fields as tasks are scheduled
   * @param availableCpus  remaining cpus per offer,
   *                       value at index 'i' corresponds to shuffledOffers[i]
   * @param availableResources remaining resources per offer,
   *                           value at index 'i' corresponds to shuffledOffers[i]
   * @param tasks tasks scheduled per offer, value at index 'i' corresponds to shuffledOffers[i]
   * @param addressesWithDescs tasks scheduler per host:port, used for barrier tasks
   * @return tuple of (no delay schedule rejects?, option of min locality of launched task)
   */
  private def resourceOfferSingleTaskSet(
      taskSet: TaskSetManager,
      maxLocality: TaskLocality,
      shuffledOffers: Seq[WorkerOffer],
      availableCpus: Array[Int],
      availableResources: Array[Map[String, Buffer[String]]],
      tasks: IndexedSeq[ArrayBuffer[TaskDescription]],
      addressesWithDescs: ArrayBuffer[(String, TaskDescription)])
    : (Boolean, Option[TaskLocality]) = {
```



