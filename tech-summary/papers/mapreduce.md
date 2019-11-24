
How MapReduce scales 
- N computers gets you Nx throughput.
    - Assuming M and R are >= N (i.e., lots of input files and map output keys).
    - Maps()s can run in parallel, since they don't interact.  Same for Reduce()s.
    - The only interaction is via the "shuffle" in between maps and reduces.
- So you can get more throughput by buying more computers.
    - Rather than special-purpose efficient parallelizations of each application.


