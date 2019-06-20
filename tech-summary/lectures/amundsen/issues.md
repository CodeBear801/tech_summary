# Issues and solutions during set up amundsen

Here records issues I met during set up amundsenfrontendlibrary/amundsendatabuilder/amundsenmetadatalibrary/amundsensearchlibrary based on code and docker image.  
Here is the instruction I followed: [link](https://github.com/lyft/amundsenfrontendlibrary/blob/master/docs/installation.md#install-standalone-application-directly-from-the-source)

## Amundsenfrontend
1. Docker image for frontend is not working  
  Issue is similar to this one [issues-64](https://github.com/lyft/amundsenfrontendlibrary/issues/64)
  Solution: Build frontend image on local and modify docker file, modify yml file to depend on local file during docker-compose

```yml
  amundsenfrontend:
      image: amundsen-frontend:local
```
  Notes: Amudsen team update docker image on cloud side and fixed this issue in (image: amundsendev/amundsen-frontend:1.0.5).  I met following issues during generating forntend docker file in local
```bash
docker build -f public.Dockerfile -t amundsen-frontend:local .
```

2. Failed to build with node
  Issue is similar to this one [issues-6](https://github.com/lyft/amundsenfrontendlibrary/issues/6)  
  Solution: Use node 8.x.x or 10.x.x, not 11.x.x  

3. Node build failed
  Error info:  
```
ERROR in Missing binding /node_modules/node-sass/vendor/darwin-x64-11/binding.node
Node Sass could not find a binding for your current environment: OS X 64-bit with Node 0.10.x
```
  Solution: [stackoverflow](https://stackoverflow.com/questions/37986800/node-sass-could-not-find-a-binding-for-your-current-environment/47197646)  

  Sample of docker file, follow the comments from [issues-64](https://github.com/lyft/amundsenfrontendlibrary/issues/64) with two modifications: 1. add 'npm rebuild node-sass', 2. change https://deb.nodesource.com/setup_11.x to setup_8.x 
```docker
FROM python:3
WORKDIR /app
COPY . /app
RUN pip3 install -r requirements3.txt
RUN apt-get update \
    && apt-get -y install curl gnupg \
    && curl -sL https://deb.nodesource.com/setup_8.x  | bash - \
    && apt-get -y install nodejs \
    && cd amundsen_application/static \
    && npm rebuild node-sass \
    && npm install \
    && npm run build
RUN python3 setup.py install

ENTRYPOINT [ "python3" ]
CMD [ "amundsen_application/wsgi.py" ]

```
  More information about [node distribution](https://github.com/nodesource/distributions)   



## Integration
1. Elestic search gc issue
  Description: While using docker compose to init 4 components on macbookpro, I alwasy failed to connect neo4j and elasticsearch, by checking logs I found following information:  
```
es_amundsen         | OpenJDK 64-Bit Server VM warning: Option UseConcMarkSweepGC was deprecated in version 9.0 and will likely be removed in a future release.
es_amundsen exited with code 137
     
```
  Solution: I feel its due to elasticsearch run out of memory of docker with 6GB memory allocated for it, I tried with following steps:
 * Setting up enviroment on macpro with 32GB physical memory allocated for docker
 * Modify yml file of docker-amudensen.yml in amundsenfrontend folder, add limit of heap size for elestic search ([reference](https://github.com/10up/wp-local-docker/issues/6))
```docker
elasticsearch:
      environment:
          ES_JAVA_OPTS: "-Xms10g -Xmx10g"
```

2. Authetication issue of neo4j
During try with following command in databuilder
```python
python example/scripts/sample_data_loader.py
```
I always met with the issue like
```
ServiceUnavailable: Failed to establish connection to xxxx,xxx
neo4j connectionrefusederror errno 61 connection refused mac
```
Soulution: 
- Update node4j's password.  Go to http://192.168.99.100:7474/browser/ or http://0.0.0.0:7474/browser, input old password and then update new password with 'test'.  Please make sure password in docker-amudensen.yml file is also test.
```
  neo4j:
      environment:
        - CREDENTIALS_PROXY_USER=neo4j
        - CREDENTIALS_PROXY_PASSWORD=test
```
- I also tried to disable authetication for neo4j, but seems it didn't work
go into container
```
docker exec -it 0f495da1c0dce3330c41d0386170423ff23be742e0b2031efbb040fe0ca9e2c7  /bin/bash
```
Modify conf/neo4j.conf, remove following line [ref](https://neo4j.com/docs/operations-manual/current/authentication-authorization/enable/)
```
dbms.security.auth_enabled=false
```


## Test
1. Test frontend  
http://0.0.0.0:5000  
http://192.168.99.100:5000  
  
2. Test neo4j  
http://localhost:7474/browser/  
http://localhost:5000/table_detail/gold/dynamo/test_schema/test_table2  




