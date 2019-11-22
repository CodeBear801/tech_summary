

# KafkaMusicExample
- [KafkaMusicExample.java](https://github.com/confluentinc/kafka-streams-examples/blob/5.3.1-post/src/main/java/io/confluent/examples/streams/interactivequeries/kafkamusic/KafkaMusicExample.java)
- [KafkaMusicExampleDriver.java](https://github.com/confluentinc/kafka-streams-examples/blob/5.3.1-post/src/main/java/io/confluent/examples/streams/interactivequeries/kafkamusic/KafkaMusicExampleDriver.java)
- [Kafka Streams Demo Application](https://docs.confluent.io/current/streams/kafka-streams-examples/docs/index.html)
- [Env Steup](https://github.com/confluentinc/kafka-streams-examples#packaging-and-running)


## Operations

Re-use the env setup of [cp-all-in-one](./kafka_example_confluent_cp_all_in_one.md)
```
docker-compose up -d --build

```

Create topic
```
docker-compose exec broker kafka-topics --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 1 --topic users

// not working???
docker-compose exec broker kafka-topics --list --zookeeper localhost:2181
```


Build package
```
// https://github.com/confluentinc/kafka-streams-examples#packaging-and-running
mvn -DskipTests=true clean package
```


Create topic
```
docker-compose exec broker kafka-topics --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 2 --topic play-events

docker-compose exec broker kafka-topics --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 2 --topic song-feed
```

```
java -cp target/kafka-streams-examples-5.3.1-standalone.jar io.confluent.examples.streams.interactivequeries.kafkamusic.KafkaMusicExample 7070

java -cp target/kafka-streams-examples-5.3.1-standalone.jar io.confluent.examples.streams.interactivequeries.kafkamusic.KafkaMusicExample 7071

java -cp target/kafka-streams-examples-5.3.1-standalone.jar io.confluent.examples.streams.interactivequeries.kafkamusic.KafkaMusicExampleDriver
```


```
 * # List all running instances of this application
http://localhost:7070/kafka-music/instances
 *
 * # List app instances that currently manage (parts of) state store "song-play-count"
http://localhost:7070/kafka-music/instances/song-play-count
 *
 * # Get the latest top five for the genre "punk"
http://localhost:7070/kafka-music/charts/genre/punk
 *
 * # Get the latest top five across all genres
http://localhost:7070/kafka-music/charts/top-five
```


## protocal

[playevent](https://github.com/confluentinc/kafka-streams-examples/blob/5.3.1-post/src/main/resources/avro/io/confluent/examples/streams/playevent.avsc) 
```
{"namespace": "io.confluent.examples.streams.avro",
 "type": "record",
 "name": "PlayEvent",
 "fields": [
     {"name": "song_id",   "type": "long"},
     {"name": "duration", "type": "long"}
 ]
}
```

[Song](https://github.com/confluentinc/kafka-streams-examples/blob/5.3.1-post/src/main/resources/avro/io/confluent/examples/streams/song.avsc) 
```
{"namespace": "io.confluent.examples.streams.avro",
 "type": "record",
 "name": "Song",
 "fields": [
     {"name": "id",     "type": "long"},
     {"name": "album",  "type": "string"},
     {"name": "artist", "type": "string"},
     {"name": "name",   "type": "string"},
     {"name": "genre",  "type": "string"}
 ]
}
```

[SongPlayCount](https://github.com/confluentinc/kafka-streams-examples/blob/5.3.1-post/src/main/resources/avro/io/confluent/examples/streams/songplaycount.avsc) 
```
{"namespace": "io.confluent.examples.streams.avro",
 "type": "record",
 "name": "SongPlayCount",
 "fields": [
     {"name": "song_id",  "type": "long"},
     {"name": "plays",   "type": "long"}
 ]
}
```
## [Driver](https://github.com/confluentinc/kafka-streams-examples/blob/5.3.1-post/src/main/java/io/confluent/examples/streams/interactivequeries/kafkamusic/KafkaMusicExampleDriver.java)

Driver will load song into array[code](https://github.com/confluentinc/kafka-streams-examples/blob/453b3ebbf64f765310943ce5b45ddf979900d3b2/src/main/java/io/confluent/examples/streams/interactivequeries/kafkamusic/KafkaMusicExampleDriver.java#L70), then create playevent producer and songproducer.  

Send all the song into topic of song-feed
```java
    songs.forEach(song -> {
      System.out.println("Writing song information for '" + song.getName() + "' to input topic " +
          KafkaMusicExample.SONG_FEED);
      songProducer.send(new ProducerRecord<>(KafkaMusicExample.SONG_FEED, song.getId(), song));
    });

    songProducer.close();
```

Simulate play event and send into topic of play-events
```java
    final long duration = 60 * 1000L;
    final Random random = new Random();

    // send a play event every 100 milliseconds
    while (true) {
      final Song song = songs.get(random.nextInt(songs.size()));
      System.out.println("Writing play event for song " + song.getName() + " to input topic " +
          KafkaMusicExample.PLAY_EVENTS);
      playEventProducer.send(
          new ProducerRecord<>(KafkaMusicExample.PLAY_EVENTS,
                                                "uk", new PlayEvent(song.getId(), duration)));
      Thread.sleep(100L);
    }
```


## [KafkaMusic](https://github.com/confluentinc/kafka-streams-examples/blob/5.3.1-post/src/main/java/io/confluent/examples/streams/interactivequeries/kafkamusic/KafkaMusicExample.java)


Build Table to store all the songs
```java
    // get table and create a state store to hold all the songs in the store
    final KTable<Long, Song>
        songTable =
        builder.table(SONG_FEED, Materialized.<Long, Song, KeyValueStore<Bytes, byte[]>>as(ALL_SONGS)
            .withKeySerde(Serdes.Long())
            .withValueSerde(valueSongSerde));
```

```java
    final SpecificAvroSerde<Song> keySongSerde = new SpecificAvroSerde<>();
    keySongSerde.configure(serdeConfig, true);

    final SpecificAvroSerde<Song> valueSongSerde = new SpecificAvroSerde<>();
    valueSongSerde.configure(serdeConfig, false);
```


Record play event
```java
    final StreamsBuilder builder = new StreamsBuilder();

    // get a stream of play events
    final KStream<String, PlayEvent> playEvents = builder.stream(
        PLAY_EVENTS,
        Consumed.with(Serdes.String(), playEventSerde));

    // final SpecificAvroSerde<PlayEvent> playEventSerde = new SpecificAvroSerde<>();

    // Accept play events that have a duration >= the minimum
    final KStream<Long, PlayEvent> playsBySongId =
        playEvents.filter((region, event) -> event.getDuration() >= MIN_CHARTABLE_DURATION)
            // repartition based on song id
            .map((key, value) -> KeyValue.pair(value.getSongId(), value));

    // [Perry] Here I don't know how region is handled from previously step
    // I don't feel like playEventSerde could handle this
    // 
    // map's purpose is using specific field from value to create key, which is required by join
    // please note that (value.getSongId(), value) -> <Long, PlayEvent>        

    // join the plays with song as we will use it later for charting
    final KStream<Long, Song> songPlays = playsBySongId.leftJoin(songTable,
        (value1, song) -> song,
        Joined.with(Serdes.Long(), playEventSerde, valueSongSerde));

    // [Perry] leftjoin: https://kafka.apache.org/20/javadoc/org/apache/kafka/streams/kstream/KStream.html#leftJoin
    // This KStream-KTable join, it allow you you to perform table lookups against a KTable everytime a new record 
    // is received from the KStream.  Return result is KStream
    // My understanding
    // (value1, song) act as (leftValue, rightValue), the following function will return final value
    //  Joined.with(Serdes.Long(), playEventSerde, valueSongSerde))-> (key, leftvalue, rightvalue), will return final key

    // [Perry] KStream + filter -> KTable for recording

    // create a state store to track song play counts
    final KTable<Song, Long> songPlayCounts = songPlays.groupBy((songId, song) -> song,
                                                                Grouped.with(keySongSerde, valueSongSerde))
            .count(Materialized.<Song, Long, KeyValueStore<Bytes, byte[]>>as(SONG_PLAY_COUNT_STORE)
                           .withKeySerde(valueSongSerde)
                           .withValueSerde(Serdes.Long()));

```


```java
    // Compute the top five charts for each genre. The results of this computation will continuously update the state
    // store "top-five-songs-by-genre", and this state store can then be queried interactively via a REST API (cf.
    // MusicPlaysRestService) for the latest charts per genre.
    // 1
    songPlayCounts
       .groupBy((song, plays) -> // 2
            KeyValue.pair(song.getGenre().toLowerCase(),// 3
                new SongPlayCount(song.getId(), plays)),
        Grouped.with(Serdes.String(), songPlayCountSerde))   
        // aggregate into a TopFiveSongs instance that will keep track
        // of the current top five for each genre. The data will be available in the
        // top-five-songs-genre store
        .aggregate(TopFiveSongs::new,   // 4
            (aggKey, value, aggregate) -> {
              aggregate.add(value);
              return aggregate;
            },
            (aggKey, value, aggregate) -> {
              aggregate.remove(value);
              return aggregate;
            },
            Materialized.<String, TopFiveSongs, KeyValueStore<Bytes, byte[]>>as(TOP_FIVE_SONGS_BY_GENRE_STORE)
                .withKeySerde(Serdes.String())
                .withValueSerde(topFiveSerde)
        );

    // [Perry]
    // 1. songPlayCounts is a KTable with Song as key and play counts as value
    // 2. groupby will generate KGroupedTable, which is needed by aggregation
    //    -> selectKey(song.getGenre().toLowerCase()).groupbykey()
    //    It will trigger repartition based on new key
    //    
    // 3. groupby(KeyValueMapper<? super K,? super V,KR> selector, Grouped<KR,V>grouped) 
    //    means group the records of this KStream on a new key that is selected using 
    //    the provided KeyValueMapper and Serdes as specified by Grouped
    //    
    //    selector is a KeyValueMapper that computers a new key for grouping, it similar 
    //    as VR apply(K key, V value).  Input is song and plays, and in the logic will 
    //    reconstruct a different key(song's Genre) and value(SongPlayCount object) 
    //    based on them
    //    
    //    grouped is used to capture the key and value Serdes and set the part of name
    //    used for repartition topics when performing groupby operations.
    //    public static <K,V> Grouped<K,V> with(org.apache.kafka.common.serialization.Serde<K> keySerde,
    //                                          org.apache.kafka.common.serialization.Serde<V> valueSerde)
    //    The two parameter is Serde for key and value and returns grouped configuration
    //    So, Grouped.with(Serdes.String(), songPlayCountSerde) means
    //    Serdes.String() is for new key song.getGenre().toLowerCase()
    //    songPlayCountSerde is for new value SongPlayCount(song.getId(), plays)
    //    
    // 4. Apply aggregation on kgroupedtable
    //    Initializer<VR> initializer ->  TopFiveSongs::new
    //    Aggregator<? super K,? super V,VR> adder(aggKey, newValue, aggValue)
    //    aggValue is the object of TopFiveSongs::new
    //    newValue means increase/update/delete of (songname, count)
    //    Here should always be increase
    //    Aggregator<? super K,? super V,VR> subtractor
 

```
The implementation of TopFiveSongs::add()
```java
//[Perry] Keep a TreeSet only have top 5 elements
    public void add(final SongPlayCount songPlayCount) {
      if(currentSongs.containsKey(songPlayCount.getSongId())) {
        topFive.remove(currentSongs.remove(songPlayCount.getSongId()));
      }
      topFive.add(songPlayCount);
      currentSongs.put(songPlayCount.getSongId(), songPlayCount);
      if (topFive.size() > 5) {
        final SongPlayCount last = topFive.last();
        currentSongs.remove(last.getSongId());
        topFive.remove(last);
      }
    }
```



## More info
- [how-to-join-two-kafka-streams-and-produce-the-result-in-a-topic-with-avro-values](https://stackoverflow.com/questions/50213221/how-to-join-two-kafka-streams-and-produce-the-result-in-a-topic-with-avro-values)
- [kafka-stream-aggregation-with-custom-object-data-type](https://stackoverflow.com/questions/53400832/kafka-stream-aggregation-with-custom-object-data-type)
- [Schema Management, how to handle multiple version](https://docs.confluent.io/current/schema-registry/index.html)





# Wiki avro example
- [Wiki feed avro file](https://github.com/confluentinc/kafka-streams-examples/blob/5.3.1-post/src/main/resources/avro/io/confluent/examples/streams/wikifeed.avsc)
- [Driver](https://github.com/confluentinc/kafka-streams-examples/blob/5.3.1-post/src/main/java/io/confluent/examples/streams/WikipediaFeedAvroExampleDriver.java)
- [AvroLambdaExample](https://github.com/confluentinc/kafka-streams-examples/blob/5.3.1-post/src/main/java/io/confluent/examples/streams/WikipediaFeedAvroLambdaExample.java)


