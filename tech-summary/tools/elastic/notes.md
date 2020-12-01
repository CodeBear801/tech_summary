

Elasticsearch Notes (11302020):

1. About synonym set

ElasticSearch需要定义analyzer, 也就是如何分词 https://www.elastic.co/guide/en/elasticsearch/reference/6.4/analyzer.html  
以及如何去建立同义词索引  参见https://www.elastic.co/guide/en/elasticsearch/reference/6.4/analysis-synonym-tokenfilter.html

- Elasticsearch 正确设置同义词近义词 https://www.jianshu.com/p/893d35d53356


2. Elasticsearch Relevance tune

Relevance Tuning Guide, Weights and Boosts
https://www.elastic.co/guide/en/app-search/current/relevance-tuning-guide.html#relevance-tuning-guide

3. Elasticsearch update index by realtime???

reindex?

How does one keep an Elasticsearch index up-to-date using elasticsearch-dsl-py?
https://stackoverflow.com/questions/48103731/how-does-one-keep-an-elasticsearch-index-up-to-date-using-elasticsearch-dsl-py


**acceptable replication delay and data consistency**

Keeping Elasticsearch in Sync
https://www.elastic.co/blog/found-keeping-elasticsearch-in-sync

How we reindexed 36 billion documents in 5 days within the same Elasticsearch cluster
https://thoughts.t37.net/how-we-reindexed-36-billions-documents-in-5-days-within-the-same-elasticsearch-cluster-cd9c054d1db8

应用实战工作坊
https://github.com/geektime-geekbang/geektime-ELK/blob/master/part-5/%E5%BA%94%E7%94%A8%E5%AE%9E%E6%88%98%E5%B7%A5%E4%BD%9C%E5%9D%8A.pdf


Elasticsearch Notes (11302020):  



GET /_cat/indices   get information of all index
https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-indices.html


让Elasticsearch飞起来！百亿级实时查询优化实战 +++
https://dbaplus.cn/news-73-2550-1.html

## Routing
类似于分布式数据库中的分片原则，将符合规则的数据存储到同一分片。ES 通过哈希算法来决定数据存储于哪个 Shard：
```
shard_num = hash(_routing) % num_primary_shards
```
不指定 Routing 的查询过程会查询每个 Shard，然后做结果聚合，总的时间大概就是所有 Shard 查询所消耗的时间之和。
指定 Routing 以后会根据 Routing 查询特定的一个或多个 Shard


```
_search?routing=beijing：按城市。
_search?routing=beijing_user123：按城市和用户。
_search?routing=beijing_android，shanghai_android：按城市和手机类型等。
```

如果某个routing下面的数据越来越多如何做?
一种解决办法是单独为这些数据量大的渠道创建独立的 Index
```
http://localhost:9200/shanghai,beijing,other/_search?routing=android
```
另一种办法是指定 Index 参数 index.routing_partition_size，来解决最终可能产生群集不均衡的问题，
```
shard_num = (hash(_routing) + hash(_id) % routing_partition_size) % num_primary_shards	
```


## 索引拆分

