
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


* Map worker crashes:
master sees worker no longer responds to pings
crashed worker's intermediate Map output is lost but is likely needed by every Reduce task!
master re-runs, spreads tasks over other GFS replicas of input.
some Reduce workers may already have read failed worker's intermediate data.
  here we depend on functional and deterministic Map()!
master need not re-run Map if Reduces have fetched all intermediate data
  though then a Reduce crash would then force re-execution of failed Map

* Reduce worker crashes.
finshed tasks are OK -- stored in GFS, with replicas.
master re-starts worker's unfinished tasks on other workers.

* Reduce worker crashes in the middle of writing its output.
GFS has atomic rename that prevents output from being visible until complete.
so it's safe for the master to re-run the Reduce tasks somewhere else.

* What if the master gives two workers the same Map() task?
perhaps the master incorrectly thinks one worker died.  it will tell Reduce workers about only one of them.
* What if the master gives two workers the same Reduce() task?
they will both try to **write the same output file** on GFS!  atomic GFS rename prevents mixing; one complete file will be visible.
* What if a single worker is very slow -- a "straggler"?
perhaps due to flakey hardware.  master starts a second copy of last few tasks.
* What if a worker computes incorrect output, due to broken h/w or s/w?
No way! MR assumes "fail-stop" CPUs and software.
* What if the master crashes?
recover from check-point, or give up on job
