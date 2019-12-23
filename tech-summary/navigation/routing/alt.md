# ALT

ALT means Astar, landmarks and triangle-inequality.  
Landmarks could help to provide better hubristic for Astar.

<img src="../resources/astar_extreme_case.png" alt="astar_extreme_case" width="400"/>
<br/>

## Algorithm
Pre-processing: select a small number of vertices Landmarks  (L)
For all nodes v store distance vector d(v, l) to all landmarks(l belongs to L)

- What is Triangle-inequality

<img src="../resources/triangle_inequality_1.png" alt="triangle_inequality_1" width="400"/>
<br/>


- How to compute distance from all vertex to specific landmark
Could get result in one dijkstra on reverse map
