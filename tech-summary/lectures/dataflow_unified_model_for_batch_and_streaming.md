# [A Unified Model for Batch and Streaming Data Processing](https://www.youtube.com/watch?v=3UfZN59Nsk8)

By Frances Perry  


## Story

<img src="resources/imgs/dataflow_frances_perry_mobile_game.png" alt="dataflow_frances_perry_mobile_game" width="400"/><br/>


Mobile game:
- Users are distributed all over the world
- User could play offline, then update status later



## Why watermarks

<img src="resources/imgs/dataflow_frances_perry_why_watermark.png" alt="dataflow_frances_perry_why_watermark" width="400"/><br/>


Watermark is a heuristic event time progress


## Key steps of Dataflow programming

<img src="resources/imgs/dataflow_frances_perry_four_steps.png" alt="dataflow_frances_perry_four_steps" width="400"/><br/>

- What results are calculated? = transformations
- Where in event-time are results calculated? = windowing
- When in processing-time are results materialized? = watermarks + triggers
- How do refinements of results relate? = accumulation


### What
<img src="resources/imgs/dataflow_frances_perry_what.png" alt="dataflow_frances_perry_what" width="400"/><br/>


### Where
<img src="resources/imgs/dataflow_frances_perry_where.png" alt="dataflow_frances_perry_where" width="400"/><br/>


Session means a period time user is active.  More info about [session-windows](https://cloud.google.com/dataflow/docs/concepts/streaming-pipelines#session-windows)

### When

How the arriving time into system will affect the processing result

`Watermarks`: A watermark is a notion of input completeness with respect to event times. A watermark with a value of time X makes the statement: “all input data with event times less than X have been observed.”   

`Triggers`: A trigger is a mechanism for declaring when the output for a window should be materialized relative to some external signal.  

<img src="resources/imgs/dataflow_frances_perry_when.png" alt="dataflow_frances_perry_when" width="400"/><br/>


Why hubristic value for `watermark`  

<img src="https://user-images.githubusercontent.com/16873751/84205333-b9901a80-aa61-11ea-8e6c-d5b043d50fba.png" alt="dataflow_frances_perry_when" width="400"/><br/>
- To balance `resource taken` for computation and correctness of all data
   + too slow
   + too fast
- Take `9` as a example, its generated around 12:01~12:02, but arrives for processing around 12:08~12:09, when calculating window of 12:00~12:02, if we want to taken `9` into consideration, the resource be used for calculating this window need to be hold for 7 more minutes(12:09 - 12:02) or more.



#### Example
<img src="resources/imgs/dataflow_frances_perry_example.png" alt="dataflow_frances_perry_example" width="400"/><br/>

<img src="resources/imgs/dataflow_frances_perry_example_fixed_window.png" alt="dataflow_frances_perry_example_fixed_window" width="400"/><br/>

<img src="resources/imgs/dataflow_frances_perry_example_fixed_window2.png" alt="dataflow_frances_perry_example_fixed_window2" width="400"/><br/>

<img src="resources/imgs/dataflow_frances_perry_example_trigger_at_watermark.png" alt="dataflow_frances_perry_example_trigger_at_watermark" width="400"/><br/>

<img src="resources/imgs/dataflow_frances_perry_example_trigger_at_watermark2.png" alt="dataflow_frances_perry_example_trigger_at_watermark2" width="400"/><br/>


- `9` has been dropped
- waiting to watermark to trigger
- 
<img src="resources/imgs/dataflow_frances_perry_example_trigger_at_watermark3.png" alt="dataflow_frances_perry_example_trigger_at_watermark3" width="400"/><br/>

<img src="resources/imgs/dataflow_frances_perry_example_trigger_at_watermark4.png" alt="dataflow_frances_perry_example_trigger_at_watermark4" width="400"/><br/>


### How
<img src="resources/imgs/dataflow_frances_perry_how.png" alt="dataflow_frances_perry_how" width="400"/><br/>

<img src="resources/imgs/dataflow_frances_perry_how_code.png" alt="dataflow_frances_perry_how_code" width="400"/><br/>

<img src="resources/imgs/dataflow_frances_perry_how_code2.png" alt="dataflow_frances_perry_how_code2" width="400"/><br/>



<img src="resources/imgs/dataflow_frances_perry_4step_summary.png" alt="dataflow_frances_perry_4step_summary" width="400"/><br/>



## Demo

Fixed bounding input data -> output aggregation for all batch data

<img src="resources/imgs/dataflow_frances_perry_demo1.png" alt="dataflow_frances_perry_demo1" width="800"/><br/>


Window
- fixed window
- session window(processing time window)
- 
<img src="resources/imgs/dataflow_frances_perry_demo2.png" alt="dataflow_frances_perry_demo2" width="800"/><br/>


Fixed window

<img src="resources/imgs/dataflow_frances_perry_demo3.png" alt="dataflow_frances_perry_demo3" width="400"/><br/>


Session Window

<img src="resources/imgs/dataflow_frances_perry_demo4.png" alt="dataflow_frances_perry_demo4" width="400"/><br/>



## Reference 
- https://cs.stanford.edu/~matei/courses/2015/6.S897/slides/dataflow.pdf
- https://beam.apache.org/documentation/programming-guide/#overview
- https://github.com/tshauck/DataflowJavaSDK-examples/tree/master/src/main/java8/com/google/cloud/dataflow/examples/complete/game
- https://github.com/jlewi/dataflow/blob/master/dataflow/src/main/java/sessions/SlidingWindowExample.java