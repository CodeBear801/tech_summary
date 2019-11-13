# Experimentation Using Event-based Systems
Martin Fowler + Toby Clemson | Kafka Summit 2018 Keynote (Experimentation Using Event-based Systems)  
[video](https://www.youtube.com/watch?time_continue=52&v=_RgUxUTuxH4)


Agile Fluency

<img src="resources/imgs/event_martin_toby_agile.png" alt="event_martin_toby_agile" width="400"/>
<br/>


Optimists situation

## Evolution 0

<img src="resources/imgs/event_martin_toby_e0.png" alt="event_martin_toby_e0" width="400"/>
<br/>


https://github.com/infrablocks

<img src="resources/imgs/event_martin_toby_e0_detail.png" alt="event_martin_toby_e0_detail" width="400"/>
<br/>


Postgis  
Communicate sync  

**Microservice: each service has its own datastore**

<img src="resources/imgs/event_martin_toby_e0_hyper_media.png" alt="event_martin_toby_e0_hyper_media" width="400"/>
<br/>


Hyper media:  
- easy to make the move
- change service's boundary
- Only provide one url for the entire system, every other url should be got steps by steps

<img src="resources/imgs/event_martin_toby_e0_hyper_media2.png" alt="event_martin_toby_e0_hyper_media2" width="400"/>
<br/>


<img src="resources/imgs/event_martin_toby_e0_procon.png" alt="event_martin_toby_e0_procon" width="400"/>
<br/>


## Event sourcing

<img src="resources/imgs/event_martin_toby_event_sourcing.png" alt="event_martin_toby_event_sourcing" width="400"/>
<br/>


At any time, could blow away application data and rebuild system based on events, recall from log  

Version control: Git  

## Evolution 1

<img src="resources/imgs/event_martin_toby_e1_1.png" alt="event_martin_toby_e1_1" width="400"/>
<br/>


Record complete information from user

<img src="resources/imgs/event_martin_toby_e1_2.png" alt="event_martin_toby_e1_2" width="400"/>
<br/>


<img src="resources/imgs/event_martin_toby_e1_procon.png" alt="event_martin_toby_e1_procon" width="400"/>
<br/>


## Evolution 2

<img src="resources/imgs/event_martin_toby_e2_1.png" alt="event_martin_toby_e2_1" width="400"/>
<br/>


All cordinate logic lives in consumer, load, replayabiliy

<img src="resources/imgs/event_martin_toby_e2_2.png" alt="event_martin_toby_e2_2" width="400"/>
<br/>

Where would we find event bus: kafka

<img src="resources/imgs/event_martin_toby_e2_3.png" alt="event_martin_toby_e2_3" width="400"/>
<br/>


Keep event feed in json, keep readability

<img src="resources/imgs/event_martin_toby_e2_4.png" alt="event_martin_toby_e2_4" width="400"/>
<br/>


Sync communication

<img src="resources/imgs/event_martin_toby_e2_procon.png" alt="event_martin_toby_e2_procon" width="400"/>
<br/>


## Future

<img src="resources/imgs/event_martin_toby_future.png" alt="event_martin_toby_future" width="400"/>
<br/>


Flexibility
