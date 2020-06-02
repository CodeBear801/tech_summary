
## multiple dimension array

### C++
- Create and init a two dimensional vector [ref](https://stackoverflow.com/questions/17663186/initializing-a-two-dimensional-stdvector)
```C++
std::vector<std::vector<int> > fog(
    NUM_4_ROWS,
    std::vector<int>(NUM_4_COLUMNS /*,  INT_MAX*/)); // Defaults to zero initial value

// or
std::vector<std::vector<int> > fog { { 1, 1, 1 },
                                    { 2, 2, 2 } };
```

### Golang
[ref](https://stackoverflow.com/questions/39804861/what-is-a-concise-way-to-create-a-2d-slice-in-go)

```go
// make creates new slice
a := make([][]uint8, dy)
for i := range a {
    a[i] = make([]uint8, dx)
}

// or
a := [][]uint8{
    {0, 1, 2, 3},
    {4, 5, 6, 7},
}
```
