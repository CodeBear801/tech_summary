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

## A Star

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


## Bidirectional A Star

- Make forward search and backward search consistent

<img src="../resources/astar_alg_2.png" alt="astar_alg_2" width="600"/>
<br/>

- Stop condition

<img src="../resources/astar_alg_stop.png" alt="astar_alg_stop" width="600"/>
<br/>


## Bellman Ford

 
<img src="../resources/bellman_ford_1.png" alt="bellman_ford_1" width="600"/>
<br/>

如果不是从起点开始 开始的loop是不是就浪费了?

Bellman Ford有这几个假设
1. 从起点到任何一个点的最短距离，顶多会经过n-1个端点(无环)  
2.相当于从起点一层一层往外推，每次更新一层的cost  

<<挑战程序竞赛>>

<img src="../resources/book_acm_1.png" alt="book_acm_1" width="600"/>
<br/>

<img src="../resources/book_acm_2.png" alt="book_acm_2" width="600"/>
<br/>

<img src="../resources/book_acm_3.png" alt="book_acm_3" width="600"/>
<br/>

<<算法导论>>

<img src="../resources/book_ita_3.png" alt="book_ita_3" width="400"/>
<br/>


## Fibonacci heap 

-  Why Fibonacci heap 
https://stackoverflow.com/questions/14252582/how-can-i-use-binary-heap-in-the-dijkstra-algorithm  
Heap data-structures have no way of getting at any particular node that is not the minimum or the last node!  
For specific node in heap, we might find a better path to him with lower cost