### Index Templates
Why - 在开发中，elasticsearch很大一部分工作是用来处理日志信息的，比如某公司对于日志处理策略是以日期为名创建每天的日志索引。并且每天的索引映射类型和配置信息都是一样的，只是索引名称改变了。如果手动的创建每天的索引，将会是一件很麻烦的事情。为了解决类似问题，elasticsearch提供了预先定义的模板进行索引创建，这个模板称作为Index Templates。通过索引模板可以让类似的索引重用同一个模板。 [ref elasticsearch之Index Templates](https://www.cnblogs.com/Neeo/articles/10869231.html)

```
PUT _template/2019
{
  "index_patterns": ["20*", "product1*"],   ①
  "settings":{   ②
    "number_of_shards": 2,
    "number_of_replicas": 1
  },
  "mappings":{  ③
    "doc":{
      "properties":{
        "ip":{
          "type":"keyword"
        },
        "method":{
          "type": "keyword"
        }
      }
    }
  }
}
```
一个 Index 可以被多个 Template 匹配，那 Settings 和 Mappings 就是多个 Template 合并后的结果。

### _rollover API 
why: 用户通过创建基于标准时间段的索引来管理其集群的数据生命周期，通常每天一个索引。如果您的工作负载有多个数据流，并且每个流有不同的数据大小，那么您可能会遇到问题：您的资源使用量（尤其是每个节点的存储使用量）可能会变得失衡，或者说发生“偏斜”。 发生这种情况时，一些节点将会过载或比其他节点更早耗尽存储，并且您的集群可能会停止运行。 可以使用_rollover API 来管理索引的大小: 使用一个定义了 Elasticsearch 何时应创建新索引并开始写入的阈值，按照固定的周期调用 _rollover。

<img src="https://d2908q01vomqb2.cloudfront.net/ca3512f4dfa95a03169c5a670a4c91a19b3077b4/2019/08/06/Rollover3.jpg" alt="replicas" width="400"/>
<br/>



<img src="https://d2908q01vomqb2.cloudfront.net/ca3512f4dfa95a03169c5a670a4c91a19b3077b4/2019/08/06/Rollover2.jpg" alt="replicas" width="400"/>
<br/>


https://d2908q01vomqb2.cloudfront.net/ca3512f4dfa95a03169c5a670a4c91a19b3077b4/2019/08/06/Rollover2.jpg

```json
POST _aliases
{
  "actions": [
    {
      "add": {
        "index": "weblogs-000001",
        "alias": "weblogs",
        "is_write_index": true
      }
    }
  ]
}

POST /weblogs/_rollover 
{
  "conditions": {
    "max_size":  "10gb"  // max_age、max_docs, max_size
  }
}
```


## Hot-Warm

to do 












## Terms

### shards& replicas

<img src="https://user-images.githubusercontent.com/16873751/100789276-1514e280-33cb-11eb-80ca-350ec0d9cba8.png" alt="replicas" width="400"/>
<br/>


<img src="https://user-images.githubusercontent.com/16873751/100789618-8f456700-33cb-11eb-961b-1aca86f16164.png" alt="replicas" width="400"/>
<br/>

mget/msearch/bulk


### Analyser 分词

ElasticSearch需要定义analyzer, 也就是如何分词 https://www.elastic.co/guide/en/elasticsearch/reference/6.4/analyzer.html

以及如何去建立同义词索引 参见https://www.elastic.co/guide/en/elasticsearch/reference/6.4/analysis-synonym-tokenfilter.html


<img src="https://user-images.githubusercontent.com/16873751/100792815-14327f80-33d0-11eb-8dc4-88ede8453e62.png" alt="replicas" width="400"/>
<br/>

<img src="https://user-images.githubusercontent.com/16873751/100793948-c28af480-33d1-11eb-9f70-446d177c124c.png" alt="replicas" width="400"/>
<br/>

### Query string

<img src="https://user-images.githubusercontent.com/16873751/100793984-cae32f80-33d1-11eb-9133-3b0b4635d731.png" alt="replicas" width="400"/>
<br/>

<img src="https://user-images.githubusercontent.com/16873751/100794006-d33b6a80-33d1-11eb-9dba-9500de350fcb.png" alt="replicas" width="400"/>
<br/>

<img src="https://user-images.githubusercontent.com/16873751/100794020-d8001e80-33d1-11eb-8f37-4b877dc78673.png" alt="replicas" width="400"/>
<br/>

<img src="https://user-images.githubusercontent.com/16873751/100794043-ddf5ff80-33d1-11eb-9f99-266b176e4974.png" alt="replicas" width="400"/>
<br/>

<img src="https://user-images.githubusercontent.com/16873751/100794077-efd7a280-33d1-11eb-97ad-a0f54251b5e0.png" alt="replicas" width="400"/>
<br/>


## aggregation

数据统计分析，


<img src="https://user-images.githubusercontent.com/16873751/100805546-b314a700-33e3-11eb-85dd-d2ba12cf12bb.png" alt="replicas" width="400"/>
<br/>
bucket -> group, metric -> count
