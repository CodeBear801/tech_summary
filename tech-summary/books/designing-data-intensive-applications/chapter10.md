<!-- TOC -->
- [Batch processing](#batch-processing)
  - [Keywords](#keywords)
  - [Questions](#questions)
  - [Notes](#notes)
    - [Batch with unix tools](#batch-with-unix-tools)
    - [MapReduce](#mapreduce)
    - [Dataflow](#dataflow)
    - [Graph processing](#graph-processing)
  - [More Info](#more-info)

# Batch processing

## Keywords

## Questions
- Why need data flow like spark, what's the issue of mapreduce?
- Mapreduce example: build index, distribute grep, distribute sort, classifier and recommender
- How could partition help sort-merge joins, broadcast hash joins, partitioned hash joins

## Notes

### Batch with unix tools
get the five most popular pages on your site
```bash
cat /var/log/nginx/access.log |
  awk '{print $7}' |
  sort             |
  uniq -c          |
  sort -r -n       |
  head -n 5        |
```


### MapReduce

- [My notes on MapReduce paper](../../papers/mapreduce.md)
- MapReduce is a programming framework with which you can write code to process large datasets in a distributed filesystem like GFS
   + Read a set of input files, and break it up into records.
   + Call the mapper function to extract a key and value from each input record.
   + Sort all of the key-value pairs by key.
   + Call the reducer function to iterate over the sorted key-value pairs.
       * Mapper: Called once for every input record, and its job is to extract the key and value from the input record.
       * Reducer: Takes the key-value pairs produced by the mappers, collects all the values belonging to the same key, and calls the reducer with an iterator over that collection of values.


How to handle hot keys?  
- In an example of a social network, small number of celebrities may have many millions of followers. Such disproportionately active database records are known as linchpin objects or hot keys.
- A single reducer can lead to significant skew that is, one reducer that must process significantly more records than the others.
- The `skewed join` method in Pig first runs a sampling job to determine which keys are hot and then records related to the hot key need to be replicated to all reducers handling that key.
- Handling the hot key over several reducers is called `shared join method`. In Crunch is similar but requires the hot keys to be specified explicitly.
- Hive's skewed join optimization requires hot keys to be specified explicitly and it uses `map-side join`. If you can make certain assumptions about your input data, it is possible to make joins faster. A MapReducer job with no reducers and no sorting, each mapper simply reads one input file and writes one output file

How to handle skew join in spark?
An example in spark from https://stackoverflow.com/questions/40373577/skewed-dataset-join-in-spark  
Say you have to join two tables A and B on A.id=B.id. Lets assume that table A has skew on id=1.  

- Approach 1:
  - Break your query/dataset into 2 parts - one containing only skew and the other containing non skewed data. In the above example.
  - The first query will not have any skew, so all the tasks of ResultStage will finish at roughly the same time.
  - If we assume that B has only few rows with B.id = 1, then it will fit into memory. So Second query will be converted to a broadcast join. This is also called `Map-side join` in Hive.  Reference: https://cwiki.apache.org/confluence/display/Hive/Skewed+Join+Optimization
  - The partial results of the two queries can then be merged to get the final results.
```sql
1. select A.id from A join B on A.id = B.id where A.id <> 1;
2. select A.id from A join B on A.id = B.id where A.id = 1 and B.id = 1;
```

- Approach 2:
  - randomize the join key by appending extra column.  Like adding salt, dividing the data into 100 bins for large df and replicating the small df 100 times.
  - Add a column in the larger table (A), say skewLeft and populate it with random numbers between 0 to N-1 for all the rows.
  - Add a column in the smaller table (B), say skewRight. Replicate the smaller table N times. So values in new skewRight column will vary from 0 to N-1 for each copy of original data. For this, you can use the explode sql/dataset operator.
  - join the 2 datasets/tables with join condition.  Reference: https://datarus.wordpress.com/2015/05/04/fighting-the-skew-in-spark/
  
```sql
*A.id = B.id && A.skewLeft = B.skewRight*
```




How to build index for full-text search?
If you need to perform a full-text search, a batch process is very effective way of building indexes: the mappers partition the set of documents as needed, each reducer builds the index for its partition, and the index files are written to the distributed filesystem. It parallelisms very well.   
Google's original use of MapReduce was to build indexes for its search engine. Hadoop MapReduce remains a good way of building indexes for Lucene/Solr.  


- We can place the nearness data processing systems on a continuum, between online systems on one end and batch processing systems on the other end (with stream processing as an intermediate; another chapter).
- Batch processing systems process data on a scheduled or as-needed basis, instead of immediate basis of an online systme.
- Thus the concerns are very different. Latency doesn't matter. We design for total application success or total failure. 
- Throughput is the most important measurement.
Batch processing is really the original programming use case, harkening back to the US Census counting card machine days!

- A significant issue in MapReduce is skew. If keys are partitioned amongst reducers naively, hot keys, as typical in a Zipf-distributed system (e.g. celebrities on Twitter), will result in very bad tail-reducer performance.
- Some on-top-of-Hadoop systems, like Apache Pig, provide a skew join facility for working with such keys.
- These keys get randomly sub-partitioned amongst the reducers to distribute the load.
- The default utilization is to perform joins on the reducer side. It's also possible to perform a mapper-side join.
You do not want to push the final producer of a MapReduce job to a database via insert operations as this is slow. It's better to just build a new database in place. A number of databases designed with batch processing in mind provide this feature (see e.g. LevelDB).


### Dataflow
- MapReduce is becoming ever more historical.
- There are better processing models, that solve problems that were discovered with MapReduce.
- One big problem with MapReduce has to do with its state materialization. In a chained MapReduce, in between every map-reduce there is a write to disc.
- Another one, each further step in the chain must wait for all of the previous job to finish before it can start its work.
- Another problem is that mappers are often redundant, and could be omitted entirely.
- The new approach is known as dataflow engines. Spark etcetera.
- Dataflow engines build graphs over the entire data workflow, so they contain all of the state (instead of just parts of it, as in MapReduce).  DAG
- They generalize the map-reduce steps to generic operators. Map-reducers are a subset of operators.
- The fact that the dataflow engines are aware of all steps allows them to optimize movement between steps in ways that were difficult to impossible with MapReduce.
- For example, if the data is small they may avoid writing to disc at all.
- Sorting is now optional, not mandatory. Steps that don't need to sort can omit that operation entirely, saving a ton of time.
- Since operators are generic, you can often combine what used to be several map-reduce steps into one operation. This saves on all of the file writes, all of the sorts, and all of the overhead in between those moves.
- It also allows you to more easily express a broad range of computational ideas. E.g. to perform some of the developer experience optimizations that the API layers that were built on top of Hadoop performed.
- On the other hand, since there may not be intermediate materialized state to back up on, in order to retain fault tolerance dataflow introduces the requirement that computations be deterministic.  
In practice, there are a lot of sneaky ways in which non-determinism may sneak into your processing.

### Graph processing
- [My notes on Pregel paper](../../papers/pregel.md)
- When look at graphs in batch processing context, the goal is to perform some kind of offline processing or analysis on an entire graph. This need often arises in machine learning applications such as recommendation engines, or in ranking systems.  "repeating until done" cannot be expressed in plain MapReduce as it runs in a single pass over the data and some extra trickery is necessary.
- An optimization for batch processing graphs, the bulk synchronous parallel (BSP) has become popular.
    + One vertex can "send a message" to another vertex, and typically those messages are sent along the edges in a graph.  The difference from MapReduce is that a vertex remembers its state in memory from one iteration to the next.  The fact that vertices can only communicate by message passing helps improve the performance of Pregel jobs, since messages can be batched.
    + Fault tolerance is achieved by periodically checkpointing the state of all vertices at the end of an interation.
    + The framework may partition the graph in arbitrary ways.
- If your graph can fit into memory on a single computer, it's quite likely that a single-machine algorithm will outperform a distributed batch process. If the graph is too big to fit on a single machine, a distributed approach such as Pregel is unavoidable.


## More Info
- [My notes on MapReduce paper](../../papers/mapreduce.md)
- [My notes on Apache Spark](../../tools/spark_index.md)
- [My notes on Pregel paper](../../papers/pregel.md)
- My notes on [flumejava](../../papers/flumejava.md) [rdd](../../papers/rdd.md)
- [My notes on gfs paper](../../papers/gfs.md)


