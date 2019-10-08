
[Book](https://www.amazon.com/Agile-Data-Science-2-0-Applications/dp/1491960116) [Code](https://github.com/rjurney/Agile_Data_Code_2) [Video](https://www.youtube.com/watch?v=QqXC0k7sxRI) [Slides](https://www.slideshare.net/rjurney/predictive-analytics-with-airflow-and-pyspark?from_action=save)


Chapter 2

**Setting up ENV with docker**

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

Chapter 4
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



