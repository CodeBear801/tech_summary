# ALT

ALT means Astar, landmarks and triangle-inequality.  
Landmarks could help to provide better hubristic for Astar.  

<img src="../resources/astar_extreme_case.png" alt="astar_extreme_case" width="400"/>
<br/>
For goal directed hubristic(Euclidean distance), the upper case always be a trouble for routing.

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

- What are good landmarks
A good landmark appears “before” v or “after” w.

<img src="../resources/alt_landmark_selection_1.png" alt="alt_landmark_selection_1" width="400"/>
<br/>

<img src="../resources/alt_landmark_selection_2.png" alt="alt_landmark_selection_2" width="200"/>
<br/>

<img src="../resources/alt_when_is_good_landmark.png" alt="alt_when_is_good_landmark" width="400"/>
<br/>



<br/>

- How to select landmarks

<img src="../resources/alt_greedy_selection_landmark.png" alt="alt_greedy_selection_landmark" width="400"/>
<br/>

<br/>

- How to compute distance from all vertex to specific landmark  
Could get result in one dijkstra on reverse map
<br/>


- How to compute distance from one vertex to a set of landmarks

<img src="../resources/alt_dijkstra_sets_node.png" alt="alt_dijkstra_sets_node" width="400"/>
<br/>



## More info
- [Landmarks - Algorithms on Graphs - University of California San Diego](https://www.coursera.org/lecture/algorithms-on-graphs/landmarks-optional-h3uOb)
- [Landmark selection - On Preprocessing the ALT-Algorithm - Fabian Fuchs from KIT](http://www.fabianfuchs.com/fabianfuchs_ALT.pdf)
