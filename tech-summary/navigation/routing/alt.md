# ALT

ALT means Astar, landmarks and triangle-inequality.  
Landmarks could help to provide better hubristic for Astar.

<img src="../resources/astar_extreme_case.png" alt="astar_extreme_case" width="400"/>
<br/>

## Algorithm
Pre-processing:  
Select a small number of vertices Landmarks (L)  
For all nodes v store distance vector d(v, l) to all landmarks(l belongs to L)  

Query:
h(u) = |dist(l, u) - dist(l, t)| <= dist (u,t)



- What is Triangle-inequality

<img src="../resources/triangle_inequality_1.png" alt="triangle_inequality_1" width="400"/>
<br/>

- Multiple landmarks

<img src="../resources/alt_multiple_landmarks.png" alt="alt_multiple_landmarks" width="400"/>
<br/>

- How to select landmarks
A good landmark appears “before” v or “after” w.

<img src="../resources/alt_landmark_selection_1.png" alt="alt_landmark_selection_1" width="400"/>
<br/>

<img src="../resources/alt_landmark_selection_2.png" alt="alt_landmark_selection_2" width="200"/>
<br/>



- How to compute distance from all vertex to specific landmark
Could get result in one dijkstra on reverse map


- How to compute distance from one vertex to a set of landmarks

<img src="../resources/alt_dijkstra_sets_node.png" alt="alt_dijkstra_sets_node" width="400"/>
<br/>




