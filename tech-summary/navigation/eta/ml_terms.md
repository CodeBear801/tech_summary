
## Predictive Analytics

`Dependent features(自变特征)` are the values we are trying to predict.   
`Independent features(因变特征)` are features describing the things we want to predict that relate to the dependent features.  
For instance, for the training data used for predicting flight delays is our atomic records: our flight delay records. A flight with its delay is an example of a record with a dependent variable. Our independent features are other things we can associate with flights: departure delay, the airline, the origin and destination.  

### regression and classification

There are two ways to approach most predictions: `regression` and `classification`
`regression` takes examples made up of features as input and produces a numeric out‐put.   
`Classification` takes examples as input and produces a categorical classification.  

```
// definition of linear regression
// https://stattrek.com/statistics/dictionary.aspx?definition=Regression
In a cause and effect relationship, the independent variable is the cause, and the dependent variable is the effect. Least squares linear regression is a method for predicting the value of a dependent variable Y, based on the value of an independent variable X.

// More info: Linear Regression in Excel https://projects.ncsu.edu/labwrite//res/gt/gt-reg-home.html
```


### feature

A `feature` is what it sounds like: a feature of an example.  Two or more features make up the training data of a statistical prediction—two being the mini‐ mum because one field is required as the one to predict, and at least one additional feature is required to make an inference about in order to create a prediction.  


### [feature engineering](https://developers.google.com/machine-learning/crash-course/representation/feature-engineering)

Mapping raw data to features

<img src="https://user-images.githubusercontent.com/16873751/84330964-bcf8d400-ab3d-11ea-8dd9-54297ca08e2d.png" alt="feature_1" width="400"/><br/>

`one-hot encoding` and `multi-hot encoding`

<img src="https://user-images.githubusercontent.com/16873751/84331012-e0238380-ab3d-11ea-97f8-81f246d79cd1.png" alt="feature_1" width="400"/><br/>

<img src="https://user-images.githubusercontent.com/16873751/84331181-65a73380-ab3e-11ea-986e-4f538cc360ea.png" alt="feature_1" width="400"/><br/>

- more info: [Text Vectorization and Transformation Pipelines](https://www.oreilly.com/library/view/applied-text-analysis/9781491963036/ch04.html)

[`sparse representation`](https://developers.google.com/machine-learning/glossary#sparse_representation) to decrease resource usage  
<img src="https://user-images.githubusercontent.com/16873751/84331022-ed407280-ab3d-11ea-9559-714f9b2c9a4d.png" alt="feature_1" width="400"/><br/>


## cross validation
[Cross-validation](https://en.wikipedia.org/wiki/Cross-validation_(statistics)) means we split our data into test and training sets, and then train the model on the training set before testing it on the test set. Cross-validation prevents `overfitting`, which is when a model seems quite accurate but fails to actually predict future events well.


### Median absolute deviation

[Median absolute error](https://en.wikipedia.org/wiki/Median_absolute_deviation) is the median of all absolute differences between the target and the prediction.  **Less is better,** more indicates a high error between target and prediction.


[The median absolute deviation(MAD)](https://www.statisticshowto.com/median-absolute-deviation/) is a robust measure of how spread out a set of data is. The variance and standard deviation are also measures of spread, but they are more affected by extremely high or extremely low values and non normality. If your data is normal, the standard deviation is usually the best choice for assessing spread. However, if your data isn’t normal, the MAD is one statistic you can use instead. 

```
Example: Find the MAD of the following set of numbers: 3, 8, 8, 8, 8, 9, 9, 9, 9.
Step 1: Find the median. The median for this set of numbers is 8.

Step 2: Subtract the median from each x-value using the formula |yi – median|.
|3 – 8| = 5
|8 – 8| = 0
|8 – 8| = 0
|8 – 8| = 0
|8 – 8| = 0
|9 – 8| = 1
|9 – 8| = 1
|9 – 8| = 1
|9 – 8| = 1

Step 3: find the median of the absolute differences. The median of the differences (0,0,0,0,1,1,1,1,5) is 1.
```

<img src="https://user-images.githubusercontent.com/16873751/84334144-82e00000-ab46-11ea-87ed-82d7ef1bd0b0.png" alt="feature_1" width="400"/><br/>


## Reference
- [For all resource imported from other side, please go to here for the source link](https://github.com/CodeBear801/tech_summary/issues/2)
