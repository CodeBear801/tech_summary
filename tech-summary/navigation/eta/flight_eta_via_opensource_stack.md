# Estimate flight ETA via Open Source Stack


### Extracting features via spark 

[code](https://github.com/CodeBear801/Agile_Data_Code_2/blob/master/ch07/extract_features.py#L81)
```python
# Load the on-time parquet file
on_time_dataframe = spark.read.parquet('data/on_time_performance.parquet')
# [perry] new api: createTempView
on_time_dataframe.registerTempTable("on_time_performance")

# Select a few features of interest
simple_on_time_features = spark.sql("""
SELECT
  FlightNum,
  FlightDate,
  DayOfWeek,
  DayofMonth AS DayOfMonth,
  ...
FROM on_time_performance
""")
simple_on_time_features.show()

# [perry] data cleaning
# filled_on_time_features = simple_on_time_features.filter(filter_func)

# [perry] format convert and feature generation
# timestamp_features = filled_on_time_features.rdd.map(enhancement_func)

# Explicitly sort the data and keep it sorted throughout. Leave nothing to chance.
# [perry] sort before vectorization explicitly to guarantee stability of the result 
sorted_features = timestamp_df.sort(
  timestamp_df.DayOfYear,
  timestamp_df.Carrier,
  timestamp_df.Origin,
  timestamp_df.Dest,
  timestamp_df.FlightNum,
  timestamp_df.CRSDepTime,
  timestamp_df.CRSArrTime,
)

# [perry] output result to json

```

### regression via scikit-learn
[code](https://github.com/CodeBear801/Agile_Data_Code_2/blob/master/ch07/Predicting%20flight%20delays%20with%20sklearn.ipynb)

```python
# Load
# Load and check the size of our training data. 
training_data = utils.read_json_lines_file('../data/simple_flight_delay_features.jsonl')

# Sample
# We need to sample our data to fit into RAM
# [perry]decrease amount of total data size
training_data = np.random.choice(training_data, 1000000) # 'Sample down to 1MM examples'

# Vectorization
# Use DictVectorizer to convert feature dicts to vectors
vectorizer = DictVectorizer()
training_vectors = vectorizer.fit_transform(training_data)
# [perry] https://datascience.stackexchange.com/questions/12321/difference-between-fit-and-fit-transform-in-scikit-learn-models


# prepare for cross validation
X_train, X_test, y_train, y_test = train_test_split(
  training_vectors,
  results_vector,
  test_size=0.1, # perry: training data 90%, testing data 10%
  random_state=43
)

# training
regressor = LinearRegression()
regressor.fit(X_train, y_train)

# predict
predicted = regressor.predict(X_test)

# evaluation

# Median absolute error is the median of all absolute differences between the target and the prediction.
# Less is better, more indicates a high error between target and prediction.
medae = median_absolute_error(y_test, predicted)

# R2 score is the coefficient of determination. Ranges from 1 - 0, 1.0 is best, 0.0 is worst.
# Measures how well future samples are likely to be predicted.
r2 = r2_score(y_test, predicted)

```

### Classifier via Spark MLlib

[code](https://render.githubusercontent.com/view/ipynb?commit=46dc4e5514d0189fff1baaffb9ab817ba2aff19f&enc_url=68747470733a2f2f7261772e67697468756275736572636f6e74656e742e636f6d2f436f6465426561723830312f4167696c655f446174615f436f64655f322f343664633465353531346430313839666666316261616666623961623831376261326166663139662f636830372f4d616b696e675f50726564696374696f6e732e6970796e62&nwo=CodeBear801%2FAgile_Data_Code_2&path=ch07%2FMaking_Predictions.ipynb&repository_id=208912886&repository_type=Repository#Bucketizing-a-Continuous-Variable-for-Classification)

```py
# Load data with specified schema
schema = StructType([
  StructField("ArrDelay", FloatType(), True),       # "ArrDelay":5.0
  StructField("CRSArrTime", TimestampType(), True), # "CRSArrTime":"2015-12..."
  StructField("CRSDepTime", TimestampType(), True), # "CRSDepTime":"2015-12..."
  StructField("Carrier", StringType(), True),       # "Carrier":"WN"
  StructField("DayOfMonth", IntegerType(), True),   # "DayOfMonth":31
  StructField("DayOfWeek", IntegerType(), True),    # "DayOfWeek":4
  StructField("DayOfYear", IntegerType(), True),    # "DayOfYear":365
  StructField("DepDelay", FloatType(), True),       # "DepDelay":14.0
  StructField("Dest", StringType(), True),          # "Dest":"SAN"
  StructField("Distance", FloatType(), True),       # "Distance":368.0
  StructField("FlightDate", DateType(), True),      # "FlightDate":"2015-12..."
  StructField("FlightNum", StringType(), True),     # "FlightNum":"6109"
  StructField("Origin", StringType(), True),        # "Origin":"TUS"
])

features = spark.read.json(
  "../data/simple_flight_delay_features.json", 
  schema=schema
)

# [perry] for how to decide bucket range, please go to original code
# excellent description in Bucketizing a Continuous Variable for Classification
# Try with multiple bucket range and visualize histogram
# target is generating a balanced category

# Use pysmark.ml.feature.Bucketizer to bucketize ArrDelay
splits = [-float("inf"), -15.0, 0, 15.0, 30.0, float("inf")]
bucketizer = Bucketizer(
  splits=splits,
  inputCol="ArrDelay",
  outputCol="ArrDelayBucket"
)
ml_bucketized_features = bucketizer.transform(features_with_route)

# Vectorized features

# Turn category fields(string) into categoric feature vectors, then drop # intermediate fields
# [perry] https://spark.apache.org/docs/latest/ml-features#stringindexer
#        new column of original_columnname_index has been generated for each string
for column in ["Carrier", "DayOfMonth", "DayOfWeek", "DayOfYear",
               "Origin", "Dest", "Route"]: 
               string_indexer = StringIndexer(
                                     inputCol=column,
                                     outputCol=column + "_index" )
ml_bucketized_features = string_indexer.fit(ml_bucketized_features)\ 
                                       .transform(ml_bucketized_features)

# [Perry] https://spark.apache.org/docs/2.0.0/api/python/pyspark.ml.html#pyspark.ml.feature.VectorAssembler
#         https://spark.apache.org/docs/2.0.0/api/python/pyspark.ml.html#pyspark.ml.linalg.Vector 
# VectorAssembler is used to combine numeric and index columns into a single feature vector

# Handle continuous numeric fields by combining them into one feature vector
numeric_columns = ["DepDelay", "Distance"]
index_columns = ["Carrier_index", "DayOfMonth_index",
                   "DayOfWeek_index", "DayOfYear_index", "Origin_index",
                   "Origin_index", "Dest_index", "Route_index"]
vector_assembler = VectorAssembler(
  inputCols=numeric_columns + index_columns,
  outputCol="Features_vec"
)
final_vectorized_features = vector_assembler.transform(ml_bucketized_features)

# Drop the index columns
for column in index_columns:
  final_vectorized_features = final_vectorized_features.drop(column)



# Test/train split
training_data, test_data = final_vectorized_features.randomSplit([0.8, 0.2])

# training
rfc = RandomForestClassifier(
featuresCol="Features_vec", labelCol="ArrDelayBucket",
      maxBins=4657 
    )
# [Perry]maxBins is an experiment value
# Fir default value 32, for initial run will throw an exception which indicates
# that unique value for one feature more than 32 and will suggest the value of 4657
model = rfc.fit(training_data)


# Evaluation

# Evaluate model using test data
predictions = model.transform(test_data)
evaluator = MulticlassClassificationEvaluator(
      labelCol="ArrDelayBucket", metricName="accuracy"
    )
accuracy = evaluator.evaluate(predictions)
```

### How to improve prediction model

Two ways to improve
- hyperparameter tuning https://spark.apache.org/docs/latest/ml-tuning.html
- feature selection https://www.analyticsvidhya.com/blog/2016/12/introduction-to-feature-selection-methods-with-an-example-or-how-to-select-the-right-variables/

how to evaluate improvement  
[MulticlassClassificationEvaluator from spark](https://spark.apache.org/docs/latest/api/python/pyspark.ml.html#pyspark.ml.evaluation.MulticlassClassificationEvaluator) offers four metrics: accuracy, weighted precision, weighted recall, and f1  
- `Accuracy`: the number of correct predictions divided by the number of predictions
- `Precision`: a measure of how useful the result is
- `Recall` describes how complete the results are.
- `f1` score incorporates both precision and recall to deter‐ mine overall quality

The importance of a feature is a measure of how important that feature was in contributing to the accuracy of the model.  If we know how important a feature is, we can use this clue to make changes that increase the accuracy of the model, such as removing unimportant features and trying to engineer features similar to those that are most important.  The state of the art for many classification and regression tasks is a [gradient boosted decision tree](https://en.wikipedia.org/wiki/Gradient_boosting)

