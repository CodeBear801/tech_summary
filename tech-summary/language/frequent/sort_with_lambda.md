
## sort array

### C++
[ref](https://stackoverflow.com/questions/5122804/how-to-sort-with-a-lambda)

```C++
sort(mMyClassVector.begin(), mMyClassVector.end(), 
    [](const MyClass & a, const MyClass & b) -> bool
{ 
    return a.mProperty > b.mProperty; 
});

```

### Golang

- method 1

```go
s := []int{4, 2, 3, 1}
sort.Ints(s)  
fmt.Println(s) // [1 2 3 4]
```
`go` also provide [sort.Float64s](https://golang.org/pkg/sort/#Float64s), [sort.Strings](https://golang.org/pkg/sort/#Strings)

- method 2: custom comparator
```go
family := []struct {
    Name string
    Age  int
}{
    {"Alice", 23},
    {"David", 2},
    {"Eve", 2},
    {"Bob", 25},
}

// Sort by age, keeping original order or equal elements.
sort.SliceStable(family, func(i, j int) bool {
    return family[i].Age < family[j].Age
})
fmt.Println(family) // [{David 2} {Eve 2} {Alice 23} {Bob 25}]
```
`sort.Slice` or `sort.SliceStable` based on function `less(i, j int) bool`

[code example1](https://github.com/Telenav/osrm-backend/blob/0b461183b97de493983ba44749c772719849fd3e/integration/service/oasis/selectionstrategy/reachable_by_single_charge.go#L173)

- method 3: implement interface for supporting sort
[golang sort interface](https://golang.org/pkg/sort/#Interface)
```go
type Interface interface {
        // Len is the number of elements in the collection.
        Len() int
        // Less reports whether the element with
        // index i should sort before the element with index j.
        Less(i, j int) bool
        // Swap swaps the elements with indexes i and j.
        Swap(i, j int)
}
```

[code example1](https://github.com/Telenav/osrm-backend/blob/0b461183b97de493983ba44749c772719849fd3e/integration/service/ranking/strategy/rankbyduration/rank.go#L25)



## Sort map

### Golang

[ref](https://yourbasic.org/golang/how-to-sort-in-go/)

A `map` is an unordered collection of key-value pairs. If you need a stable iteration order, you must maintain a separate data structure.
```go
m := map[string]int{"Alice": 2, "Cecil": 1, "Bob": 3}

keys := make([]string, 0, len(m))
for k := range m {
    keys = append(keys, k)
}
sort.Strings(keys)

for _, k := range keys {
    fmt.Println(k, m[k])
}
// Output:
// Alice 2
// Bob 3
// Cecil 1
```

## Search in sorted array

## C++
[`lower_bound`](https://en.cppreference.com/w/cpp/algorithm/lower_bound) is an STL function, it use **binary search** over a sorted array a to find smallest pointer which meet condition of a[i] >=k.  
[`uppter_bound`](https://en.cppreference.com/w/cpp/algorithm/upper_bound) finds smallest pointer which meet condition of a[i]>k

```C++
// find how many k from sorted array a whose size is n
upper_bound(a, a+n, k) - lower_bound(a, a+n, k)
```


- In `C++`, `std::map` and `std::set` is sorted(rbtree)
```C++
std::set<pair<int, int>> bookingRecord; 
public: 
    bool book(int s, int e) { 
        auto next = bookingRecord.lower_bound({s, e}); // first element with key not go before k (i.e., either it is equivalent or goes after). 
    }
```


## Golang
[golang sort.Search's example](https://golang.org/pkg/sort/#example_Search)



## Insert element into sorted array

### Golang
If existing slice's capacity is not enough, a new slice will be created and all values of existing slice will be copied.  [ref](https://stackoverflow.com/questions/42746972/golang-insert-to-a-sorted-slice)
```go
func Insert(ss []string, s string) []string {
    i := sort.SearchStrings(ss, s)
    ss = append(ss, "")
    copy(ss[i+1:], ss[i:])
    ss[i] = s
    return ss
}

```
[code example1](https://github.com/Telenav/osrm-backend/blob/0b461183b97de493983ba44749c772719849fd3e/integration/service/oasis/stationconnquerier/station_conn_querier.go#L141)



