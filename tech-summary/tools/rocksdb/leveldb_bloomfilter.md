# bloom filter

Basic idea of bloom filter  
For each key, use multiple hash function to translate which into `bit map`.  

<img src="https://user-images.githubusercontent.com/16873751/96755613-cb76b800-1387-11eb-8d00-73e45103ffb1.png" alt="bloomfilter" width="600"/>


The character of bloom filter  
- If bloom filter detect key is not in the set, then it must not be: never false negative
- If bloom filter detect the key inside the set, then there is high possibility its in, but also possible not exist, which is false positive

Bloom filter's performance is related with manage "false positive".  If the big array is too small, then probably each bit is 1, then there is high percentage of false positive.

<img src="https://user-images.githubusercontent.com/16873751/96754331-ffe97480-1385-11eb-80dc-36f066638556.png" alt="bloomfilter" width="600"/>

More info:
- [Bloom Filters - the math](http://pages.cs.wisc.edu/~cao/papers/summary-cache/node8.html)
- [Less Hashing, Same Performance: Building a Better Bloom Filter](https://www.eecs.harvard.edu/~michaelm/postscripts/rsa2008.pdf) 

