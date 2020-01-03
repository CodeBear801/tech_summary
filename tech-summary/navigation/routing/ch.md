# Contraction Hierarchies

Pre select "important" nodes, during query time route to far away only use specific nodes.

## Hierarchical approach

### Basic intuition

Route selection on Google maps  
close to source and target: mainly white(local) roads  
a bit further away: mainly yellow(national/express) roads  
even further away: mainly orange(motorway/highway) roads

### Function class


<img src="../resources/ch_hierarchical_approach_fc.png" alt="ch_hierarchical_approach_fc.png" width="400"/>
<br/>

### Highway Hierarchies


<img src="../resources/ch_hierarchical_approach_hh.png" alt="ch_hierarchical_approach_hh.png" width="400"/>
<br/>

### Contraction Hierarchies


<img src="../resources/ch_hierarchical_approach_ch.png" alt="ch_hierarchical_approach_ch.png" width="400"/>
<br/>

## How to contract

<img src="../resources/ch_basic_idea.png" alt="ch_basic_idea.png" width="600"/>
<br/>

## Node Ordering

Node ordering is very important for CH's performance.  

Bad case:  
<img src="../resources/ch_bad_case.png" alt="ch_bad_case.png" width="400"/>
<br/>

Maintain the nodes in a priority queue, in the order of how attractive it is to contract the respective node next  
The less shortcuts we have to add, the better  

## ED

<img src="../resources/ch_ed.png" alt="ch_ed.png" width="400"/>
<br/>

<img src="../resources/ch_ed2.png" alt="ch_ed2.png" width="400"/>
<br/>

### Spatial diversity heuristic
For each node maintain a count of the number of neig y hbours that have already been contracted, and add this to the ED.
**The more neighbours have already been contracted, the later this node will be contracted**

<img src="../resources/ch_spatial_diversity.png" alt="ch_spatial_diversity.png" width="400"/>
<br/>


### EDS5


<img src="../resources/ch_eds5.png" alt="ch_eds5.png" width="600"/>
<br/>

### N5

<img src="../resources/ch_n5.png" alt="ch_n5.png" width="600"/>
<br/>


## More info
- [Contraction Hierarchies: Faster and Simpler Hierarchical Routing in Road Networks](http://algo2.iti.kit.edu/schultes/hwy/contract.pdf)
- [Contraction Hierarchies briefly explained](../resources/Contraction&#32;Hierarchies&#32;briefly&#32;explained.pdf)
- [Adaptive Routing on contracted graphs](https://www.youtube.com/watch?v=O74F-hpaKWM)
- [Contraction Hierarchies animation](https://www.mjt.me.uk/posts/contraction-hierarchies/)

