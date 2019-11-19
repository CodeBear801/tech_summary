# Confluent example cp-all-in-one

This example coming from [confluentinc/examples/cp-all-in-one](https://github.com/confluentinc/examples/tree/5.3.1-post/cp-all-in-one), which use kafka connector to communicate with other data sources and use ksql persistent result.


## Operation
Verify that all the services are up and running.
```
docker-compose ps
```
If the state is not up, run
```
docker-compose up -d
```

## Topic

Create topic named as `users` and `pageviews`
```
docker-compose exec broker kafka-topics --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 1 --topic users

docker-compose exec broker kafka-topics --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 1 --topic pageviews
```


## Connector

```
wget https://github.com/confluentinc/kafka-connect-datagen/raw/master/config/connector_pageviews_cos.config 
curl -X POST -H "Content-Type: application/json" --data @connector_pageviews_cos.config http://localhost:8083/connectors

wget https://github.com/confluentinc/kafka-connect-datagen/raw/master/config/connector_users_cos.config
curl -X POST -H "Content-Type: application/json" --data @connector_users_cos.config http://localhost:8083/connectors
```

## Ksql

```
docker-compose exec ksql-cli ksql http://ksql-server:8088

CREATE STREAM pageviews (viewtime BIGINT, userid VARCHAR, pageid VARCHAR) WITH (KAFKA_TOPIC='pageviews', VALUE_FORMAT='AVRO');


CREATE TABLE users (registertime BIGINT, gender VARCHAR, regionid VARCHAR, userid VARCHAR) WITH (KAFKA_TOPIC='users', VALUE_FORMAT='AVRO', KEY = 'userid');

```

Create persistent query and write result into stream
```
CREATE STREAM pageviews_female AS SELECT users.userid AS userid, pageid, regionid, gender FROM pageviews LEFT JOIN users ON pageviews.userid = users.userid WHERE gender = 'FEMALE';
```

Create persistent query, result from this query are writtern to a kafka topic pageviews_enriched_r8_r9

```
CREATE STREAM pageviews_female_like_89 WITH (kafka_topic='pageviews_enriched_r8_r9', value_format='AVRO') AS SELECT * FROM pageviews_female WHERE regionid LIKE '%_8' OR regionid LIKE '%_9';
```

[Tumbling window](https://docs.confluent.io/current/streams/developer-guide/dsl-api.html#windowing-tumbling) example

```
CREATE TABLE pageviews_regions AS SELECT gender, regionid , COUNT(*) AS numusers FROM pageviews_female WINDOW TUMBLING (size 30 second) GROUP BY gender, regionid HAVING COUNT(*) > 1;
```

Explain
```
DESCRIBE EXTENDED pageviews_female_like_89;

SHOW QUERIES;

EXPLAIN CTAS_PAGEVIEWS_REGIONS_1
```

## More info
- [confluentinc/examples/cp-all-in-one](https://github.com/confluentinc/examples/tree/5.3.1-post/cp-all-in-one)
- [Quick Start using Community Components (Docker)](https://docs.confluent.io/current/quickstart/cos-docker-quickstart.html#cos-docker-quickstart)
