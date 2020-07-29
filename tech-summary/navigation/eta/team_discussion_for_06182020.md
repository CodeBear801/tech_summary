
resources for [discussion](https://github.com/Telenav/osrm-backend/issues/356#issuecomment-646378400)

## Why XGBoost
- https://github.com/dmlc/xgboost
- http://datascience.la/xgboost-workshop-and-meetup-talk-with-tianqi-chen/

<img src="https://user-images.githubusercontent.com/16873751/85156321-c1dd1800-b20e-11ea-8d2a-a4b1cb908f3b.png" alt="drawing" width="600"/><br/>

decision tree, weak learner(individually they are quite inaccurate, but slightly better when work together)

<img src="https://user-images.githubusercontent.com/16873751/85156771-63646980-b20f-11ea-83e9-dcb2341a41cc.png" alt="drawing" width="600"/><br/>

second tree must provide positive effort when combine with first tree


## PCA

- https://setosa.io/ev/principal-component-analysis/

<img src="https://user-images.githubusercontent.com/16873751/85156926-9c9cd980-b20f-11ea-859a-5da8e06157a5.png" alt="drawing" width="800"/><br/>

- Dimension for original data is reasonable for human beings, but might not friendly for decision tree

- Try to represent data with different `coordinate system`

- Make the dimension the principal component with most variation - (Minimize difference between min value and max value)

## Why Parquet

- Dremel paper https://storage.googleapis.com/pub-tools-public-publication-data/pdf/36632.pdf
- Google Dremel 原理 - 如何能3秒分析1PB https://www.twblogs.net/a/5b82787e2b717766a1e868d0
- Approaching NoSQL Design in DynamoDB https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/bp-general-nosql-design.html#bp-general-nosql-design-approach
- Spark + Parquet In Depth https://www.youtube.com/watch?v=_0Wpwj_gvzg&t=1501s
- https://medium.com/@rajnishtiwari2010/conversion-of-json-to-parquet-format-using-apache-parquet-in-java-b694a0a7487d

### Target of column based data organize
- Only load needed data
- Organize data storage based on needs(query, filter)

Sample data input   
<img src="https://user-images.githubusercontent.com/16873751/85159671-ac69ed00-b212-11ea-9f06-0623b2a568da.png" alt="drawing" width="600"/><br/>

Flat data vs. Nested data

<img src="https://user-images.githubusercontent.com/16873751/85159736-bb509f80-b212-11ea-877d-e6b7c68004f6.png" alt="drawing" width="600"/><br/>


Option on Parquet with `scala`
```java
val flatDF = sc.read.option("delimiter", "\t")
                    .option("header", "true")
                    .csv(flatInput)
                    .rdd
                    .map(r => transformRow(r))
                    .toDF

flatDF.write.option("compression", "snappy")
            .parquet(flatOutput)

var nestedDF = sc.read.json(nestedInput)

nestedDF.write.option("compression", "snappy")
              .parquet(nestedOutput)
```

### How to avoid reading un-related column?


<img src="https://user-images.githubusercontent.com/16873751/85160645-72e5b180-b213-11ea-9d54-3a13eff17aef.png" alt="drawing" width="600"/><br/>

By record data column oriented, we could use different way to compress the data  
Incrementally(record diff), or use dictionary  

More details could be found in [parquet encoding](https://github.com/apache/parquet-format/blob/master/Encodings.md)

### How to save space

<img src="https://user-images.githubusercontent.com/16873751/85160785-a88a9a80-b213-11ea-9a31-59ac44708bae.png" alt="drawing" width="600"/><br/>

<img src="https://user-images.githubusercontent.com/16873751/85160804-afb1a880-b213-11ea-842e-b1198150ac43.png" alt="drawing" width="600"/><br/>

### Parquet file internal

tree structure

![image](https://user-images.githubusercontent.com/16873751/85160861-ca841d00-b213-11ea-83d6-77a347f35734.png)

[parquet file format](https://github.com/apache/parquet-format)  

- file metadata: schema, num of rows
- Row group: each time when write massive data, we are not going to write them all together but piece by piece
- Column Chunk: consider each column individually(if whole column is no, directly skip)
- Page header: size
- Page: record meta data like count, max, min to help quick filter

### The power of partition
```java
dataFrame.write
         .partitionBy("Year", "Month", "Day", "Hour")
         .parquet(outputFile)
```

### Understanding the case in paper
https://github.com/CodeBear801/tech_summary/blob/master/tech-summary/papers/dremel.md#understanding-the-case-in-paper

### Experiment

https://github.com/apache/parquet-mr

#### set up env
- via source code
   +  https://github.com/apache/parquet-mr/tree/master/parquet-tools
   + https://repo1.maven.org/maven2/org/apache/parquet/parquet-tools/1.11.0/parquet-tools-1.11.0.jar

- via docker
   + https://github.com/NathanHowell/parquet-tools/blob/master/Dockerfile
```
docker pull nathanhowell/parquet-tools:latest
```

#### Experiment via docker
- https://github.com/apache/drill/blob/master/exec/java-exec/src/test/resources/lateraljoin/nested-customer.json
- https://github.com/apache/drill/blob/master/exec/java-exec/src/test/resources/lateraljoin/nested-customer.parquet

```
wget https://github.com/apache/drill/blob/master/exec/java-exec/src/test/resources/lateraljoin/nested-customer.parquet

docker run -it --rm nathanhowell/parquet-tools:latest --help

docker run --rm -it -v /yourlocalpath:/test nathanhowell/parquet-tools:latest schema test/nested-customer.parquet
```

```
➜  tmp docker run --rm -it -v /Users/xunliu/Downloads/tmp:/test nathanhowell/parquet-tools:latest schema test/nested-customer.parquet
message root {
  optional binary _id (UTF8);
  optional binary c_address (UTF8);
  optional double c_id;
  optional binary c_name (UTF8);
  repeated group orders {
    repeated group items {
      optional binary i_name (UTF8);
      optional double i_number;
      optional binary i_supplier (UTF8);
    }
    optional double o_amount;
    optional double o_id;
    optional binary o_shop (UTF8);
  }
}
```


- For how to generate parquet file from csv, I think the best ways is using `Apache Spark`, `Apache Drill` or other cloud platforms, such as [Databricks notbook](https://docs.databricks.com/data/data-sources/read-parquet.html)
- I failed on following command line tools
   + golang https://github.com/xitongsys/parquet-go failed to build
   + python 2.7 https://github.com/redsymbol/csv2parquet failed to build


## Why apache dataframe

- why rdd https://github.com/CodeBear801/tech_summary/blob/master/tech-summary/papers/rdd.md
- dataframe vs rdd https://github.com/CodeBear801/tech_summary/blob/master/tech-summary/lectures/spark_rdd_dataframe_dataset.md
- https://github.com/CodeBear801/tech_summary/blob/master/tech-summary/papers/flumejava.md
- http://why-not-learn-something.blogspot.com/2016/07/apache-spark-rdd-vs-dataframe-vs-dataset.html

### Why resilient distribute data

map reduce flow
<img src="https://user-images.githubusercontent.com/16873751/85180850-56f50680-b239-11ea-95f3-05ebea4329ca.png" alt="drawing" width="600"/><br/>

Let's say there are multiple stage of map reduce
- how to represent distribute data for programming language
   + [google file system](https://github.com/CodeBear801/tech_summary/blob/master/tech-summary/papers/gfs.md)  
<img src="https://user-images.githubusercontent.com/16873751/85181162-1944ad80-b23a-11ea-934e-efce7c5610e9.png" alt="drawing" width="400"/><br/>

- what if middle step failed, how to recover
- how to let programmer easy to write mr program
- let's say we want to first add 1 on all numbers then filter odd numbers, can we optimize calculation?


RDD means Resilient Distributed Datasets, an RDD is a collection of partitions of records.
```
The main challenge in designing RDDs is defining a programming interface 
that can provide fault tolerance efficiently. 
```

What is RDD
```
每个RDD都包含：
（1）一组RDD分区（partition，即数据集的原子组成部分）；
（2）对父RDD的一组依赖，这些依赖描述了RDD的Lineage；
（3）一个函数，即在父RDD上执行何种计算；
（4）元数据，描述分区模式和数据存放的位置。
例如，一个表示HDFS文件的RDD包含：各个数据块的一个分区，并知道各个数据块放在哪些节点上。
而且这个RDD上的map操作结果也具有同样的分区，map函数是在父数据上执行的。
```

Example code

```java
val rdd = sc.textFile("/mnt/wikipediapagecounts.gz")
var parsedRDD = rdd.flatMap{
    line => line.split("""\s+""") match {
        case Array(project, page, numRequests,-)=>Some((project, page, numRequests))
        case _=None
    }
}

// filter only english pages; count pages and requests to it
parsedRDD.filter{case(project, page, numRequests) => project == "en"}
         .map{ case(_, page, numRequests) => (page, numRequests)}
         .reduceByKey(_+_)
         .take(100)
         .foreach{case (page, requests) => println(s"$page:$requests")}
```

why not rdd

<img src="https://user-images.githubusercontent.com/16873751/85183302-12209e00-b240-11ea-9320-315dcc40ec12.png" alt="drawing" width="400"/><br/>

**Spark don't look into lambda functions, and he don't know what's the data/type**

<img src="https://user-images.githubusercontent.com/16873751/85183330-32e8f380-b240-11ea-930f-1de4bf6de5d5.png" alt="drawing" width="400"/><br/>

DataFrame Sample

```java
// convert RDD -> DF with colum names
val df = parsedRDD.toDF("project", "page", "numRequests")
// filter, groupBy, sum, and then agg()
df.filter($"project" === "en")
  .groupBy($"page")
  .agg(sum($"numRequests").as("count"))
  .limit(100)
  .show(100)
```

project | page | numRequests
---|---|---
en | 23 | 45
en | 24 | 200


<img src="https://user-images.githubusercontent.com/16873751/85183481-c02c4800-b240-11ea-8f2e-2778368988c8.png" alt="drawing" width="600"/><br/>

<img src="https://user-images.githubusercontent.com/16873751/85183489-c3bfcf00-b240-11ea-9ccf-e714153a0b1f.png" alt="drawing" width="600"/><br/>

<img src="https://user-images.githubusercontent.com/16873751/85183495-c6babf80-b240-11ea-9a99-045b364bbbcb.png" alt="drawing" width="600"/><br/>


### Experiment
- https://github.com/CodeBear801/tech_summary/blob/master/tech-summary/tools/spark/spark_docker_setup.md

#### More info
- https://github.com/CodeBear801/tech_summary/blob/master/tech-summary/papers/flumejava.md
- [pcollection in apache beam](https://beam.apache.org/documentation/programming-guide/#pcollections) 



