# Uber H3 Experiment

## index.c

```shell
cc -lh3 ../examples/index.c -o example_index
```

```c
location.lat = degsToRads(40.689167);
location.lon = degsToRads(-74.044444);
// print index for all resolutions
for (int i = 0; i < 16; ++i)
{
H3Index indexed = geoToH3(&location, i);
printf("The index for resolution %d is: %" PRIx64 "\n", i, indexed);
}
```


```
The index for resolution 0 is:  802bfffffffffff
The index for resolution 1 is:  812a3ffffffffff
The index for resolution 2 is:  822a17fffffffff
The index for resolution 3 is:  832a10fffffffff
The index for resolution 4 is:  842a107ffffffff
The index for resolution 5 is:  852a1073fffffff
The index for resolution 6 is:  862a1072fffffff
The index for resolution 7 is:  872a1072bffffff
The index for resolution 8 is:  882a1072b5fffff
The index for resolution 9 is:  892a1072b5bffff
The index for resolution 10 is: 8a2a1072b59ffff
The index for resolution 11 is: 8b2a1072b59bfff
The index for resolution 12 is: 8c2a1072b59b9ff
The index for resolution 13 is: 8d2a1072b59b97f
The index for resolution 14 is: 8e2a1072b59b967
The index for resolution 15 is: 8f2a1072b59b960

```

## neighbors.c

```
k-ring 0 is defined as the origin index, k-ring 1 is defined as k-ring 0 and all neighboring indices, and so on.
```

K= 1
```
Neighbors:
8a2a1072b59ffff
8a2a1072b597fff
8a2a1070c96ffff
8a2a1072b4b7fff
8a2a1072b4a7fff
8a2a1072b58ffff
8a2a1072b587fff
```

K =2
```
Neighbors:
8a2a1072b59ffff
8a2a1072b597fff
8a2a1070c96ffff
8a2a1072b4b7fff
8a2a1072b4a7fff
8a2a1072b58ffff
8a2a1072b587fff
8a2a1072b5b7fff
8a2a1072a2cffff
8a2a1070c967fff
8a2a1070c947fff
8a2a1070c94ffff
8a2a1072b497fff
8a2a1072b487fff
8a2a1072b4affff
8a2a1072b417fff
8a2a1072b437fff
8a2a1072b5affff
8a2a1072b5a7fff

```

## compact

```c
H3Index input[] = {
// All with the same parent index
0x8a2a1072b587fffL, 0x8a2a1072b5b7fffL, 0x8a2a1072b597fffL,
0x8a2a1072b59ffffL, 0x8a2a1072b58ffffL, 0x8a2a1072b5affffL,
0x8a2a1072b5a7fffL,
// These don't have the same parent index as above.
0x8a2a1070c96ffffL, 0x8a2a1072b4b7fffL, 0x8a2a1072b4a7fffL};

```


```
Starting with 10 indexes.
Compacted:
8a2a1070c96ffff
8a2a1072b4b7fff
8a2a1072b4a7fff
892a1072b5bffff
Compacted to 4 indexes.
Uncompacted:
8a2a1070c96ffff
8a2a1072b4b7fff
8a2a1072b4a7fff
8a2a1072b587fff
8a2a1072b58ffff
8a2a1072b597fff
8a2a1072b59ffff
8a2a1072b5a7fff
8a2a1072b5affff
8a2a1072b5b7fff
Uncompacted to 10 indexes.

```

