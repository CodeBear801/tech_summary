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

