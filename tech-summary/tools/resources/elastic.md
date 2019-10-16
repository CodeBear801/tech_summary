# ELK tool set

This task records my initial experience with ELK tech stack related with [this task](https://github.com/Telenav/osrm-backend/issues/65).  

## Beats
Filebeats is a light weight shipper tool for collecting logs.  They helps moving logs from different places to logstash or elasticsearch.  

The best way to experience beats is this [example setups](https://github.com/xeraa/elastic-docker) for Elasticsearch, Kibana, Logstash, and Beats with docker-compose.  

## Logstash

![image](https://www.elastic.co/guide/en/logstash/current/static/images/logstash.png)

Examples

```
Example #1 

55.3.244.1 GET /index.html 15824 0.043

%{IP:client} %{WORD:method} %{URIPATHPARAM:request} %{NUMBER:bytes} %{NUMBER:duration}

Example #2

2016-07-11T23:56:42.000+00:00 INFO [MySecretApp.com.Transaction.Manager]:Starting transaction for session -464410bf-37bf-475a-afc0-498e0199f008

%{TIMESTAMP_ISO8601:timestamp} %{LOGLEVEL:log-level} \[%{DATA:class}\]:%{GREEDYDATA:message}
 
Example #3

[2019-08-06T11:47:42 UTC] [info] Used 535602493 speeds from LUA profile or input map

\[%{TIMESTAMP_ISO8601:timestamp} UTC\] \[%{LOGLEVEL:log-level}\] Used %{NUMBER:lua-speed-items} speeds from LUA profile or input map


Example #4
03-30-2017 13:26:13 [00089] TIMER XXX.TimerLog: entType [organization], queueType [output], memRecno = 446323718, audRecno = 2595542711, elapsed time = 998ms

^%{DATE_US:dte}\s*%{TIME:tme}\s*\[%{GREEDYDATA}elapsed time\s*=\s*%{BASE10NUM:time}

Example #5

2015-04-17 16:32:03.805 ERROR [grok-pattern-demo-app,BDS567TNP,2424PLI34934934KNS67,true] 54345 --- [nio-8080-exec-1] org.qbox.logstash.GrokApplicarion : this is a sample message

%{TIMESTAMP_ISO8601:timestamp} *%{LOGLEVEL:level} \[%{DATA:application},%{DATA:minQId},%{DATA:maxQId},%{DATA:debug}] %{DATA:pid} --- *\[%{DATA:thread}] %{JAVACLASS:class} *: %{GREEDYDATA:log}

Example #6

[2019-08-13T05:25:48 UTC] [info] 13-08-2019 05:25:48 0.503655ms 10.189.160.19 - Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36 200 /tile/v1/car/tile(5259,12696,15).mvt


\[%{TIMESTAMP_ISO8601:timestamp} UTC\]\s\[%{LOGLEVEL:log-level}\]\s%{DATE_EU:request-date}\s%{TIME:request-time}\s%{DATA:response_duration}\s%{IP:request-ip}\s\W\s(?<user-agent>.+?(?=\s\d{3}\s))\s(?<response-code>\d{3})\s%{URIPATH:uri}

```


### Debug grok

For testing single matching, you could use Kibana->DevTools-> Debug grok.

If you finished your config file and you want to check format, you could use following command
```
/bin/logstash --config.test_and_exit -f xxxx.your.conf
```

### Logstash links
- [groks pattern github](https://github.com/logstash-plugins/logstash-patterns-core/blob/master/patterns/grok-patterns)  <span>&#9733;</span><span>&#9733;</span><span>&#9733;</span>
- [elastic doc groks filter plugin](https://www.elastic.co/guide/en/logstash/current/plugins-filters-grok.html)
-  [elastic doc custom grok pattern](https://www.elastic.co/guide/en/logstash/current/plugins-filters-grok.html#_custom_patterns)
- [elastic doc all plugins](https://www.elastic.co/guide/en/logstash/current/filter-plugins.html)
- [elastic doc json filter plugin](https://www.elastic.co/guide/en/logstash/current/plugins-filters-json.html)

- [Using Multiple Grok Statements to Parse a Java Stack Trace](https://dzone.com/articles/using-multiple-grok-statements)  <span>&#9733;</span><span>&#9733;</span><span>&#9733;</span>
- [A Beginner’s Guide to Logstash Grok](https://logz.io/blog/logstash-grok/) 
- [How to debug logstash](https://logz.io/blog/debug-logstash/)
- [Custom Regex Patterns in Logstash](https://medium.com/statuscode/using-custom-regex-patterns-in-logstash-fa3c5b40daab)
- [Logstash grok filter tutorial patterns](https://qbox.io/blog/logstash-grok-filter-tutorial-patterns)



## Elasticsearch

```

// delete index
curl -XDELETE url:80/osrm-test

// check result
curl -XGET 'url:80/osrm-test/_search?pretty'

// query for specific tags
curl -XGET 'url:80/osrm-test/_search?pretty=true&q=lua-speed-items'

// query for elastic mapping
// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping.html
curl -XGET 'url:80/osrm-test/_mapping?pretty'

```

### Elasticsearch useful links
- [elastic doc search](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-search.html)


## Kibana

### Kibana Useful links
- [elastic doc xy-chart](https://www.elastic.co/guide/en/kibana/current/xy-chart.html)



## Issues

### Error related with logstash match pattern
Problem description: Logstash Configuration Error: “Expected one of #,”
Root cause: Finally, I found its due to I past regular expression from microsoft onenote which secretly changed the character set, especially for space.
Related links: 
https://discuss.elastic.co/t/logstash-configuration-error-expected-one-of/150827
https://discuss.elastic.co/t/configurationerror-message-expected-one-of-input-filter-output-at-line-1-column-1/156154
https://discuss.elastic.co/t/logstash-configurationerror-message-expected-one-of-at-line/128999

### Error related Kibana
Problem description: While trying to generate a virtualization graph for traffic-speed-value(key-value), I found following errors
```
Visualize: Fielddata is disabled on text fields by default. Set fielddata=true on [lua-speed-items] in order to load fielddata in memory by uninverting the inverted index. Note that this can however use significant memory. Alternatively use a keyword field instead.
```
Solution: From link [here](https://www.elastic.co/guide/en/elasticsearch/reference/current/fielddata.html), it mentioned how to handle fielddata for text, but it also mentioned text is not for such purpose.  
Originally, I use following strategy to change lua-speed-items from string to int.z
```
      mutate {
         convert => { "lua-speed-items" => "integer" }
      }
```
Finally use [this way](https://discuss.elastic.co/t/field-turns-into-text-in-elasticsearch-after-being-mapped-as-int-in-grok-and-being-mutated-into-integer/112381)
```
pattern:lua-speed-items:int
```



