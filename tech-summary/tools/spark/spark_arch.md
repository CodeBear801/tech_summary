- [Spark Architecture](#spark-architecture)
  - [How an application is executed](#how-an-application-is-executed)
  - [Mode](#mode)
    - [stand alone mode](#stand-alone-mode)
    - [running on YARN](#running-on-yarn)
    - [running on K8S](#running-on-k8s)
  - [Terms](#terms)
    - [Driver](#driver)
      - [SparkSession](#sparksession)
      - [Application Planning](#application-planning)
      - [Application Orchestration](#application-orchestration)
    - [Worker](#worker)
    - [Executor](#executor)
    - [Spark Master](#spark-master)
    - [Cluster Manager](#cluster-manager)
  - [More details](#more-details)
    - [Job, stage, task](#job-stage-task)

# Spark Architecture

<img src="https://user-images.githubusercontent.com/16873751/108144204-0bbc1d80-707e-11eb-8794-2b1cdadf3784.png" alt="spark_arch" width="600"/>   

(from: https://docs.cloud.sdu.dk/Apps/spark-cluster.html)
<br/>

***
## How an application is executed


<img src="https://user-images.githubusercontent.com/16873751/108236879-3baa0600-70fc-11eb-8fc0-0870fa38c613.png" alt="spark_task_commit" width="600"/> 

(from: 深入理解Spark核心思想与源码分析 Chapter 5)
<br/><br/><br/>


<img src="https://user-images.githubusercontent.com/16873751/108428807-fe28a400-71f3-11eb-9a20-e354f8280caf.png" alt="spark_task_commit" width="800"/> 
<br/>

## Mode

### stand alone mode

<img src="https://user-images.githubusercontent.com/16873751/108144216-1971a300-707e-11eb-95af-25bc04015d74.png" alt="spark_arch" width="400"/> 

(from: https://docs.cloud.sdu.dk/Apps/spark-cluster.html)
<br/>


### running on YARN

<img src="https://user-images.githubusercontent.com/16873751/108144309-3efeac80-707e-11eb-9b4f-bce5fd2d8cf9.png" alt="spark_arch" width="400"/> 

(from: https://docs.cloud.sdu.dk/Apps/spark-cluster.html)
<br/>

### running on K8S

<img src="https://user-images.githubusercontent.com/16873751/108259552-14603280-7116-11eb-827e-480f05671987.png" alt="spark_arch" width="600"/> 

(from: https://www.youtube.com/watch?v=3EbTr79wLkU)
<br/>

<img src="https://user-images.githubusercontent.com/16873751/108259591-1cb86d80-7116-11eb-908a-d4f78c1e0df9.png" alt="spark_arch" width="600"/> 

(from: https://www.youtube.com/watch?v=3EbTr79wLkU)
<br/>

<img src="https://user-images.githubusercontent.com/16873751/108259598-1fb35e00-7116-11eb-9e03-20db25a9f79c.png" alt="spark_arch" width="600"/> 

(from: https://www.youtube.com/watch?v=3EbTr79wLkU)
<br/>

- spark-submit: can't specify node-selector
- spark-operator is an kubernetes' object, you could config which in yaml
- more notes: https://my.oschina.net/u/4287583/blog/4443764

***
## Terms


<img src="https://user-images.githubusercontent.com/16873751/108428027-d258ee80-71f2-11eb-9413-6a882443c9d8.png" alt="spark_arch" width="800"/> 

(from: https://spark.apache.org/docs/1.1.0/cluster-overview.html)
<br/>

### Driver
The life of a Spark application starts and finishes with the Spark Driver.    
The Driver is the process that clients use to submit applications in Spark.   
The Driver is also responsible for planning and coordinating the execution of the Spark program and returning status and/or results (data) to the client. The Driver can physically reside on a client or on a node in the cluster   

#### SparkSession 

The Spark Driver is responsible for creating the SparkSession. The SparkSession object represents a connection to a Spark cluster

#### Application Planning

One of the main functions of the Driver is to plan the application. The Driver takes the application processing input and plans the execution of the program. The Driver takes all the requested transformations (data manipulation operations) and actions (requests for output or prompts to execute programs) and creates a directed acyclic graph (DAG) of nodes, each representing a transformational or computational step.  

A Spark application DAG consists of tasks and stages. A task is the smallest unit of schedulable work in a Spark program. A stage is a set of tasks that can be run together. Stages are dependent upon one another; in other words, there are stage dependencies.

#### Application Orchestration

- Keeping track of available resources to execute tasks
- Scheduling tasks to run “close” to the data where possible

### Worker

A Worker node—which hosts the Executor process

### Executor

Spark Executors are the processes on which Spark DAG tasks run. Executors reserve CPU and memory resources on slave nodes, or Workers, in a Spark cluster. An Executor is dedicated to a specific Spark application and terminated when the application completes. A Spark program normally consists of many Executors, often working in parallel.  

### Spark Master 

The Spark Master is the process that requests resources in the cluster and makes them available to the Spark Driver.

### Cluster Manager 

The Cluster Manager is the process responsible for monitoring the Worker nodes and reserving resources on these nodes upon request by the Master. The Master then makes these cluster resources available to the Driver in the form of Executors.

***
## More details

### Job, stage, task

<img src="https://user-images.githubusercontent.com/16873751/108260091-be3fbf00-7116-11eb-9387-75239b87cae1.png" alt="spark_arch" width="600"/> 

(from: https://juejin.cn/post/6844904047011430407#heading-16)
<br/>

- Job是以Action方法为界，遇到一个Action方法则触发一个Job
- Stage是Job的子集，以RDD宽依赖(即Shuffle)为界，遇到Shuffle做一次划分(groupByKey,reduceByKey,countByKey)
- Task是Stage的子集，以并行度(分区数)来衡量，分区数是多少，则有多少个task

<img src="https://user-images.githubusercontent.com/16873751/108260293-0363f100-7117-11eb-945c-4d516cd412f5.png" alt="spark_arch" width="400"/> 

(from: https://www.usenix.org/system/files/conference/nsdi12/nsdi12-final138.pdf)
<br/>

Spark RDD通过其Transactions操作，形成了RDD血缘关系图，即DAG，最后通过Action的调用，触发Job并调度执行。DAGScheduler负责Stage级的调度，主要是将job切分成若干Stages，并将每个Stage打包成TaskSet交给TaskScheduler调度。TaskScheduler负责Task级的调度，将DAGScheduler给过来的TaskSet按照指定的调度策略分发到Executor上执行，调度过程中SchedulerBackend负责提供可用资源，其中SchedulerBackend有多种实现，分别对接不同的资源管理系统

