[No More Silos: How to Integrate Your Databases with Apache Kafka and CDC](https://www.confluent.io/blog/no-more-silos-how-to-integrate-your-databases-with-apache-kafka-and-cdc)  [SDE](https://softwareengineeringdaily.com/2019/09/23/kafka-data-pipelines-with-robin-moffatt/ )

## Why Kafka: 
Its never just one pipeline for your system, its never just the set of data is used in this one place over there.
Its always, the data needs to populate a search index, and we want to use it for analytics, and we want to share it with this other department, and we want to use it to drive this application.
You either build a specific architecture or use streaming systems like kafka.

## Apache connector
Moving data from database to kafka or vice verse

## ETL process changed
The idea of events and the idea of real-time processing. Building you a data platform around events

## Time
In batch processing, you know what your time window is.  
In streaming processing, need to be aware the difference between event time and system time  

With Kafka, you could have the change data being buffered from your transactional data store into Kafka and then you could have your search system reading from kafka and updating its search indexes.  

## Database, Kafka, Kafka connector
If I want to build search index for my database, at first will take a snapshot of the current database state, record SDK, the points in the transaction log at which that was taken.  Then from then on, capture any of the events cut of the transaction log.  In summary, you get current status through a selector server plus transactional log from that point in using the SDK to make sure you don't miss anything.  

## Why Kafka is high performant
- Write ahead log: how to implement : [Crashes and Recovery: Write-ahead Logging, columbia](https://www.cs.columbia.edu/~du/ds/assets/lectures/lecture15.pdf)
- Events: events not message, it could rebuild state
- There are more...


Workshop: Real-time SQL Stream Processing at Scale with Apache Kafka and KSQL (Mac/Linux)
https://github.com/confluentinc/demo-scene/blob/master/ksql-workshop/ksql-workshop.adoc
