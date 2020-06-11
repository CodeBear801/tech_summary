
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

[`sparse representation`](https://developers.google.com/machine-learning/glossary#sparse_representation) to decrease resource usage  
<img src="https://user-images.githubusercontent.com/16873751/84331022-ed407280-ab3d-11ea-9559-714f9b2c9a4d.png" alt="feature_1" width="400"/><br/>


## Reference
- [For all resource imported from other side, please go to here for the source link](https://github.com/CodeBear801/tech_summary/issues/2)
