# Scaling Instagram Infra 
@[Lisa Guo](https://www.linkedin.com/in/lisa-guo-5b56972/)   
[video](https://www.youtube.com/watch?v=hnpzNAPiC0E)

## Take away

## Notes

### scale out

scale_instagram_infra_lisa_arch.png


	- User's request come to Django
	- Jango will fwd to other services
	- Use postgresql to persist data
	- Use memcache as front end cache
	- Use cassandra to cache user feeds   https://blog.pythian.com/cassandra-use-cases/


<img src="resources/imgs/scale_instagram_infra_lisa_postgis.png" alt="scale_instagram_infra_lisa_postgis" width="600"/>

<img src="resources/imgs/scale_instagram_infra_lisa_cassandra.png" alt="scale_instagram_infra_lisa_cassandra" width="600"/>

[Perry] static information recorded in postgresql, for dynamic information recorded in cassandra

<img src="resources/imgs/scale_instagram_infra_lisa_infra2.png" alt="scale_instagram_infra_lisa_infra2" width="600"/>

Django + RabbitMQ as one package
[RabbitMQ vs Celery](https://stackoverflow.com/questions/9077687/why-use-celery-instead-of-rabbitmq)

Why memcache

<img src="resources/imgs/scale_instagram_infra_lisa_memcached.png" alt="scale_instagram_infra_lisa_memcached" width="600"/>


Memcache for single data center

<img src="resources/imgs/scale_instagram_infra_lisa_memcache2.png" alt="scale_instagram_infra_lisa_memcache2" width="600"/>


Memcache for multiple data center (How to avoid stale read)

<img src="resources/imgs/scale_instagram_infra_lisa_memcache3.png" alt="scale_instagram_infra_lisa_memcache3.png" width="600"/>


User R will read stale information about user C's comments
<img src="resources/imgs/scale_instagram_infra_lisa_memcache4.png" alt="scale_instagram_infra_lisa_memcache4.png" width="600"/>


When PostgreSQL replica data, invalid cache  
When User R want to read data again, he will go to database  
But what if there are millions of User R query from DC2  
Select count from media_likes where media_id = 12345  

Solution: cache lease: only first request trigger database load, for following ones wait or use stale value  

<img src="resources/imgs/scale_instagram_infra_lisa_memcache5.png" alt="scale_instagram_infra_lisa_memcache5" width="600"/>




[mark]Don't count servers, make the servers count[/mark]


### scale up

Use as few CPU instructions as possible
How: monitor, analyze, optimize

#### Monitor
- collect cpu based on user request, with meta data(data center, time, …)
                 when you find some jump, it could be regression in new feature

- python cprofile find bottleneck
- 
<img src="resources/imgs/scale_instagram_infra_lisa_pythonprofile.png" alt="scale_instagram_infra_lisa_pythonprofile" width="600"/>


#### Analyze

Analysis based on running specific version(snapshot)
Analysis based on duration of running (regression)

#### Optimize
Different size of images

<img src="resources/imgs/scale_instagram_infra_lisa_imagesize.png" alt="scale_instagram_infra_lisa_imagesize" width="600"/>


**Use C++ to replace python for core part**
Facebook folly's std::future https://instagram-engineering.com/c-futures-at-instagram-9628ff634f49


Number of process is limited due to system memory

<img src="resources/imgs/scale_instagram_infra_lisa_mem.png" alt="scale_instagram_infra_lisa_mem.png" width="600"/>


How to have more processes: reduce code, share more(move config to shared memory, disable GC to avoid private mem allocation)

<img src="resources/imgs/scale_instagram_infra_lisa_asyncio.png" alt="scale_instagram_infra_lisa_asyncio" width="600"/>


