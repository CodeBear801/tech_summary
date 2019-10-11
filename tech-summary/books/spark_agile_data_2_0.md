

# Agile Data Science 2.0

[Book](https://www.amazon.com/Agile-Data-Science-2-0-Applications/dp/1491960116) [Code](https://github.com/rjurney/Agile_Data_Code_2) [Video](https://www.youtube.com/watch?v=QqXC0k7sxRI) [Slides](https://www.slideshare.net/rjurney/predictive-analytics-with-airflow-and-pyspark?from_action=save)


## Notes

### Chapter 2

#### Setting up ENV with docker

- <span style="color:blue">Oracle java authentication issue </span></br>
 Description: When come to java part met following issue
```
Package 'oracle-java8-installer' has no installation candidate
```
Cause: Oracle JDK License has changed starting from April 16, 2019
Solution: I switched to OPenJDK, but there are other solutions
https://newfivefour.com/docker-java8-auto-install.html
https://askubuntu.com/questions/1136104/get-error-when-install-oracle-jdk-8-on-ubuntu-18-04

- <span style="color:blue">Version mis-match issue while installing python dependencies. </span></br>
Description: Lots of errors related with python version in requirement.txt
```
ERROR: flask-appbuilder 1.13.1 has requirement click<8,>=6.7, but you'll have click 6.6 which is incompatible.
ERROR: flask-appbuilder 1.13.1 has requirement Flask<2,>=0.12, but you'll have flask 0.11.1 which is incompatible.
ERROR: flask-jwt-extended 3.23.0 has requirement Flask>=1.0, but you'll have flask 0.11.1 which is incompatible.
ERROR: flask-jwt-extended 3.23.0 has requirement Werkzeug>=0.14, but you'll have werkzeug 0.11.11 which is incompatible.
ERROR: pendulum 1.4.4 has requirement python-dateutil<3.0.0.0,>=2.6.0.0, but you'll have python-dateutil 2.5.3 which is incompatible.
ERROR: flask-caching 1.3.3 has requirement Werkzeug>=0.12, but you'll have werkzeug 0.11.11 which is incompatible.
ERROR: jsonschema 3.0.2 has requirement six>=1.11.0, but you'll have six 1.10.0 which is incompatible.
ERROR: apache-airflow 1.10.5 has requirement flask<2.0,>=1.1.0, but you'll have flask 0.11.1 which is incompatible.
ERROR: apache-airflow 1.10.5 has requirement jinja2<2.11.0,>=2.10.1, but you'll have jinja2 2.8 which is incompatible.
ERROR: apache-airflow 1.10.5 has requirement requests<3,>=2.20.0, but you'll have requests 2.11.1 which is incompatible.
```
Cause: The [requirements.txt](https://github.com/rjurney/Agile_Data_Code_2/blob/dcc4fb20d1a7f23304244786dca3f6a9be89322d/requirements.txt#L1) I used didn't pin to specific versions and in docker it uses ['python=3.4'](https://github.com/rjurney/Agile_Data_Code_2/blob/dcc4fb20d1a7f23304244786dca3f6a9be89322d/Dockerfile#L22). 
Solution: Upgrade python version to 3.5 and pip some dependencies to specific version solve this part of issue.  You could find changes [here](https://github.com/CodeBear801/Agile_Data_Code_2/commit/55ac5b2b47edcd733028f9dbde9437a967c1fd55)

- <span style="color:blue">Package 'mongodb-org' has no installation candidate</span>

Description: MongoDB authentication failed
Solution: Update 'apt-key --recv 0C49F3730359A14518585931BC711F9BA15703C6' and using 'mongodb' instead of 'mongodb-org'.  This fix cuased a lot of time to try with different key.
https://askubuntu.com/questions/842592/apt-get-fails-on-16-04-or-18-04-installing-mongodb   

- <span style="color:blue">Fix invalid link issue for kafka/zeppelin</span>
Solution: Old url is invalid, update to correct link.


Notes: All the upper fix could be found [here](https://github.com/CodeBear801/Agile_Data_Code_2/commit/55ac5b2b47edcd733028f9dbde9437a967c1fd55).  Docker image could be successfully generated after this modifications.  But there are still other issues inside the image.

- <span style="color:blue">**ModuleNotFoundError: No module named 'pyspark'**</span>
Solution:
```
export PYTHONPATH=$SPARK_HOME/python/:$PYTHONPATH
```

- <span style="color:blue">Elasticsearch could not be started with root</span>
Solution: Start elastic search with mapuser
```
adduser mapuser
usermod -aG sudo mapuser
chown mapuser:mapuser -R /root/elasticsearch/
chown mapuser:mapuser -R /root/elasticsearch/data/
// In mapuser, use chmod 777 to enable /root/elasticsearch/data/ 
```
Use following command to run elasticsearch in background
```
/root/elasticsearch/bin/elasticsearch -d
```
Then we could test elasticsearch with 'curl localhost:9200'

- <span style="color:blue">Elasticsearch could not started</span>  

Problem: {"acknowledged":true,"shards_acknowledged":false}

Debug: curl -XGET 'localhost:9200/_cluster/allocation/explain' ?

Solution: Kill all elasticsearch process and restart, problem solved

### Chapter 4
- <span style="color:blue">While starting docker container in previous step, I didn't do port mapping, but if we want to test with browser we need open 8080 and 5000.   </span>Here is my solution
```
// Save docker container and then restart another one with port mapping
docker commit 095f788b525c agile_data_2:v2
docker run -v /Users/ngxuser/Desktop/data/agiledata:/data -p 5000:5000 -p 8080:8080 -it agile_data_2:v2 /bin/bash
Docker exec -it c4626fff799d /bin/bash
```
- <span style="color:blue">During running python script: NameError: name 'spark' is not defined </span>

Solution1: https://stackoverflow.com/questions/39541204/pyspark-nameerror-name-spark-is-not-defined
```python
from pyspark.context import SparkContext
from pyspark.sql.session import SparkSession
sc = SparkContext.getOrCreate()
spark = SparkSession(sc)
```

Solution 2: (recommended)
```python
  try:
    sc and spark
  except NameError as e:
    import findspark
    findspark.init()
    import pyspark
    import pyspark.sql
    
    sc = pyspark.SparkContext()
    spark = pyspark.sql.SparkSession(sc).builder.appName(APP_NAME).getOrCreate()

```
- <span style="color:blue">Write mongo db failed</span>
Issue description
```
root@c4626fff799d:~/Agile_Data_Code_2/ch04# mongo agile_data_science
MongoDB shell version v3.6.3
connecting to: mongodb://127.0.0.1:27017/agile_data_science
2019-09-30T22:34:32.020+0000 W NETWORK  [thread1] Failed to connect to 127.0.0.1:27017, in(checking socket for error after poll), reason: Connection refused
2019-09-30T22:34:32.022+0000 E QUERY    [thread1] Error: couldn't connect to server 127.0.0.1:27017, connection attempt failed :
connect@src/mongo/shell/mongo.js:251:13
@(connect):1:6
exception: connect failed
```
Solution: **service mongodb restart**

- Spark Sql API: [registerTempTable](http://spark.apache.org/docs/2.2.0/api/python/pyspark.sql.html#pyspark.sql.DataFrame.registerTempTable)


### Chapter 5

- Code note: drop un-needed fields, only keep elements needed for the following steps 
```
# Filter down to the fields we need to identify and link to a flight
flights = on_time_dataframe.rdd.map(
    lambda x: 
  {
      'Carrier': x.Carrier, 
      'FlightDate': x.FlightDate, 
      'FlightNum': x.FlightNum, 
      'Origin': x.Origin, 
      'Dest': x.Dest, 
      'TailNum': x.TailNum
  }
)
```

- Example about spark mr
```python
flights_per_airplane = flights\
  .map(lambda record: (record['TailNum'], [record]))\
  .reduceByKey(lambda a, b: a + b)\
  .map(lambda tuple:
      {
        'TailNum': tuple[0], 
        'Flights': sorted(tuple[1], key=lambda x: (x['FlightNum'], x['FlightDate'], x['Origin'], x['Dest']))
      }
    )
```
1. Take `TailNum` as key and retrieve from list
2. If Two record has the same key, which is `TailNum`, merge content together
3. sort elements

- OSError: [Errno 98] Address already in use  
Kill related process: https://stackoverflow.com/questions/19071512/socket-error-errno-48-address-already-in-use




### Chapter 6
- When downloading data from openflights, please be aware to download raw data format.  The format I downloaded contains lots of html tags which is incorrect.  https://raw.githubusercontent.com/jpatokal/openflights/master/data/routes.dat

- Collect all data together
```
os.system("cp data/our_airlines.json/part* data/our_airlines.jsonl")
```

- pySpark table join
```
tail_num_plus_inquiry = unique_tail_numbers.join(
  faa_tail_number_inquiry,
  unique_tail_numbers.TailNum == faa_tail_number_inquiry.TailNum,
)
tail_num_plus_inquiry = tail_num_plus_inquiry.drop(unique_tail_numbers.TailNum)
```

- pySpark data frame drop
https://spark.apache.org/docs/2.1.0/api/python/pyspark.sql.html#pyspark.sql.DataFrame.drop

- The reason of using mongodb is recording processed result, such as several table's join result

- Issue of the book
```
airplanes.write.format("org.elasticsearch.spark.sql")\
  .option("es.resource","agile_data_science/airplane")\
  .mode("overwrite")\
  .save()


curl -XGET 'localhost:9200/agile_data_science/airplanes/_search?q=*'
// There should be 's' after airplanes
```

### Chapter 7

- Some link related with regression
https://stattrek.com/statistics/dictionary.aspx?definition=Regression

https://projects.ncsu.edu/labwrite/res/gt/gt-reg-home.html

- late_flights.sample example
```
late_flights.sample(False, 0.01).show()
```
https://spark.apache.org/docs/latest/api/python/pyspark.sql.html?highlight=tojson#pyspark.sql.DataFrame.sample

- Code need more investigation

```
weather_delay_histogram = on_time_dataframe\
  .select("WeatherDelay")\
  .rdd\
  .flatMap(lambda x: x)\
  .histogram([1, 15, 30, 60, 120, 240, 480, 720, 24*60.0])
print(weather_delay_histogram)
```
output
```
>>> print(weather_delay_histogram)
([0.0, 121.1, 242.2, 363.29999999999995, 484.4, 605.5, 726.5999999999999, 847.6999999999999, 968.8, 1089.8999999999999, 1211.0], [1057655, 4477, 873, 216, 94, 55, 29, 20, 15, 5])

>>> print(weather_delay_histogram)
([1, 15, 30, 60, 120, 240, 480, 720, 1440.0], [19740, 16007, 13569, 9442, 4598, 1136, 152, 72])

```

- Error in book
```
Testing model persistance...
Traceback (most recent call last):
  File "train_sklearn_model.py", line 135, in <module>
    model_f = open(regressor_path, 'wb')
FileNotFoundError: [Errno 2] No such file or directory: '/Agile_Data_Code_2/models/sklearn_regressor.pkl'
```


### Chapter 8
- Restart mongo
```
service mongodb restart
```
- A better way to init sc and spark
```
  try:
    sc and spark
  except NameError as e:
    import findspark
    findspark.init()
    import pyspark
    import pyspark.sql
    
    sc = pyspark.SparkContext()
    spark = pyspark.sql.SparkSession(sc).builder.appName(APP_NAME).getOrCreate()
```

- Commands
```
python fetch_prediction_requests.py 2019-10-04 .

cat data/prediction_tasks_daily.json/2019-10-04/part-00000

python make_predictions.py 2019-10-04 .

Input path does not exist: file:/root/Agile_Data_Code_2/ch08/models/string_indexer_model_DayOfMonth.bin/metadata

python load_prediction_results.py 2019-10-04 .
```

- Airflow
```
ln -s $PROJECT_HOME/ch02/airflow_setup
airflow test agile_data_science_batch_prediction_model_training pyspark_extract_features 2019-10-04
airflow backfil -s 2019-10-04 -e 2019-10-04 agile_data_science_batch_predictions_daily
```

- Kafka
```
kafka/bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic flight_delay_classification_request
kafka/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic flight_delay_classification_request --from-beginning
kafka/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic flight_delay_classification_request --from-beginning
```

- Issue related with linking spark-streaming-kafka lib
```
  Spark Streaming's Kafka libraries not found in class path. Try one of the following.

  1. Include the Kafka library and its dependencies with in the
     spark-submit command as

     $ bin/spark-submit --packages org.apache.spark:spark-streaming-kafka-0-8:2.1.1 ...

  2. Download the JAR of the artifact from Maven Central http://search.maven.org/,
     Group Id = org.apache.spark, Artifact Id = spark-streaming-kafka-0-8-assembly, Version = 2.1.1.
     Then, include the jar in the spark-submit command as

#!/usr/bin/env python
     $ bin/spark-submit --jars <spark-streaming-kafka-0-8-assembly.jar> ...
```
reference:
    https://community.cloudera.com/t5/Support-Questions/getting-error-while-submitting-spark-job/td-p/129955  
    https://spark.apache.org/docs/latest/streaming-kafka-0-8-integration.html  
    https://repo1.maven.org/maven2/org/apache/spark/spark-streaming-kafka-0-8-assembly_2.11/2.1.0/  
how I fix it:
```
// https://stackoverflow.com/questions/35910427/java-lang-noclassdeffounderror-kafka-common-topicandpartition
/root/spark/bin/spark-submit --jars /root/spark/jars/spark-streaming-kafka-0-8-assembly_2.11-2.1.0.jar make_predictions_streaming.py .
```



## More information

### Spark
- [Spark programming-guide](https://spark.apache.org/docs/2.1.1/programming-guide.html)
- [PySpark SQL](https://spark.apache.org/docs/latest/api/python/pyspark.sql.html)
- [Apache Spark Examples](https://spark.apache.org/examples.html)

### MongoDB
- [Mongo db methods](https://docs.mongodb.com/manual/reference/method/)

### Elasticsearch
- [Elasticsearch学习](https://blog.csdn.net/laoyang360/article/details/52244917)
- [Configuration](https://www.elastic.co/guide/en/elasticsearch/hadoop/current/configuration.html)

### Airflow
- [Airbnb airflow](https://medium.com/airbnb-engineering/airflow-a-workflow-management-platform-46318b977fd8)
- [Airflow介绍](http://lxwei.github.io/posts/airflow%E4%BB%8B%E7%BB%8D.html)

### D3.js
- [Mike Bostock’s Blocks](https://bl.ocks.org/mbostock) <span>&#9733;</span><span>&#9733;</span><span>&#9733;</span><span>&#9733;</span>

### Frontend
- [菜鸟网站](https://www.runoob.com/)
- [W3 school](https://www.w3schools.com/html/default.asp)

### BootStrap
- [Twitter Bootstrap](http://www.runoob.com/bootstrap/bootstrap-tutorial.html)

### Jinjia2 
- [github Jinjia flask example](https://github.com/mjhea0/thinkful-mentor/tree/master/python/jinja/flask_example)
- [Primer on jinja templating](https://realpython.com/primer-on-jinja-templating/)

### Others
- [Zsh shortcuts](http://www.geekmind.net/2011/01/shortcuts-to-improve-your-bash-zsh.html)
- [Updated data URL for openflights](https://github.com/jpatokal/openflights/tree/master/data)
- [On time performance file](https://s3.amazonaws.com/agile_data_science/On_Time_On_Time_Performance_2015.csv.bz2)



