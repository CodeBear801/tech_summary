# Easy, Scalable, Fault Tolerant Stream Processing with Structured Streaming in Apache Spark

[Spark Summit 2017 session](https://databricks.com/session/easy-scalable-fault-tolerant-stream-processing-with-structured-streaming-in-apache-spark)


Question: How spark continuously update query result?

## High level overview

Example: Streaming word count

### Step 1

define source, here input is kafka

<img src="resources/imgs/spark_structured_stream_tahadas_example_s1.png" alt="spark_structured_stream_tahadas_example_s1" width="600"/>

### Step 2

Convert records into string as key and count number of each key

<img src="resources/imgs/spark_structured_stream_tahadas_example_s2.png" alt="spark_structured_stream_tahadas_example_s2" width="600"/>


<img src="resources/imgs/spark_structured_stream_tahadas_example_s3.png" alt="spark_structured_stream_tahadas_example_s3" width="800"/>

### Step3

What to do with final wordcount -> define sink


<img src="resources/imgs/spark_structured_stream_tahadas_example_s4.png" alt="spark_structured_stream_tahadas_example_s4" width="600"/>

### Step 4

When -> how you want it to be executed


<img src="resources/imgs/spark_structured_stream_tahadas_example_s5.png" alt="spark_structured_stream_tahadas_example_s5" width="600"/>

### Step 5

Fault tolerant


<img src="resources/imgs/spark_structured_stream_tahadas_example_s6.png" alt="spark_structured_stream_tahadas_example_s6" width="600"/>


<img src="resources/imgs/spark_structured_stream_tahadas_example_s6_cp.png" alt="spark_structured_stream_tahadas_example_s6_cp" width="600"/>

**compatible**

## Details

Traditional way of design complex streaming ETL



<img src="resources/imgs/spark_structured_stream_tahadas_traditional_etl.png" alt="spark_structured_stream_tahadas_traditional_etl" width="600"/>

Issue: for data to be usable, it takes hours, long time

<img src="resources/imgs/spark_structured_stream_tahadas_spark_etl.png" alt="spark_structured_stream_tahadas_spark_etl" width="600"/>

### How

Another example

<img src="resources/imgs/spark_structured_stream_tahadas_example2_1.png" alt="spark_structured_stream_tahadas_example2_1" width="800"/>


<img src="resources/imgs/spark_structured_stream_tahadas_example2_2.png" alt="spark_structured_stream_tahadas_example2_2" width="800"/>


Raw data -> Dataframe, a collection of row

<img src="resources/imgs/spark_structured_stream_tahadas_example2_3.png" alt="spark_structured_stream_tahadas_example2_3" width="600"/>



3 steps:  
1. cast to string, then into json format  
2. json string to nested columns  
3. flatten:nested -> un-nested  

=> powerful build-in APIs to perform complex data transformations
https://docs.databricks.com/_static/notebooks/transform-complex-data-types-scala.html


<img src="resources/imgs/spark_structured_stream_tahadas_example2_4.png" alt="spark_structured_stream_tahadas_example2_4" width="600"/>



<img src="resources/imgs/spark_structured_stream_tahadas_example2_5.png" alt="spark_structured_stream_tahadas_example2_5" width="600"/>



<img src="resources/imgs/spark_structured_stream_tahadas_example2_6.png" alt="spark_structured_stream_tahadas_example2_6" width="600"/>

Return a handle to streaming query


<img src="resources/imgs/spark_structured_stream_tahadas_example2_7.png" alt="spark_structured_stream_tahadas_example2_7" width="600"/>


<img src="resources/imgs/spark_structured_stream_tahadas_example2_8.png" alt="spark_structured_stream_tahadas_example2_8" width="600"/>

How parquet table is updated incrementally?  My understanding is parquet file is partitioned by time, `parseData`'s result will continuously, small patch by small patch, append to certain parquet file.   

More info:  
https://databricks.com/spark/getting-started-with-apache-spark/streaming

### Play with Time

#### Event time


<img src="resources/imgs/spark_structured_stream_tahadas_eventtime_issue.png" alt="spark_structured_stream_tahadas_eventtime_issue" width="600"/>

**Windowing is another kind of grouping**


<img src="resources/imgs/spark_structured_stream_tahadas_eventtime_agg.png" alt="spark_structured_stream_tahadas_eventtime_agg" width="600"/>


#### How to aggregate

Important slide:  

<img src="resources/imgs/spark_structured_stream_tahadas_eventtime_agg_how.png" alt="spark_structured_stream_tahadas_eventtime_agg_how" width="600"/>


- Inside spark, there is running aggregation going on for every window.
- To keep this aggregations alive across micro baches
- Keep state for every trigger in distrubute env
- State record in excutor's memory, write-ahead log, including check point location 


<img src="resources/imgs/spark_structured_stream_tahadas_handle_late_data.png" alt="spark_structured_stream_tahadas_handle_late_data" width="600"/>

### Watermarking

- How long to keep each window open? 
- Limit the size of state to be aggregate
- System keep on tracking `max event time`(most latest)


<img src="resources/imgs/spark_structured_stream_tahadas_watermark1.png" alt="spark_structured_stream_tahadas_handle_late_data" width="600"/>


<img src="resources/imgs/spark_structured_stream_tahadas_watermark2.png" alt="spark_structured_stream_tahadas_handle_late_data" width="600"/>

How late data do you want


<img src="resources/imgs/spark_structured_stream_tahadas_watermark3.png" alt="spark_structured_stream_tahadas_handle_late_data" width="600"/>

### 3 time to be distinguish



<img src="resources/imgs/spark_structured_stream_tahadas_3_time.png" alt="spark_structured_stream_tahadas_3_time" width="600"/>