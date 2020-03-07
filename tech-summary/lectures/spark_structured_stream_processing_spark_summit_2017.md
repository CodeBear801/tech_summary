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