# How Kafka achieve Exactly Once

## Problem set

Treat kafka as black box

<img src="https://user-images.githubusercontent.com/16873751/105798871-01f43e00-5f48-11eb-9f87-c709635b32b4.png" alt="kafka_streaming" width="600"/>
<br/>

We want:
    - message has been correctly be published into Kafka, 
    - then all the status has been correctly be updated
    - then result be subscribed by consumers
    - then input be marked as consumed

All or nothing!

## Problems 

<img src="https://user-images.githubusercontent.com/16873751/105787881-3827c280-5f34-11eb-87fd-e9c43c975ff2.png" alt="kafka_streaming" width="600"/>
<br/>

Solution:  

- Write to input topic
- Write to change log topics
- Write to output topic
- Write to offset commit log

Key points

- write path
   + idempotency
   + atomic multi-partition writes

- Read path
   + only read commited data
 
- Exactly once processing


<span style="font-size:500%;color:blue;">&starf;</span>How to avoid publish duplicate data into Kafka?    

Broker side de-duplication via producer IDS and sequence numbers(persist in logs)  


<img src="https://user-images.githubusercontent.com/16873751/105799650-e25e1500-5f49-11eb-8845-684d406e624d.png" alt="kafka_streaming" width="600"/>
<br/>



<span style="font-size:500%;color:blue;">&starf;</span>How to avoid problems when transaction happened for multiple times?    

Transactions:
- Atomic multi-partition writes
- 2 phase commit

Success situation:  

<img src="https://user-images.githubusercontent.com/16873751/105799865-5e585d00-5f4a-11eb-890f-0634adf97f20.png" alt="kafka_streaming" width="600"/>
<br/>
Begin txn -> send -> prepare commit ->  commit txn(committed)

Abort situation  

<img src="https://user-images.githubusercontent.com/16873751/105799881-644e3e00-5f4a-11eb-9c08-6a25937cdf5c.png" alt="kafka_streaming" width="600"/>
<br/>





<img src="https://user-images.githubusercontent.com/16873751/105799891-69ab8880-5f4a-11eb-9f5f-7ca893ece4f5.png" alt="kafka_streaming" width="600"/>
<br/>

Intermediate status just replay of logs, based on changelogs if we got abort, system will keep correct status

<span style="font-size:500%;color:blue;">&starf;</span>How to handle the situation of setting read-commit-log incorrectly

<img src="https://user-images.githubusercontent.com/16873751/105800194-0a01ad00-5f4b-11eb-8eaa-bb2b67d20fb9.png" alt="kafka_streaming" width="600"/>
<br/>

Producer will guarantee to update consumer__offset and make them in the transaction


<img src="https://user-images.githubusercontent.com/16873751/105800221-14bc4200-5f4b-11eb-8158-dd0fba9541fb.png" alt="kafka_streaming" width="600"/>
<br/>
If something goes wrong, we will guarantee abort marker will be put everywhere.

*** 

### Producer side's offset management


<img src="https://user-images.githubusercontent.com/16873751/106366607-dd0a1d00-62f1-11eb-841d-6790991b16b0.png" alt="kafka_streaming" width="600"/>
<br/>

more info please go to [How kafka achieve HA](./kafa_ha.md)  

*** 

### Consumer side's offset management
消费的位移信息存在了kafka内部的主题__consumer_offset中， 存储(持久化)并且提交， 2pc  

<img src="https://user-images.githubusercontent.com/16873751/106366482-0d04f080-62f1-11eb-9535-423cb5f33cd0.png" alt="kafka_consumer" width="600"/>
<br/>

<img src="https://user-images.githubusercontent.com/16873751/106366489-142bfe80-62f1-11eb-977b-ecd68586c79d.png" alt="kafka_consumer" width="600"/>
<br/>
kafka提供了position(TopicPartition) 以及 committed(TopicPartition) 方法，来获得下次拉取信息的便宜以及提交消费过的位移  

<img src="https://user-images.githubusercontent.com/16873751/106366490-19894900-62f1-11eb-8750-d6e52b107bdd.png" alt="kafka_consumer" width="600"/>
<br/>


- kafka默认是自动提交，就是定时更新，使用自动提交会有两个潜在的问题
   + 如果拿到了 [x+2, x+8] 就提交x + 8, 那么5 ～ 8 没有被处理
   + 如果提交x+ 2 ，那么又可能会重复消费2 ~ 5
自动提交是在poll方法的逻辑中完成的，每次真正向服务器发起拉取请求之前都会检查是否可以进行位移提交，每5秒一次

比较好的方式是手动提交，当用户完成了相关操作后，比如落盘，写入数据库，或者写入本地缓存，完成更复杂的业务才提交offset  
commitAsync()提交失败，重试时如何avoid重复提交: 递增序号来维护异步提交的顺序  
problem: 如果某一次异步提交的消费位移为x，但是提交失败，在下一次又异步提交了消费位移为x+y并成功。如果简单重试，那么重复提交x，如果修改成功，那又吧消费移回了x，可能造成潜在的重复消费     
solution: 每次位移成功提交之后就递增序列号的相应的值。如果遇到位移提交失败需要重试的时候，可以检查所提交的位移和序号值的大小。如果前者小于后者，那就说明有更大的位移已经提交了，不需要在进行本次重试；如果两者相同，说明可以重试提交  

## More info
- [Don’t Repeat Yourself Introducing Exactly Once Semantics in Apache Kafka - Kafka Summit 2018
](https://www.youtube.com/watch?v=zm5A7z95pdE&t=697s)
