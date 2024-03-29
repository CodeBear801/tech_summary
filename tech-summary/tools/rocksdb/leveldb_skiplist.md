# skiplist

Basic idea: trade space for time  
To insert a <k, v> pair into DB, the quickest way for insert itself is adding a node in `list`.  
But if we want also to provide a quick search, we'd better to make the list in sorted order and then we could use binary search, which is not easy in traditional list due to we need `O(n)` to travel nodes one by one.  What if we build additional layer of index on top of normal list to provide `second level index`

<img src="https://user-images.githubusercontent.com/16873751/96521686-a9b1ef80-1226-11eb-8a5e-bd0f1df3a0da.png" alt="skiplist" width="600"/>   

(image from: https://zhuanlan.zhihu.com/p/54869087)  

<br/>
<img src="https://user-images.githubusercontent.com/16873751/96521884-23e27400-1227-11eb-9359-4bb0c1472fc9.png" alt="skiplist" width="600"/>

- [code of skiplist.h](https://github.com/google/leveldb/blob/b7d302326961fb809d92a95ce813e2d26fe2e16e/db/skiplist.h#L42)
- [code of skiplist testing](https://github.com/google/leveldb/blob/b7d302326961fb809d92a95ce813e2d26fe2e16e/db/skiplist_test.cc#L151)

### CAS
<img src="https://user-images.githubusercontent.com/16873751/107069091-4bd5f300-6796-11eb-8087-fd1fbad82a28.png" alt="skiplist" width="600"/>
<br/>
from: http://ifeve.com/cas-skiplist/


## More info
- [跳表SkipList](https://www.cnblogs.com/xuqiang/archive/2011/05/22/2053516.html)
