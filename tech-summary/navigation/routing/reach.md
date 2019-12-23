# Reach

Reach is a **pruning strategy**.  Reach value could help to filter un-reasonable nodes during exploring.

## What is Reach value

<img src="../resources/reach_basic.png" alt="reach_basic" width="400"/>
<br/>

## Why Reach

Vertices on highways have high reach  
Vertices on local roads have low reach  

## How Reach works

Reach is used to prune the search during an s-t query.  During scanning an edge (v, w)  
If `reach(w) < min{d(s, v) + l(v, w), LB(w, t)}`, then w can be pruned  



