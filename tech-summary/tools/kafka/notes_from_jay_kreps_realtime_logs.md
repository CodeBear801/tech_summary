# notes about real time logs by Jay Kreps

https://engineering.linkedin.com/distributed-systems/log-what-every-software-engineer-should-know-about-real-time-datas-unifying

## logs
- what is log 
- how to use logs for data integration, real time processing, and system building.

## What is log
It is an append-only, totally-ordered sequence of records ordered by time



<img src="../resources/real_time_logs_jay_kreps_logs1.png" alt="real_time_logs_jay_kreps_logs1.png" width="400"/>
<br/>

## why log is important
The answer is that logs have a specific purpose: they record what happened and when
The two problems a log solves—ordering changes and distributing data—are even more important in distributed data systems.

One of the beautiful things about this approach is that the time stamps that index the log now act as the clock for the state of the replicas—you can describe each replica by a single number, the timestamp for the maximum log entry it has processed. This timestamp combined with the log uniquely captures the entire state of the replica.


<img src="../resources/real_time_logs_jay_kreps_logs2.png" alt="real_time_logs_jay_kreps_logs2.png" width="400"/>
<br/>


State-Machine-> active active -> The active-active approach might log out the transformations to apply, say "+1", "*2", etc. Each replica would apply these transformations and hence go through the same set of values.

Primary-Backup -> Active-Passive ->  a single master execute the transformations and log out the result, say "1", "3", "6", etc. 

A log, after all, represents a series of decisions on the "next" value to append. 


