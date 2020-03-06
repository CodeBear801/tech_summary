# A Deep Dive into Spark SQL's Catalyst Optimizer with Yin Huai

[video](https://www.youtube.com/watch?v=RmUn5vHlevc)


## Example 

### Original version


<img src="./resources/imgs/spark_sql_catalyst_optimizer_example_original.png" alt="spark_sql_catalyst_optimizer_example_original" width="600"/>

Analysis the query


<img src="./resources/imgs/spark_sql_catalyst_optimizer_example_analysis.png" alt="spark_sql_catalyst_optimizer_example_analysis" width="600"/>


<img src="./resources/imgs/spark_sql_catalyst_optimizer_example_analysis2.png" alt="spark_sql_catalyst_optimizer_example_analysis2" width="600"/>

### Optimized version

**Write a special query plan rule to match the cases when we join two intervals and calculate intersection directly and then put back to query plan.**


// high light code is the condition which triggers the rule:

<img src="./resources/imgs/spark_sql_catalyst_optimizer_example_optimized_1.png" alt="spark_sql_catalyst_optimizer_example_optimized_1" width="600"/>
<br/>

<img src="./resources/imgs/spark_sql_catalyst_optimizer_example_optimized_2.png" alt="spark_sql_catalyst_optimizer_example_optimized_2" width="600"/>

The result:

<img src="./resources/imgs/spark_sql_catalyst_optimizer_example_optimized_3.png" alt="spark_sql_catalyst_optimizer_example_optimized_3" width="600"/>


## Deep Dive

### Why structure API

- Structure will limit what can be impressed -> enable optimizations
- In practice, accommodate the vast majority of computations

<img src="./resources/imgs/spark_sql_catalyst_optimizer_why_structure_api.png" alt="
spark_sql_catalyst_optimizer_why_structure_api" width="600"/>

### How to take advantage of optimization opportunities?

```
Get an optimizer that automatically finds out
the most efficient plan to execute data
operations specified in the user's program
```

### Trees

abstraction of users programs

#### Expression

<img src="./resources/imgs/spark_sql_catalyst_optimizer_expressions.png" alt="spark_sql_catalyst_optimizer_expressions" width="600"/>


#### Query plan

<img src="./resources/imgs/spark_sql_catalyst_optimizer_query_plan.png" alt="spark_sql_catalyst_optimizer_query_plan" width="600"/>

Expression -> new value  
Query plan -> new dataset  
Query plan use one or more expressions


#### Transformation

<img src="./resources/imgs/spark_sql_catalyst_optimizer_transformation.png" alt="spark_sql_catalyst_optimizer_transformation" width="600"/>


A `transformation` is defined as a partial function, which is a function that is defined for a subset of its possible arguments.

<img src="./resources/imgs/spark_sql_catalyst_optimizer_transformation_example.png" alt="spark_sql_catalyst_optimizer_transformation_example" width="600"/>



#### Optimization

// predicate pushdown: t2.id>50000 only apply to t2

<img src="./resources/imgs/spark_sql_catalyst_optimizer_predicate_pushdown.png" alt="spark_sql_catalyst_optimizer_predicate_pushdown" width="600"/>



// column pruning: only need 3 columns: from t1: column id and value, from t2: column id  
// reduce IO cost(column data format)

<img src="./resources/imgs/spark_sql_catalyst_optimizer_column_pruning.png" alt="spark_sql_catalyst_optimizer_column_pruning" width="600"/>


// combine rules

<img src="./resources/imgs/spark_sql_catalyst_optimizer_combine_rules.png" alt="spark_sql_catalyst_optimizer_combine_rules" width="600"/>


Rule executor to combine multiple rules  

Two strategy to apply rule  
Fixed point: apply rules to batch again and again, until the tree is not changed anymore   
Once: apply all the rules to the same batch just once, get them all trigger  
<img src="./resources/imgs/spark_sql_catalyst_optimizer_multiple_rules.png" alt="spark_sql_catalyst_optimizer_multiple_rules" width="600"/>



#### Physical plan
<img src="./resources/imgs/spark_sql_catalyst_optimizer_logical_2_physical.png" alt="spark_sql_catalyst_optimizer_logical_2_physical" width="600"/>



## Overview
<img src="./resources/imgs/spark_sql_catalyst_optimizer_overview.png" alt="spark_sql_catalyst_optimizer_overview" width="600"/>




