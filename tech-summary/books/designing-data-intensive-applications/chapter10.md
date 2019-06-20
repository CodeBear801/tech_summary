<!-- TOC -->
- [Batch processing](#batch-processing)
  - [Keywords](#keywords)
  - [Questions](#questions)
  - [Notes](#notes)
    - [MapReduce](#mapreduce)
    - [Dataflow](#dataflow)
    - [Graph processing](#graph-processing)

# Batch processing

## Keywords

## Questions
- Why need data flow like spark, what's the issue of mapreduce?
- Mapreduce example: build index, distribute grep, distribute sort, classifier and recommender
- How could partition help sort-merge joins, broadcast hash joins, partitioned hash joins

## Notes

### MapReduce
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
- What about graph processing, e.g. processing data using graph algorithms?
- Dataflow engines have implemented this feature using the bulk sychronous parallel model of computation. This was populated by Google Pregel (Google again...).
- Pregel has 3 components: master, worker, aggregator
- The insight is that most of these algorithms can be implemented by processing one node at a time and "walking" the graph.
- This algorithm archetype is known as a transitive closure.
- In BSP nodes are processed in stages. At each stage you process the nodes that match some condition, and evaluate what nodes to step to next.
- When you run out of node to jump to you stop.
- It is possible to parallelize this algorithm across multiple partitions. Ideally you want to partition on the neighborhoods you are going to be within during the walks, but this is hard to do, so most schemes just partition the graph arbitrarily.
- This creates unavoidable message overhead, when nodes of interest are on different machines.
Ongoing area of research.

