# Elasticsearch template

[Difference between mapping and template](https://stackoverflow.com/a/52823881).  Template helps to create index implicitly by writing data to it.

If I want add indexes created start with `logstash-` applied specific settings, such as
- disabling the _all fields search
- set the default attribute to @message
- To saving sapce and indexing time, I disallowed create index for source/source_host/source_path etc


I could do as following  
1. Define logstash-template.json
```json
{
    "template": "logstash-*",
    "settings" : {
        "index.number_of_shards" : 3,
        "index.number_of_replicas" : 1,
        "index.query.default_field" : "@message",
        "index.routing.allocation.total_shards_per_node" : 2,
        "index.auto_expand_replicas": false
    },
    "mappings": {
        "_default_": {
            "_all": { "enabled": false },
            "_source": { "compress": false },
            "dynamic_templates": [
                {
                    "fields_template" : {
                        "mapping": { "type": "string", "index": "not_analyzed" },
                        "path_match": "@fields.*"
                    }
                },
                {
                    "tags_template" : {
                        "mapping": { "type": "string", "index": "not_analyzed" },
                        "path_match": "@tags.*"
                    }
                }
            ],
            "properties" : {
                "@fields": { "type": "object", "dynamic": true, "path": "full" },
                "@source" : { "type" : "string", "index" : "not_analyzed" },
                "@source_host" : { "type" : "string", "index" : "not_analyzed" },
                "@source_path" : { "type" : "string", "index" : "not_analyzed" },
                "@timestamp" : { "type" : "date", "index" : "not_analyzed" },
                "@type" : { "type" : "string", "index" : "not_analyzed" },
                "@message" : { "type" : "string", "analyzer" : "whitespace" }

             }
        }
    }
}
```
2. apply the settings to cluster
```
curl -XPUT 'http://localhost:9200/_template/template_logstash/' -d @logstash-template.json
```
