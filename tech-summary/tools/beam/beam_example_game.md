
## Reference link
- https://cloud.google.com/dataflow/examples/gaming-example
- https://github.com/apache/beam/tree/master/examples/java/src/main/java/org/apache/beam/examples/complete/game
- 

## wordcount
- https://beam.apache.org/get-started/wordcount-example/
- https://github.com/apache/beam/tree/master/examples/java

Generate pom
```
mvn archetype:generate \
      -DarchetypeGroupId=org.apache.beam \
      -DarchetypeArtifactId=beam-sdks-java-maven-archetypes-examples \
      -DarchetypeVersion=2.21.0 \
      -DgroupId=org.example \
      -DartifactId=word-count-beam \
      -Dversion="0.1" \
      -Dpackage=org.apache.beam.examples \
      -DinteractiveMode=false
```

Run WordCount
```
mvn compile exec:java -Dexec.mainClass=org.apache.beam.examples.WordCount \
     -Dexec.args="--inputFile=input_1.txt --output=counts" -Pdirect-runner
```

Test
```
more counts*
```

Run MinimalWordCount
```
mvn compile exec:java -Dexec.mainClass=org.apache.beam.examples.MinimalWordCount
```


## Game

Generate pom
```
mvn archetype:generate \
      -DarchetypeGroupId=org.apache.beam \
      -DarchetypeArtifactId=beam-sdks-java-maven-archetypes-examples \
      -DarchetypeVersion=2.21.0 \
      -DgroupId=org.example \
      -DartifactId=word-count-beam \
      -Dversion="0.1" \
      -Dpackage=org.apache.beam.examples \
      -DinteractiveMode=false
```


Generate `beamgame.txt` in local
```
mvn compile exec:java -Dexec.mainClass=org.apache.beam.examples.complete.game.injector.Injector -Dexec.args="beamgame none beamgame.txt"
```


### UserScore

```
mvn compile exec:java -Dexec.mainClass=org.apache.beam.examples.complete.game.UserScore -Dexec.args="--output=userscore" -Pdirect-runner
```


```
cat userscore-0000* > userscore-total
sort -rn -t $'\t' -k2,2rn userscore-total
```
- https://superuser.com/questions/148456/merging-and-sorting-multiple-files-with-sort



### HourlyTeamScore
```
mvn compile exec:java -Dexec.mainClass=org.apache.beam.examples.complete.game.HourlyTeamScore -Dexec.args="--output=userscorehourly" -Pdirect-runner
```


### LeaderBoard
steps:
- Input 
  + Injector simulates input, it could publish result to pubsub
  + create `pub/sub` topic, following description [here](https://developers.google.com/identity/protocols/application-default-credentials) or refer to comments in `injector.java`
  + export GOOGLE_APPLICATION_CREDENTIALS=54416cf616c5.json
  + `pub/sub` topic could be deleted during re-test.  More advaced operation could go to here: [pubsub delete all topics until now](https://stackoverflow.com/questions/39398173/best-practices-for-draining-or-clearing-a-google-cloud-pubsub-topic)
  + https://console.cloud.google.com/cloudpubsub/topic/list?project=constant-jigsaw-272415

```
mvn compile exec:java -Dexec.mainClass=org.apache.beam.examples.complete.game.injector.Injector -Dexec.args="constant-jigsaw-272415 beamgame none"

// * Injector <project-name> <topic-name> none
// Injector constant-jigsaw-272415 beamgame none

// delete topic
gcloud pubsub subscriptions seek projects/constant-jigsaw-272415/topics/beamgame --time=$(date +%Y-%m-%dT%H:%M:%S) 
```

```java
// some modification to make it running slower
  // QPS ranges from 800 to 1000.
  private static final int MIN_QPS = 80;
  private static final int QPS_RANGE = 20;
  // How long to sleep, in ms, between creation of the threads that make API requests to PubSub.
  private static final int THREAD_SLEEP_MS = 5000;
```


- LeaderBoard
   + Create bigquery dataset https://console.cloud.google.com/bigquery?p=bigquery-public-data&page=project


```
mvn compile exec:java -Dexec.mainClass=org.apache.beam.examples.complete.game.LeaderBoard -Dexec.args="--project=constant-jigsaw-272415 --dataset=beamgame2 --topic=projects/constant-jigsaw-272415/topics/beamgame"
```

## Experiment

<img src="
https://user-images.githubusercontent.com/16873751/84213835-8d32c900-aa76-11ea-827a-f17062bb1b4e.png" alt="dataflow_frances_perry_mobile_game" width="400"/><br/>


<img src="
https://user-images.githubusercontent.com/16873751/84213843-902db980-aa76-11ea-8794-a80d7224ff0c.png" alt="dataflow_frances_perry_mobile_game" width="400"/><br/>

