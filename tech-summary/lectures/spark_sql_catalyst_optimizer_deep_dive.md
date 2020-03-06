# A Deep Dive into Spark SQL's Catalyst Optimizer with Yin Huai

[video](https://www.youtube.com/watch?v=RmUn5vHlevc)


## Example 

### Original version


<img src="./resources/imgs/spark_sql_catalyst_optimizer_example_original.png" alt="spark_sql_catalyst_optimizer_example_original" width="600"/>

Analysis the query


<img src="./resources/imgs/spark_sql_catalyst_optimizer_example_analysis.png" alt="spark_sql_catalyst_optimizer_example_analysis" width="600"/>


<img src="./resources/imgs/spark_sql_catalyst_optimizer_example_analysis2.png" alt="spark_sql_catalyst_optimizer_example_analysis2" width="400"/>

### Optimized version

**Write a special query plan rule to match the cases when we join two intervals and calculate intersection directly and then put back to query plan.**


// high light code is the condition which triggers the rule:

<img src="./resources/imgs/spark_sql_catalyst_optimizer_example_optimized_1.png" alt="spark_sql_catalyst_optimizer_example_optimized_1" width="600"/>
<br/>

<img src="./resources/imgs/spark_sql_catalyst_optimizer_example_optimized_2.png" alt="spark_sql_catalyst_optimizer_example_optimized_2" width="600"/>

The result:

<img src="./resources/imgs/spark_sql_catalyst_optimizer_example_optimized_3.png" alt="spark_sql_catalyst_optimizer_example_optimized_3" width="400"/>

