# Elasticsearch Commands

## Basic
Throughout `{endpoint}` refers to the ElasticSearch index type (aka table). Note that ElasticSearch often let’s you run the same queries on both “indexes” (aka database) and types.
```
// query
{endpoint}/_search?size=5&pretty=true
// schema
{endpoint}/_mapping
```

### Curl
```
curl 'localhost:9200/_cat/indices?v' 
```

### Python
```python
import urllib2
import json
# Store some data

url = '{endpoint}'
data = {
    'title': 'jones',
    'amount': 5.7
    }
# have to send the data as JSON
data = json.dumps(data)

req = urllib2.Request(url, data, headers)
out = urllib2.urlopen(req)
print out.read()
```

## System

```
// start
./elasticsearch -d
// find pid
ps -ef | grep elastic

```


## DBOP

- Create/Delete index
```
curl -XPUT 'localhost:9200/customer?pretty'

curl -XDELETE 'localhost:9200/customer'
```

- Insert/Get/Update/Delete data
```
   curl -XPUT 'localhost:9200/customer/external/1?pretty' -d '
　　{
       　　  "name": "John Doe"
　　}'

   curl -XGET 'localhost:9200/customer/external/1?pretty'

　　curl -XPUT 'localhost:9200/customer/external/1?pretty' -d '
　　{
 　　 "name": "Jane Doe"
　　}'

   curl -XDELETE 'localhost:9200/customer?pretty' 
```

- Batch
```
   // update data with id = 1 and then delete data with id = 2
　　curl -XPOST 'localhost:9200/customer/external/_bulk?pretty' -d '
　　{"update":{"_id":"1"}}
　　{"doc": { "name": "John Doe becomes Jane Doe" } }
　　{"delete":{"_id":"2"}}
```

- Import
```
    curl -XPOST 'localhost:9200/bank/account/_bulk?pretty' --data-binary "@accounts.json"

    curl 'localhost:9200/_cat/indices?v'
```


## Query
- Query for 10 result
```
 　　curl -XPOST 'localhost:9200/bank/_search?pretty' -d '
　　{
  　　"query": { "match_all": {} },
　　  "from": 10,
 　　 "size": 10
　　}'
```

- Query for specific match
```
　  curl -XPOST 'localhost:9200/bank/_search?pretty' -d '
　　{
 　　 "query": { "match": { "account_number": 20 } }
　　}'
```

- match_phrase
```
   // address contain mill or lane
　　curl -XPOST 'localhost:9200/bank/_search?pretty' -d '
　　{
  　　"query": { "match": { "address": "mill lane" } }
　　}' 

   // address contain a phrase `mill lane`
   curl -XPOST 'localhost:9200/bank/_search?pretty' -d '
　　{
  　　"query": { "match_phrase": { "address": "mill lane" } }
　　}' 
```
- boolean query
```
   // all should be true in must
　　curl -XPOST 'localhost:9200/bank/_search?pretty' -d '
　　{
　　  "query": {
 　　   "bool": {
    　　  "must": [
     　　   { "match": { "address": "mill" } },
     　　   { "match": { "address": "lane" } }
    　　  ]
    　　}
  　　}
　　}'

   // Any one is true will return
　　curl -XPOST 'localhost:9200/bank/_search?pretty' -d '
　　{
  　　"query": {
  　　  "bool": {
    　　  "should": [
     　　   { "match": { "address": "mill" } },
      　　  { "match": { "address": "lane" } }
     　　 ]
   　　 }
  　　}
　　}'

   // no one should be true
　　curl -XPOST 'localhost:9200/bank/_search?pretty' -d '
　　{
 　　 "query": {
  　　  "bool": {
    　　  "must_not": [
      　　  { "match": { "address": "mill" } },
       　　 { "match": { "address": "lane" } }
      　　]
    　　}
  　　}
　　}'
```

- filter
```
   // return data with balance in the range of [20000, 30000]
　　curl -XPOST 'localhost:9200/bank/_search?pretty' -d '
　　{
　　　　  "query": {
  　　　　  "bool": {
    　　　　  "must": { "match_all": {} },
     　　　　 "filter": {
        　　　　"range": {
          　　"balance": {
          　　  "gte": 20000,
           　　 "lte": 30000
         　　 }
       　　 }
     　　 }
   　　 }
  　　}
　　}'
```

- [aggregations](https://www.elastic.co/guide/en/elasticsearch/reference/2.1/_executing_aggregations.html)

```
// SELECT state, COUNT(*) FROM bank GROUP BY state ORDER BY COUNT(*) DESC
curl -XPOST 'localhost:9200/bank/_search?pretty' -d '
{
  "size": 0,
  "aggs": {
    "group_by_state": {
      "terms": {
        "field": "state"
      }
    }
  }
}'

   // Compare with previous example, return the avg of balance in the result
　　curl -XPOST 'localhost:9200/bank/_search?pretty' -d '
　　{
  　　"size": 0,
  　　"aggs": {
   　　 "group_by_state": {
    　　  "terms": {
       　　 "field": "state"
     　　 },
    　　  "aggs": {
       　　 "average_balance": {
        　　  "avg": {
         　　   "field": "balance"
         　　 }
       　　 }
     　　 }
    　　}
  　　}
　　}' 
```

- search_type=count
```
// similar to SELECT UNIQUE
GET /index_streets/_search?search_type=count
{
 "aggs": {
   "street_values": {
     "terms": {
       "field": "name.raw",
       "size": 0
     }
   }
 } 
} 
```


## Examples
