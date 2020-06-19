
## Why XGBoost
- https://github.com/dmlc/xgboost
- http://datascience.la/xgboost-workshop-and-meetup-talk-with-tianqi-chen/

![image](https://user-images.githubusercontent.com/16873751/85156321-c1dd1800-b20e-11ea-8d2a-a4b1cb908f3b.png)

decision tree, weak learner(individually they are quite inaccurate, but slightly better when work together)

![image](https://user-images.githubusercontent.com/16873751/85156771-63646980-b20f-11ea-83e9-dcb2341a41cc.png)

second tree must provide positive effort when combine with first tree


## PCA

- https://setosa.io/ev/principal-component-analysis/

![image](https://user-images.githubusercontent.com/16873751/85156926-9c9cd980-b20f-11ea-859a-5da8e06157a5.png)


- Dimension for original data is reasonable for human beings, but might not friendly for decision tree

- Try to represent data with different `coordinate system`

- Make the dimension the principal component with most variation - (Minimize difference between min value and max value)

## Why Parquet

- Dremel paper https://storage.googleapis.com/pub-tools-public-publication-data/pdf/36632.pdf
- Google Dremel 原理 - 如何能3秒分析1PB https://www.twblogs.net/a/5b82787e2b717766a1e868d0
- Approaching NoSQL Design in DynamoDB https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/bp-general-nosql-design.html#bp-general-nosql-design-approach
- Spark + Parquet In Depth https://www.youtube.com/watch?v=_0Wpwj_gvzg&t=1501s

### Target of column based data organize
- Only load needed data
- Organize data storage based on needs(query, filter)

Sample data input   
![image](https://user-images.githubusercontent.com/16873751/85159671-ac69ed00-b212-11ea-9f06-0623b2a568da.png)


Flat data vs. Nested data
![image](https://user-images.githubusercontent.com/16873751/85159736-bb509f80-b212-11ea-877d-e6b7c68004f6.png)

Option on Parquet with `scala`
```java
val flatDF = sc.read.option("delimiter", "\t")
                    .option("header", "true")
                    .csv(flatInput)
                    .rdd
                    .map(r => transformRow(r))
                    .toDF

flatDF.write.option("compression", "snappy")
            .parquet(flatOutput)

var nestedDF = sc.read.json(nestedInput)

nestedDF.write.option("compression", "snappy")
              .parquet(nestedOutput)
```

### How to avoid reading un-related column?

![image](https://user-images.githubusercontent.com/16873751/85160645-72e5b180-b213-11ea-9d54-3a13eff17aef.png)

By record data column oriented, we could use different way to compress the data  
Incrementally(record diff), or use dictionary  

More details could be found in [parquet encoding](https://github.com/apache/parquet-format/blob/master/Encodings.md)

### How to save space

![image](https://user-images.githubusercontent.com/16873751/85160785-a88a9a80-b213-11ea-9a31-59ac44708bae.png)

![image](https://user-images.githubusercontent.com/16873751/85160804-afb1a880-b213-11ea-842e-b1198150ac43.png)


### Parquet file internal

tree structure

![image](https://user-images.githubusercontent.com/16873751/85160861-ca841d00-b213-11ea-83d6-77a347f35734.png)

[parquet file format](https://github.com/apache/parquet-format)  

- file metadata: schema, num of rows
- Row group: each time when write massive data, we are not going to write them all together but piece by piece
- Column Chunk: consider each column individually(if whole column is no, directly skip)
- Page header: size
- Page: record meta data like count, max, min to help quick filter

### The power of partition
```java
dataFrame.write
         .partitionBy("Year", "Month", "Day", "Hour")
         .parquet(outputFile)
```


