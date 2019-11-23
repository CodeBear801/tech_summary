# A Tale of Three Apache Spark APIs: RDDs, DataFrames, and Datasets 
By Jules Damji 2017 [Video](https://www.youtube.com/watch?v=Ofk7G3GD9jk)

## RDD

### What is RDD

#### 1. Distributed data abstraction 
  - I could write functions and make them run in parallel  

<img src="resources/imgs/spark_rdd_df_ds_2017_rdd_1.png" alt="spark_rdd_df_ds_2017_rdd_1" width="600"/>
<br/>

#### 2. Resilient & Immutable
- I could re-create the RDDs at any time
- RDD's metadata will record lineage

<img src="resources/imgs/spark_rdd_df_ds_2017_rdd_2.png" alt="spark_rdd_df_ds_2017_rdd_2" width="600"/>
<br/>


#### 3. Compile-time type-safe

<img src="resources/imgs/spark_rdd_df_ds_2017_rdd_3.png" alt="spark_rdd_df_ds_2017_rdd_3" width="600"/>
<br/>


#### 4. Data could be unstructured or structured data

- RDD don't understand the data, its just string, upper user could use format to decoding them

<img src="resources/imgs/spark_rdd_df_ds_2017_rdd_4.png" alt="spark_rdd_df_ds_2017_rdd_4" width="600"/>
<br/>


#### 5. Lazy
- actions means the entire DAG start to execution

<img src="resources/imgs/spark_rdd_df_ds_2017_rdd_5.png" alt="spark_rdd_df_ds_2017_rdd_5" width="600"/>
<br/>

### RDD operations

<img src="resources/imgs/spark_rdd_df_ds_2017_rdd_operations.png" alt="spark_rdd_df_ds_2017_rdd_operations" width="600"/>
<br/>

### Why RDDs?

<img src="resources/imgs/spark_rdd_df_ds_2017_rdd_why.png" alt="spark_rdd_df_ds_2017_rdd_why" width="400"/>
<br/>

An example: Pick most popular wiki pages in english language

<img src="resources/imgs/spark_rdd_df_ds_2017_rdd_example.png" alt="spark_rdd_df_ds_2017_rdd_example" width="600"/>
<br/>

### When to use RDD


<img src="resources/imgs/spark_rdd_df_ds_2017_rdd_when.png" alt="spark_rdd_df_ds_2017_rdd_when" width="400"/>
<br/>

### What's the problem of RDD

Spark don't look into lambda functions, and he don't know what's the data


<img src="resources/imgs/spark_rdd_df_ds_2017_rdd_what_problem.png" alt="spark_rdd_df_ds_2017_rdd_what_problem" width="400"/>
<br/>


Inadvertent inefficient example -> reduceByKey first then do filter.  reduceByKey will take all data and shuffering the data

### What's inside RDDs

<img src="resources/imgs/spark_rdd_df_ds_2017_rdd_what_inside.png" alt="spark_rdd_df_ds_2017_rdd_what_inside" width="400"/>
<br/>

Functions with RDD, spark don't know what it is.   
Spark will serialize the code and send which to executor.   


## DataFrames & DataSets

<img src="resources/imgs/spark_rdd_df_ds_2017_df_ds_sql.png" alt="spark_rdd_df_ds_2017_df_ds_sql" width="600"/>
<br/>

- DataFrame-> access non-exists colum, run time error
- Datasets-> similar to java.bin
- In Scala, data frame is dataset
- In java, there is no type, everything is dataset
- In python, there is only data frame

<img src="resources/imgs/spark_rdd_df_ds_2017_df_example.png" alt="spark_rdd_df_ds_2017_df_example" width="600"/>
<br/>

Comments: 
declarative column types
tell spark what to do, let spark to optimize

<img src="resources/imgs/spark_rdd_df_ds_2017_sql_example.png" alt="spark_rdd_df_ds_2017_sql_example" width="600"/>
<br/>
<img src="resources/imgs/spark_rdd_df_ds_2017_df_example_2.png" alt="spark_rdd_df_ds_2017_df_example_2" width="600"/>
<br/>

### Why structure APIs

<img src="resources/imgs/spark_rdd_df_ds_2017_structure_api.png" alt="spark_rdd_df_ds_2017_structure_api" width="600"/>
<br/>

Structure gives the ability to do things in a declative matter, similar to database query

### Catalyst

spark_rdd_df_ds_2017_catalyst
1. Any query(sql, dataframe, dataset) will create unresolved logical plan
2. Then check for `catalog` to see which column you are referring to
3. Then create logical plan and optimized on
4. Create lots of physical plans and ranking them
5. Selected
6. RDD -> everything in spark will go into RDD

<img src="resources/imgs/spark_rdd_df_ds_2017_df_optimization.png" alt="spark_rdd_df_ds_2017_df_optimization" width="600"/>
<br/>



Step 3:   
- scan is operate on database  
- event could coming from parquet file  
- Push filter to lower level which is more efficient on operation  
- Bring less data to spark  

<img src="resources/imgs/spark_rdd_df_ds_2017_df_optimization_2.png" alt="spark_rdd_df_ds_2017_df_optimization_2" width="600"/>
<br/>


### Why & When


<img src="resources/imgs/spark_rdd_df_ds_2017_df_ds_why_when.png" alt="spark_rdd_df_ds_2017_df_ds_why_when" width="600"/>
<br/>

