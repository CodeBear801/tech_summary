# Basic Routing Algorithms

Summary about basic routing algorithms, based on questions [here](https://github.com/Telenav/open-source-spec/blob/master/routing_basic/doc/routing_alg_questions.md)

## Dijkstra

- Grow a ball around s and stop when t is settled
- PQ + Hash map
- Algorithm detail

<img src="../resources/dijkstra_alg_1.png" alt="dijkstra_alg_1" width="400"/>
<br/>


- Complexity

  O(VNlogV) -> O(ElogV)  
  https://stackoverflow.com/questions/26547816/understanding-time-complexity-calculation-for-dijkstra-algorithm  

- Proof

- Difference between visited / settled

- pseudo code
```C++
do{
    Pop()
    UpdateData()
    Relax()
} while(Stop())
```


## Bidirectional Dijkstra
- Grow a ball around end(s and t) until they meet

- Unlike uni-direction dijkstra, first meet could not guarantee best solution:


<img src="../resources/dijkstra_bidir_stop.png" alt="dijkstra_bidir_stop" width="400"/>
<br/>

- Proof

<img src="../resources/dijkstra_bidir_proof.png" alt="dijkstra_bidir_proof" width="400"/>
<br/>

## A*

- Goal directed, add heuristic, make the ball become ellipse

- Algorithm detail

<img src="../resources/astar_alg_1.png" alt="astar_alg_1" width="400"/>
<br/>

- Why is A* equivalent to Dijkstra on the modified graph?

<img src="../resources/astar_why_1.png" alt="astar_why_1" width="400"/>
<br/>


- Why is π(v) a lower bound on dist(v, t) when π is feasible and π(t)=0?

<img src="../resources/astar_why_2.png" alt="astar_why_2" width="400"/>
<br/>


- How to make sure l' is > 0
Triangle inequality ensures 

<img src="../resources/astar_fomular_1.png" alt="astar_formular_1" width="400"/>
<br/>

- What's the meaning of Dijkstra's algorithm only explore the shortest path?
In extreme case, only edges on shortest path would pop-out.  
This could prove in ideal situation, why A* works.  
Let's assume in all vertex you record shortest path to t, so node on shortest path be pops because:  

<img src="../resources/astar_fomular_2.png" alt="astar_formular_2" width="400"/>
<br/>