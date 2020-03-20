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
[sample code](https://github.com/google/s2geometry/blob/bec06921d72068fb22ef2100830c718659a19b58/src/s2/s2region_coverer.h#L40)

[implementation](https://github.com/google/s2geometry/blob/bec06921d72068fb22ef2100830c718659a19b58/src/s2/s2region_coverer.cc#L240)


```C++
  S2RegionCoverer::Options options;
  options.set_max_cells(5);
  S2RegionCoverer coverer(options);
  S2Cap cap(center, radius);
  S2CellUnion covering = coverer.GetCovering(cap);
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
You could find more samples in this file [s2closest_point_query_test.cc](https://github.com/google/s2geometry/blob/bec06921d72068fb22ef2100830c718659a19b58/src/s2/s2closest_point_query_test.cc#L42)


## Generate search terms

[code]()

```C++
  std::vector<S2Point> documents;
  documents.reserve(FLAGS_num_documents);
  for (int docid = 0; docid < FLAGS_num_documents; ++docid) {
    documents.push_back(S2Testing::RandomPoint());
  }

  // We use a hash map as our inverted index.  The key is an index term, and
  // the value is the set of "document ids" where this index term is present.
  std::unordered_map<string, std::vector<int>> index;

  // Create an indexer suitable for an index that contains points only.
  // (You may also want to adjust min_level() or max_level() if you plan
  // on querying very large or very small regions.)
  S2RegionTermIndexer::Options options;
  options.set_index_contains_points_only(true);
  S2RegionTermIndexer indexer(options);

  // Add the documents to the index.
  for (int docid = 0; docid < documents.size(); ++docid) {
    S2Point index_region = documents[docid];
    for (const auto& term : indexer.GetIndexTerms(index_region, kPrefix)) {
      // [perry]
      // - term string generated for multiple layers
      // - for near by location, they will share similar prefix
      // term is s2:9b1
      // term is s2:9b0c
      // term is s2:9b09
      // term is s2:9b094
      // term is s2:9b091
      // term is s2:9b091c
      // term is s2:9b091b
      // term is s2:9b091b4
      // term is s2:9b091b7
      // term is s2:9b091b6c
      // term is s2:9b091b6f
      // term is s2:9b091b6fc
      // term is s2:9b091b6fd
      index[term].push_back(docid);
    }
  }
```

