# Big table

```
A Bigtable is a sparse, distributed, persistent multidimensional sorted map.

The map is indexed by a row key, column key, and a timestamp; each value in the map is an uninterpreted array of bytes.
```
## Background

### Motivation

<img src="resources/pictures/bigtable_motivation.png" alt="bigtable_motivation" width="400"/>
<br/>


### Why not commercial DB
- Google's data is too large to scale
- Cost is too high
- Low level storage opertimization helps performance significantly

### What
<img src="resources/pictures/bigtable_what.png" alt="bigtable_what" width="400"/>
<br/>


### System Structure

<img src="resources/pictures/bigtable_system_structure.png" alt="bigtable_system_structure" width="400"/>
<br/>


<img src="resources/pictures/bigtable_building_blocks.png" alt="bigtable_building_blocks" width="400"/>
<br/>



### Data model
<img src="resources/pictures/bigtable_data_model.png" alt="bigtable_data_model" width="400"/>
<br/>

## Key words

### [map](https://en.wikipedia.org/wiki/Associative_array)
At its core, BigTable is a map.  A map is "an abstract data type composed of a collection of keys and a collection of values, where each key is associated with one value."  Example: 

```
{
  "zzzzz" : "woot",
  "xyz" : "hello",
  "aaaab" : "world",
  "1" : "x",
  "aaaaa" : "y"
}
```

### Persistent
Persistence merely means that the data you put in this special map "persists" after the program that created or accessed it is finished.  You could go to [GFS](./gfs.md) to understand how state is persistent.

### Distributed
BigTable is built upon distributed file systems so that the underlying file storage can be spread out among an array of independent machines.  Data is replicated across a number of participating nodes.  Refer to [GFS](./gfs.md) for more details.

### Sorted
key/value pairs are kept in strict alphabetical order. 
For example, row for the key "aaaab" should be right next to the row with key "aaaaa" and very far from the row with key "zzzzz".

```
{
  "1" : "x",
  "aaaaa" : "y",
  "aaaab" : "world",
  "xyz" : "hello",
  "zzzzz" : "woot"
}
```
- Sorting means better locality: During your table scan, the items of greatest interest to you are near each other.
- It is important when choosing a row key convention.
   + `com.jimbojw.www` better than `www.jimbojw.com`
   + `mail.jimbojw.com` should be near to `www.jimbojw.com` rather than `mail.xyz.com`
- Value is not sorted


### Multidimensional


### Sparse

 A given row can have any number of columns in each column family, or none at all. The other type of sparseness is row-based gaps which merely means that there may be gaps between keys.


## More info
- [Bigtable: A Distributed Storage System for Structured Data](https://static.googleusercontent.com/media/research.google.com/en//archive/bigtable-osdi06.pdf)
- [BigTable: A Distributed Structured Storage System By Jeff Dean 2014](https://www.youtube.com/watch?v=2cXBNQClehA) <span>&#9733;</span><span>&#9733;</span><span>&#9733;</span><span>&#9733;</span>
- [Designs, Lessons and Advice from Building Large Distributed Systems By Jeff Dean](https://www.cs.cornell.edu/projects/ladis2009/talks/dean-keynote-ladis2009.pdf) 
- [Building Software Systems At Google and Lessons Learned By Jeff Dean 2011](https://www.youtube.com/watch?v=modXC5IWTJI)
- [Introduction to HBase Schema Design](http://0b4af6cdc2f0c5998459-c0245c5c937c5dedcca3f1764ecc9b2f.r43.cf2.rackcdn.com/9353-login1210_khurana.pdf)
- [Understanding HBase and BigTable](https://dzone.com/articles/understanding-hbase-and-bigtab) [CN](https://www.cnblogs.com/ajianbeyourself/p/7789952.html)
- [Introduction to HBase Schema Design By Amandeep Khurana](http://0b4af6cdc2f0c5998459-c0245c5c937c5dedcca3f1764ecc9b2f.r43.cf2.rackcdn.com/9353-login1210_khurana.pdf)
- [Google Cloud Bigtable](https://cloud.google.com/bigtable/) 
   + [Bigtable in action (Google Cloud Next '17)](https://www.youtube.com/watch?v=KaRbKdMInuc&list=PLZuzb42lP7X7LlVIPJ0iaJgd4thV3qxfL&index=13)
- [github - spotify/simple-bigtable](https://github.com/spotify/simple-bigtable)
- [谷歌技术"三宝"之BigTable](https://blog.csdn.net/OpenNaive/article/details/7532589)
- [Rutgers CS 417 Distribute Systems - BigTable](https://www.cs.rutgers.edu/~pxk/417/notes/content/bigtable.html)
- [Bittiger Bigtable](https://posts.careerengine.us/p/5c1f17f86f311d533b55a93e)
