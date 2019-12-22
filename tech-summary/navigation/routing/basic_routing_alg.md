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

