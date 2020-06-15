
## Machine Learning Logistic regression via Cloud Spark

Target
```
y = 0 if arrival delay >= 15 minutes
y = 1 if arrival delay < 15 minutes
// marching learning algorithm predict the probability that the flight is on time
```


### Cloud Dataproc cluster with initialization actions for Datalab

### Code Analysis

[code](https://github.com/CodeBear801/data-science-on-gcp/blob/feature/experiment/07_sparkml_and_bqml/logistic_regression.ipynb)

```py
# Load csv from gs
traindays = spark.read \
    .option("header", "true") \
    .csv('gs://{}/flights/trainday.csv'.format(BUCKET))
traindays.createOrReplaceTempView('traindays')

# Define the header
header = 'FL_DATE,UNIQUE_CARRIER,AIRLINE_ID,CARRIER,FL_NUM,ORIGIN_AIRPORT_ID,ORIGIN_AIRPORT_SEQ_ID,ORIGIN_CITY_MARKET_ID,ORIGIN,DEST_AIRPORT_ID,DEST_AIRPORT_SEQ_ID,DEST_CITY_MARKET_ID,DEST,CRS_DEP_TIME,DEP_TIME,DEP_DELAY,TAXI_OUT,WHEELS_OFF,WHEELS_ON,TAXI_IN,CRS_ARR_TIME,ARR_TIME,ARR_DELAY,CANCELLED,CANCELLATION_CODE,DIVERTED,DISTANCE,DEP_AIRPORT_LAT,DEP_AIRPORT_LON,DEP_AIRPORT_TZOFFSET,ARR_AIRPORT_LAT,ARR_AIRPORT_LON,ARR_AIRPORT_TZOFFSET,EVENT,NOTIFY_TIME'

def get_structfield(colname):
   if colname in ['ARR_DELAY', 'DEP_DELAY', 'DISTANCE', 'TAXI_OUT']:
      return StructField(colname, FloatType(), True)
   else:
      return StructField(colname, StringType(), True)
# https://spark.apache.org/docs/latest/api/java/org/apache/spark/sql/types/StructType.html
schema = StructType([get_structfield(colname) for colname in header.split(',')])


inputs = 'gs://{}/flights/tzcorr/all_flights-00000-*'.format(BUCKET) # 1/30th;  you may have to change this to find a shard that has training data
#inputs = 'gs://{}/flights/tzcorr/all_flights-*'.format(BUCKET)  # FULL
flights = spark.read\
            .schema(schema)\
            .csv(inputs)

# this view can now be queried ...
flights.createOrReplaceTempView('flights')

# Clean up
# ...


trainquery = """
SELECT
  DEP_DELAY, TAXI_OUT, ARR_DELAY, DISTANCE
FROM flights f
JOIN traindays t
ON f.FL_DATE == t.FL_DATE
WHERE
  t.is_train_day == 'True' AND
  f.dep_delay IS NOT NULL AND 
  f.arr_delay IS NOT NULL
"""
traindata = spark.sql(trainquery)


# [perry]For the simple version, it just use 3 features: DEP_DELAY, TAXI_OUT, DISTANCE
def to_example(fields):
  return LabeledPoint(\
              float(fields['ARR_DELAY'] < 15), #ontime? \
              [ \
                  fields['DEP_DELAY'], \
                  fields['TAXI_OUT'],  \
                  fields['DISTANCE'],  \
              ])

examples = traindata.rdd.map(to_example)
lrmodel = LogisticRegressionWithLBFGS.train(examples, intercept=True)

# [perry] `lrmodel` could be used for prediction
print(lrmodel.predict([6.0,12.0,594.0]))

# https://spark.apache.org/docs/latest/api/java/org/apache/spark/ml/classification/ProbabilisticClassifier.html
lrmodel.setThreshold(0.7) # cancel if prob-of-ontime < 0.7
print(lrmodel.predict([6.0,12.0,594.0])) # output 1 or 0
lrmodel.clearThreshold()
print(lrmodel.predict([6.0,12.0,594.0])) # output probability like 0.956161192349

# ...

# Evaluate model
# load test data
def eval(labelpred):
    cancel = labelpred.filter(lambda (label, pred): pred < 0.7)
    nocancel = labelpred.filter(lambda (label, pred): pred >= 0.7)
    corr_cancel = cancel.filter(lambda (label, pred): label == int(pred >= 0.7)).count()
    corr_nocancel = nocancel.filter(lambda (label, pred): label == int(pred >= 0.7)).count()
    
    cancel_denom = cancel.count()
    nocancel_denom = nocancel.count()
    if cancel_denom == 0:
        cancel_denom = 1
    if nocancel_denom == 0:
        nocancel_denom = 1
    return {'total_cancel': cancel.count(), \
            'correct_cancel': float(corr_cancel)/cancel_denom, \
            'total_noncancel': nocancel.count(), \
            'correct_noncancel': float(corr_nocancel)/nocancel_denom \
           }

# Evaluate model
lrmodel.clearThreshold() # so it returns probabilities
# example is testdata.rdd.map(to_example)
labelpred = examples.map(lambda p: (p.label, lrmodel.predict(p.features)))
print(eval(labelpred))

```

- Notes, when scale up with spark with following command, I found that increasing the number of workers didn’t actually spread out the processing to all the nodes
```
gcloud dataproc clusters create \
--num-workers=35 \ 
--num-preemptible-workers=15 \
...
```
Because spark estimate the number of partitions based on the raw input data size, if just few gigabytes is too low, you could explicit repartitioning:
```
traindata = spark.sql(trainquery).repartition(1000) 
```


## Machine Learning Classifier via TensorFlow




## Reference
- [spark Extracting, transforming and selecting features](https://spark.apache.org/docs/latest/ml-features)
- [spark ML Tuning: model selection and hyperparameter tuning](https://spark.apache.org/docs/latest/ml-tuning.html)
- [google cloud ml samples](https://github.com/GoogleCloudPlatform/cloudml-samples)
- [Wide & Deep Learning for Recommender Systems](https://arxiv.org/abs/1606.07792)
   + [Wide & Deep Learning: Memorization + Generalization with TensorFlow (TensorFlow Dev Summit 2017)](https://www.youtube.com/watch?v=NV1tkZ9Lq48)
   + Memorization vs Generalization and relevance vs diversity [知乎](https://zhuanlan.zhihu.com/p/53361519)

