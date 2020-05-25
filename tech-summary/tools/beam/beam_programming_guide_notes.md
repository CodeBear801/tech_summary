# [Apache Beam Programming Guide](https://beam.apache.org/documentation/programming-guide/)

## Window - Where in event time

[Windowing](https://beam.apache.org/documentation/programming-guide/#windowing) subdivides a PCollection according to the timestamps of its individual elements.   

<img src="https://beam.apache.org/images/windowing-pipeline-bounded.svg" alt="windowing-pipeline-bounded" width="600"/>

### Fixed time windows

A fixed time window represents a consistent duration, non overlapping time interval in the data stream.  

<img src="https://beam.apache.org/images/fixed-time-windows.png" alt="[windowing-pipeline-bounded](https://beam.apache.org/images/fixed-time-windows.png)" width="600"/>

```java
    PCollection<String> fixedWindowedItems = items.apply(
        Window.<String>into(FixedWindows.of(Duration.standardSeconds(60))));
```

### Sliding time windows

A sliding time window also represents time intervals in the data stream; however, sliding time windows can overlap.  

<img src="https://beam.apache.org/images/sliding-time-windows.png" alt="https://beam.apache.org/images/sliding-time-windows.png" width="600"/>


```java
  PCollection<String> slidingWindowedItems = items.apply(
        Window.<String>into(SlidingWindows.of(Duration.standardSeconds(30)).every(Duration.standardSeconds(5))));

```

### Session windows
A session window function defines windows that contain elements that are within a certain gap duration of another element.  Used to merge small gap of data into same window.

<img src="https://beam.apache.org/images/session-windows.png" alt="https://beam.apache.org/images/session-windows.png" width="600"/>


```java
PCollection<String> sessionWindowedItems = items.apply(
        Window.<String>into(Sessions.withGapDuration(Duration.standardSeconds(600))));
```


Beam’s default windowing configuration tries to determine when all data has arrived (based on the type of data source) and then advances the watermark past the end of the window. This default configuration does not allow late data. Triggers allow you to modify and refine the windowing strategy for a PCollection. You can use triggers to decide when each individual window aggregates and reports its results, including how the window emits late elements.


## Watermark - When in processing time

[watermarks](https://beam.apache.org/documentation/programming-guide/#watermarks-and-late-data): there is a certain amount of lag between the time a data event occurs and the time the actual data element gets processed at any stage in your pipeline.  Data isn’t always guaranteed to arrive in a pipeline in time order, or to always arrive at predictable intervals. Beam tracks a watermark, which is the system’s notion of when all data in a certain window can be expected to have arrived in the pipeline. Once the watermark progresses past the end of a window, any further element that arrives with a timestamp in that window is considered late data.


```java
    PCollection<String> fixedWindowedItems = items.apply(
        Window.<String>into(FixedWindows.of(Duration.standardMinutes(1)))
              .withAllowedLateness(Duration.standardDays(2)));
```

When you set .withAllowedLateness on a PCollection, that allowed lateness propagates forward to **any subsequent PCollection derived** from the first PCollection you applied allowed lateness to.

## Tigger - When in processing time and how do refinements relate.

[triggers](https://beam.apache.org/documentation/programming-guide/#triggers), determines when to emit the results of aggregation as unbounded data arrives.  

- Event time triggers
- Processing time triggers
- Data-driven triggers
- Composite triggers

Window accumulation modes: Accumulating, Discarding

Example 1
- On Beam’s estimate that all the data has arrived (the watermark passes the end of the window)
- Any time late data arrives, after a ten-minute delay
- After two days, we assume no more data of interest will arrive, and the trigger stops executing

```java
  .apply(Window
      .configure()
      .triggering(AfterWatermark
           .pastEndOfWindow()
           .withLateFirings(AfterProcessingTime
                .pastFirstElementInPane()
                .plusDelayOf(Duration.standardMinutes(10))))
      .withAllowedLateness(Duration.standardDays(2)));
```

Example 2

 a simple composite trigger that fires whenever the pane has at least 100 elements, or after a minute.

```java
Repeatedly.forever(AfterFirst.of(
      AfterPane.elementCountAtLeast(100),
      AfterProcessingTime.pastFirstElementInPane().plusDelayOf(Duration.standardMinutes(1))))

```


