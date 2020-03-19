- How `Covering` find the most optimal solution
- Software layering: S2Point, S1Angle, S2LatLng -> S2 Region -> S2 Builder -> S2ShapeIndex -> S2ClosestEdgeQuery
- Robustness is achieved through the use of the techniques: 
    + conservative error bounds (which measure the error in certain calculations) 
    + exact geometric predicates (which determine the true mathematical relationship among geometric objects)
    + snap rounding (a technique that allows rounding results to finite precision without creating topological errors).