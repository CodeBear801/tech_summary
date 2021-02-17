# Spark Architecture

<img src="https://user-images.githubusercontent.com/16873751/108144204-0bbc1d80-707e-11eb-8794-2b1cdadf3784.png" alt="spark_arch" width="600"/>   

(from: https://docs.cloud.sdu.dk/Apps/spark-cluster.html)
<br/>

## How an application is executed

stand alone mode

<img src="https://user-images.githubusercontent.com/16873751/108144216-1971a300-707e-11eb-95af-25bc04015d74.png" alt="spark_arch" width="400"/> 

(from: https://docs.cloud.sdu.dk/Apps/spark-cluster.html)
<br/>


running on YARN

<img src="https://user-images.githubusercontent.com/16873751/108144309-3efeac80-707e-11eb-9b4f-bce5fd2d8cf9.png" alt="spark_arch" width="400"/> 

(from: https://docs.cloud.sdu.dk/Apps/spark-cluster.html)
<br/>

***
## Terms

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

