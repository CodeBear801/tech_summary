# Kafka Docker
[Docker image example](https://github.com/wurstmeister/kafka-docker/blob/master/Dockerfile)

## docker-compose file

kafka + zookeeper

```
version: '2'
services:
  zookeeper:
    image: wurstmeister/zookeeper
    ports:
      - "2181:2181"
    restart: unless-stopped
    hostname: zookeeper
    container_name: zookeeper
  kafka1:
    image: wurstmeister/kafka
    ports:
      - "9092:9092"
    environment:
      KAFKA_ADVERTISED_HOST_NAME: localhost
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      KAFKA_BROKER_ID: 1
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
      KAFKA_CREATE_TOPICS: "stream-in:1:1,stream-out:1:1"
    depends_on:
      - zookeeper
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock

```

## Docker commands

Start docker containers
```
docker-compose -f docker-compose.yml up -d
```

List all containers
```
docker ps
```

Check kafka version

```
docker exec docker-kafka_kafka1_1 find / -name \*kafka_\* | head -
```

Check zookeeper version
```
docker exec zookeeper pwd
```

Scale docker broker
```
docker-compose -f docker-compose.yml up -d --scale kafka1=2
```

## Topic



Create topic
```
docker exec docker-kafka_kafka1_1 kafka-topics.sh --create --topic topic001 --partitions 4 --zookeeper zookeeper:2181 --replication-factor 1
```

List topic
```
docker exec docker-kafka_kafka1_1 kafka-topics.sh --list --zookeeper zookeeper:2181 topic001
```

List topic, broker, replica
```
docker exec docker-kafka_kafka1_1 kafka-topics.sh --describe --topic topic001 --zookeeper zookeeper:2181
```

Consumer
```
docker exec docker-kafka_kafka1_1 kafka-console-consumer.sh --topic topic001 --bootstrap-server docker-kafka_kafka1_1:9092
```

Producer
```
docker exec -it docker-kafka_kafka1_1 kafka-console-producer.sh --topic topic001 --broker-list docker-kafka_kafka1_1:9092
```

For more command and information, please go to [here](https://hub.docker.com/r/wurstmeister/kafka/)


## Others
- [Fast data dev (20 +connectors)](https://github.com/lensesio/fast-data-dev/wiki)

