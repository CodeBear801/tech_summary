

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
