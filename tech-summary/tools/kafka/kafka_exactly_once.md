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