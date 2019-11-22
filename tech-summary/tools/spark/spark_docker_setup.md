# Spark docker setup

## Docker file
[Dockerfile](./Dockerfile)

## Operations
```
docker build -t spark-local:v2.4.1 -f Dockerfile .
docker run -v /Users/xunliu/Desktop/git/kafka/docker-example/code:/code -it spark-local:v2.4.1 /bin/bash
```

## Issues
- When using pyspark, met error message "No module named pyspark"
```
root@76899858b727:/usr/spark-2.4.1/examples/src/main/python# python wordcount.py /code/osrm-customize-experiment-raw-log.txt
Traceback (most recent call last):
  File "wordcount.py", line 23, in <module>
    from pyspark.sql import SparkSession
ImportError: No module named 'pyspark'
```
Solution: 
```
export PYTHONPATH=$SPARK_HOME/python/:$PYTHONPATH
```

- pandas not installed
```
apt-get install python3-pandas
pip install pyarrow
```

## Examples

### python case 1
Running case  
https://github.com/apache/spark/blob/v2.4.1/examples/src/main/python/wordcount.py  

Experiment:
View result in reverse order.   [ref](https://stackoverflow.com/a/8459243)
```python
output = counts.collect()
for (word, count) in sorted(output, key = lambda x:x[1], reverse = True):
print("%s: %i" % (word, count))
```

### python case 2

Running case   
https://github.com/apache/spark/blob/v2.4.1/examples/src/main/python/status_api_demo.py

Reference:   
https://spark.apache.org/docs/2.2.0/api/python/pyspark.html#pyspark.StatusTracker  
https://spark.apache.org/docs/2.2.0/api/python/pyspark.html#pyspark.StatusTracker.getStageInfo  


### python case 3
Running case   
https://github.com/apache/spark/blob/v2.4.1/examples/src/main/python/sql/basic.py


### python case 4

[hotspot.py](./hotspot.py)

Reference:  
https://spark.apache.org/docs/2.1.0/api/python/pyspark.sql.html  
https://spark.apache.org/docs/2.1.0/api/python/pyspark.html  

### python case 5

[get_medium.py](./get_medium.py)


The issue I always met with python
```
...
TypeError: <lambda>() missing 1 required positional argument: 'v'
...
```
The rdd data is
```
[(0, 0), (0, 1),
```
The error happened here:  

```python
rdd.filter(lambda k, v : k == index)
```

Solution: https://www.reddit.com/r/learnpython/comments/4zo7q3/python2_python3_sortinglambda_error/  
"Because it's not the same code. (k,v) does not mean the same thing as k,v. I want key=lambda s: s[0]"  

```python
rdd.filter(lambda k : k[0] == index)
```




## More info
- [Reference docker file from Getty Images](https://github.com/gettyimages/docker-spark)



