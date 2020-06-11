# Estimate flight ETA via Open Source Stack


### Extracting features

via spark [code](https://github.com/CodeBear801/Agile_Data_Code_2/blob/master/ch07/extract_features.py#L81)
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

# R2 score is the coefficient of determination. Ranges from 1-0, 1.0 is best, 0.0 is worst.
# Measures how well future samples are likely to be predicted.
r2 = r2_score(y_test, predicted)


```
