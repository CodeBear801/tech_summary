# S2 Samples

## Get max cover

[code](https://github.com/google/s2geometry/blob/9398b7c8d55c15c4ad7cdc645c482232ea7c087a/src/s2/s2shape_index_region.h#L49
)
```C++
S2CellUnion GetCovering(const S2ShapeIndex& index) {
  S2RegionCoverer coverer;
  coverer.mutable_options()->set_max_cells(20);
  S2CellUnion covering;
  coverer.GetCovering(MakeS2ShapeIndexRegion(&index), &covering);
  return covering;
}
```

## Test containment
[code](https://github.com/google/s2geometry/blob/9398b7c8d55c15c4ad7cdc645c482232ea7c087a/src/s2/mutable_s2shape_index.h#L79)

```C++
  void TestContainment(const vector<S2Point>& points,
                       const vector<S2Polygon*>& polygons) {
    MutableS2ShapeIndex index;
    for (auto polygon : polygons) {
      index.Add(absl::make_unique<S2Polygon::Shape>(polygon));
    }
    auto query = MakeS2ContainsPointQuery(&index);
    for (const auto& point : points) {
      for (S2Shape* shape : query.GetContainingShapes(point)) {
        S2Polygon* polygon = polygons[shape->id()];
        ... do something with (point, polygon) ...
      }
    }
  }
```

## S2ClosestPointQuery

[code](https://github.com/google/s2geometry/blob/9398b7c8d55c15c4ad7cdc645c482232ea7c087a/src/s2/s2closest_point_query.h#L150)

```C++
  // Build an index containing random points anywhere on the Earth.
  S2PointIndex<int> index;
  for (int i = 0; i < FLAGS_num_index_points; ++i) {
    index.Add(S2Testing::RandomPoint(), i);
  }

  // Create a query to search within the given radius of a target point.
  S2ClosestPointQuery<int> query(&index);
  query.mutable_options()->set_max_distance(
      S1Angle::Radians(S2Earth::KmToRadians(FLAGS_query_radius_km)));

  // Repeatedly choose a random target point, and count how many index points
  // are within the given radius of that point.
  int64 num_found = 0;
  for (int i = 0; i < FLAGS_num_queries; ++i) {
    S2ClosestPointQuery<int>::PointTarget target(S2Testing::RandomPoint());
    num_found += query.FindClosestPoints(&target).size();
  }

  printf("Found %lld points in %d queries\n", num_found, FLAGS_num_queries);

```