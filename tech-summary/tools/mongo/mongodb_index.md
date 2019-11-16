# MongoDB index

## 怎么建索引更能提高查询性能？
在查询时，索引是否高效，要注意它的cardinality（cardinality越高表示该键可选择的值越多），在组合索引中，让cardinality高的放在前面。注意这里跟分布式环境选择shard key的不同。

```
index cardinality（索引散列程度），表示的是一个索引所对应到的值的多少，散列程度越低，则一个索引对应的值越多，索引效
果越差：在使用索引时，高散列程度的索引可以更多的排除不符合条件的文档，让后续的比较在一个更小的集合中执行，这更高效。所
以一般选择高散列程度的键做索引，或者在组合索引中，把高散列程度的键放在前面。

《MongoDB——The Definitive Guide 2nd Edition》
```

## More info
- [Mongodb index](https://docs.mongodb.com/manual/indexes/)

- [Optimizing MongoDB Compound Indexes](https://emptysqua.re/blog/optimizing-mongodb-compound-indexes/?spm=a2c4e.10696291.0.0.734519a4U8dAwp)
