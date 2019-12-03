# Structuring Apache Spark 2.0: SQL, DataFrames, Datasets And Streaming
by Michael Armbrust [video](https://www.youtube.com/watch?v=1a4pgYzeFwE&t=54s)

## Why not RDD

### What is a RDD


<img src="resources/imgs/spark_structure_spark_2_0_what_rdd.png" alt="spark_structure_spark_2_0_what_rdd" width="400"/>
<br/>


- when you call rdd, it will return you an iterator
- data itself is serializer to obejct from byte array, or reversely, can't look specific column, change to uncompressed, etc
- optimization is limited

### What Structure
Define **common pattern** of the data analysis, let you describe your computation in that way, like join, select, etc.

### Why Structure


<img src="resources/imgs/spark_structure_spark_2_0_why_structure.png" alt="spark_structure_spark_2_0_why_structure" width="400"/>
<br/>


### DataSet

<img src="resources/imgs/spark_structure_spark_2_0_dataset.png" alt="spark_structure_spark_2_0_dataset" width="400"/>
<br/>


Scala's type safe class or Java's javabins


### DataFrame

<img src="resources/imgs/spark_structure_spark_2_0_dataframe.png" alt="spark_structure_spark_2_0_dataframe" width="400"/>
<br/>


**User don't know the row ahead of time, you don't want to compile the object to hold your data**

<img src="resources/imgs/spark_structure_spark_2_0_python.png" alt="spark_structure_spark_2_0_python" width="400"/>
<br/>

## Structured Spark Internal


<img src="resources/imgs/spark_structure_spark_2_0_columns.png" alt="spark_structure_spark_2_0_columns" width="400"/>
<br/>
<img src="resources/imgs/spark_structure_spark_2_0_functions.png" alt="spark_structure_spark_2_0_functions" width="400"/>
<br/>
<br/>
<img src="resources/imgs/spark_structure_spark_2_0_columns2.png" alt="spark_structure_spark_2_0_columns2" width="600"/>
<br/>
<br/>
<img src="resources/imgs/spark_structure_spark_2_0_columns3.png" alt="spark_structure_spark_2_0_columns3" width="400"/>
<br/>
<br/>
<img src="resources/imgs/spark_structure_spark_2_0_columns4.png" alt="spark_structure_spark_2_0_columns4" width="400"/>

- Hash partition the data over cluster using a shuffle
- Make sure both side have a hash table
- Merge sort will result in the complexity of nlogn
- Spark take a giant dataset and sort it, if the column is the same they will be sort to the **same place**



## Data model
<img src="resources/imgs/spark_structure_spark_2_0_data_model.png" alt="spark_structure_spark_2_0_data_model" width="400"/>
<br/>


<img src="resources/imgs/spark_structure_spark_2_0_data_model2.png" alt="spark_structure_spark_2_0_data_model2" width="400"/>
<br/>

<img src="resources/imgs/spark_structure_spark_2_0_data_model3.png" alt="spark_structure_spark_2_0_data_model3" width="400"/>
<br/>


**Spark knows where the data is and he will try to jump to the location when generating low level code, no need to construct entire object**

## Structured streaming

<img src="resources/imgs/spark_structure_spark_2_0_batch_agg.png" alt="spark_structure_spark_2_0_batch_agg" width="400"/>
<br/>

<img src="resources/imgs/spark_structure_spark_2_0_stream_agg.png" alt="spark_structure_spark_2_0_stream_agg" width="400"/>
<br/>

<img src="resources/imgs/spark_structure_spark_2_0_execution.png" alt="spark_structure_spark_2_0_execution" width="400"/>
<br/>



- Instead scanning file every time, spark will record the location of read index and only scan new data
- stateful aggregation: when aggregation data, will scan data you could see + previous scan result
