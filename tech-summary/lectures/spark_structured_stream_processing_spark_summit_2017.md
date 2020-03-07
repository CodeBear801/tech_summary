# Easy, Scalable, Fault Tolerant Stream Processing with Structured Streaming in Apache Spark

Spark Summit 2017

Question: How could spark continues update query result?

## High level overview

Example: Streaming word count

### Step 1

define source, here input is kafka

<img src="resources/imgs/spark_structured_stream_tahadas_example_s1.png" alt="spark_structured_stream_tahadas_example_s1" width="600"/>

### Step 2

Convert records into string as key and count number of each key

<img src="resources/imgs/spark_structured_stream_tahadas_example_s2.png" alt="spark_structured_stream_tahadas_example_s2" width="600"/>


### Step 3


<img src="resources/imgs/spark_structured_stream_tahadas_example_s3.png" alt="spark_structured_stream_tahadas_example_s3" width="800"/>

### Step4

What to do with final wordcount -> define sink

spark_structured_stream_tahadas_example_s4

### Step 5

When -> how you want it to be executed

spark_structured_stream_tahadas_example_s5

### Step 6

Fault tolerant

spark_structured_stream_tahadas_example_s6

spark_structured_stream_tahadas_example_s6_cp

**compatible**

## Details

Traditional way of design complex streaming ETL

spark_structured_stream_tahadas_traditional_etl

Issue: for data to be usable, it takes hours, long time

spark_structured_stream_tahadas_spark_etl

### How

Another example

spark_structured_stream_tahadas_example2_1

spark_structured_stream_tahadas_example2_2

Raw data -> Dataframe, a collection of row

spark_structured_stream_tahadas_example2_3


3 steps:

1.cast to string, then into json format
2.json string to nested columns
3.flatten:nested -> un-nested

=> powerful build-in APIs to perform complex data transformations
https://docs.databricks.com/_static/notebooks/transform-complex-data-types-scala.html

spark_structured_stream_tahadas_example2_4


spark_structured_stream_tahadas_example2_5


spark_structured_stream_tahadas_example2_6

Return a handle to streaming query

spark_structured_stream_tahadas_example2_7

spark_structured_stream_tahadas_example2_8

How parquet table is updated incrementally?  我理解的应该是只是不停的在写同一个文件

More info:  
https://databricks.com/spark/getting-started-with-apache-spark/streaming

## Play with Time

### Event time

spark_structured_stream_tahadas_eventtime_issue

**Windowing is another kind of grouping**

spark_structured_stream_tahadas_eventtime_agg


### How to aggregate

spark_structured_stream_tahadas_eventtime_agg_how


- Inside spark, there is running aggregation going on for every window.
- To keep this aggregations alive across micro baches
- Keep state for every trigger in distrubute env
- State record in excutor's memory, write-ahead log, including check point location 

spark_structured_stream_tahadas_handle_late_data

### Watermarking

- How long to keep each window open? 
- Limit the size of state to be aggregate
- System keep on tracking `max event time`(most latest)

spark_structured_stream_tahadas_watermark1

spark_structured_stream_tahadas_watermark2

How late data do you want

spark_structured_stream_tahadas_watermark3

### 3 time to be distinguish

spark_structured_stream_tahadas_3_time