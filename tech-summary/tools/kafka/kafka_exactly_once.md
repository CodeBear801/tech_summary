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

## More info
- [Don’t Repeat Yourself Introducing Exactly Once Semantics in Apache Kafka - Kafka Summit 2018
](https://www.youtube.com/watch?v=zm5A7z95pdE&t=697s)