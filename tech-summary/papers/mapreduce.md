
How MapReduce scales 
- N computers gets you Nx throughput.
    - Assuming M and R are >= N (i.e., lots of input files and map output keys).
    - Maps()s can run in parallel, since they don't interact.  Same for Reduce()s.
    - The only interaction is via the "shuffle" in between maps and reduces.
- So you can get more throughput by buying more computers.
    - Rather than special-purpose efficient parallelizations of each application.


How does detailed design reduce effect of slow network?
- Map input is read from GFS replica on local disk, not over network.
- Intermediate data goes over network just once.  Map worker writes to local disk, not GFS.
- Intermediate data partitioned into files holding many keys. (Q: Why not stream the records to the reducer (via TCP) as they are being produced by the mappers?)

How do they get good load balance?
Critical to scaling -- bad for N-1 servers to wait for 1 to finish.  But some tasks likely take longer than others.
[packing variable-length tasks into workers]
Solution: 
- many more tasks than workers.
- Master hands out new tasks to workers who finish previous tasks.
- So no task is so big it dominates completion time (hopefully).
- So faster servers do more work than slower ones, finish abt the same time.

fault tolerance
what if a server crashes during a MR job?

MR re-runs just the failed Map()s and Reduce()s.
MR requires them to be pure functions:
  they don't keep state across calls,
  they don't read or write files other than expected MR inputs/outputs,
  there's no hidden communication among tasks.
So re-execution yields the same output.
The requirement for pure functions is a major limitation of MR compared to other parallel programming schemes.
But it's critical to MR's simplicity.