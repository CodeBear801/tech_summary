# Kafka example about producer and consumer

This page follows [example given by MapR](https://github.com/mapr-demos/kafka-sample-programs.git) to experience Kafka's producer and consumer API.

## Set up env

Please first follow the [docker setup](./kafka_docker_setup.md) to create a kafka+zookeeper running env.  But for development env, you need following tools.


### JDK
```
apk add openjdk-8-jdk
```

### mvn
- [mvn docker example](https://github.com/Zenika/alpine-maven/blob/master/jdk8/Dockerfile)
- [mvn install instruction](https://www.baeldung.com/install-maven-on-windows-linux-mac#installing-maven-on-linux)
- [mvn download information](https://maven.apache.org/download.cgi)

```
wget http://us.mirrors.quenda.co/apache/maven/maven-3/3.6.2/binaries/apache-maven-3.6.2-bin.zip
unzip -d /usr/local/ apache-maven-3.6.2-bin.zip
export M2_HOME=/usr/local/apache-maven-3.6.2
export M2=$M2_HOME/bin
export MAVEN_OPTS=-Xms256m -Xmx512m
export PATH=$M2:$PATH
```

### gradle
- [gradle install instruction](https://docs.gradle.org/current/userguide/installation.html)

```
wget https://downloads.gradle-dn.com/distributions/gradle-6.0-bin.zip
mkdir /opt/gradle
unzip -d /opt/gradle gradle-6.0.1-bin.zip
export PATH=$PATH:/opt/gradle/gradle-6.0/bin
```

### Issues
- zookeeper timeout
```
// enter zookeeper's docker image
docker exec -it 3b8e7da52099 /bin/bash
// restart zookeeper
./zkServer.sh restart
```

## Operations

create topic
```
docker exec docker-kafka_kafka1_1 kafka-topics.sh --create --topic fast-messages --partitions 1 --zookeeper zookeeper:2181 --replication-factor 1

docker exec docker-kafka_kafka1_1 kafka-topics.sh --create --topic summary-markers --partitions 1 --zookeeper zookeeper:2181 --replication-factor 1

```

list all topics
```
/opt/kafka/bin/kafka-topics.sh --list --zookeeper zookeeper:2181
```

start application
```
target/kafka-example consumer

target/kafka-example producer
```


## Code analysis



[Producer side](https://github.com/CodeBear801/kafka-sample-programs/blob/master/src/main/java/com/mapr/examples/Producer.java)
```java
KafkaProducer<String, String> producer;
// ...
producer = new KafkaProducer<>(properties);
// ... 
producer.send(new ProducerRecord<String, String>(
        "fast-messages",
        String.format(Locale.US, "{\"type\":\"test\", \"t\":%.3f, \"k\":%d}", System.nanoTime() * 1e-9, i)));

// ...
producer.flush();
// ...
producer.close();
```


[Consumer side](https://github.com/CodeBear801/kafka-sample-programs/blob/master/src/main/java/com/mapr/examples/Consumer.java)
```java
KafkaConsumer<String, String> consumer;
// ... 
consumer = new KafkaConsumer<>(properties);
// ...
consumer.subscribe(Arrays.asList("fast-messages", "summary-markers"));
// ...
while (true) {
   ConsumerRecords<String, String> records = consumer.poll(200);
   for (ConsumerRecord<String, String> record : records) {
       switch (msg.get("type").asText()) {
       case "test":
           // ...
       }
   }

```

## others
- [mapr-demos/kafka-sample-programs](https://github.com/mapr-demos/kafka-sample-programs)
- [Apache Kafka Tutorial](https://mapr.com/blog/getting-started-sample-programs-apache-kafka-09/)
