# Raft

## Presentation from Deigo
[An Introduction to Raft (CoreOS Fest 2015)](https://www.youtube.com/watch?v=6bBggO6KN_k)  


<img src="resources/pictures/raft_deigo_1.png" alt="raft_deigo_1" width="500"/>  <br/>

NextIndex: what to send to this specific follower to sync  
Once nextindex == end of server.log, then appendentries become heartbeat only  

<img src="resources/pictures/raft_deigo_2.png" alt="raft_deigo_2" width="500"/>  <br/>

Leader could mark one entry committed once it be replicated by majority of clusters  
Sever will send his committed index to the followers, followers could upgrade status  

<img src="resources/pictures/raft_deigo_3.png" alt="raft_deigo_3" width="500"/>  <br/>

Missing entry  -> S4, get all result of S2  

<img src="resources/pictures/raft_deigo_4.png" alt="raft_deigo_4" width="500"/>  <br/>

Conflict entry -> S1, S2 will just remove S1's entries  
               Term 3 will become leader  

<br/>

**SAFETY**


<img src="resources/pictures/raft_deigo_5.png" alt="raft_deigo_5" width="500"/>  <br/>

Even S5 become leader, but due to its TERM is low, will change to follower  

<img src="resources/pictures/raft_deigo_6.png" alt="raft_deigo_6" width="500"/>  <br/>

Middle server will never vote bottom server, index    
Upper server will never vote middle server, term   

<img src="resources/pictures/raft_deigo_7.png" alt="raft_deigo_7" width="500"/>  <br/>



## Presentation from John Ousterhout

[Designing for Understandability: The Raft Consensus Algorithm](https://www.youtube.com/watch?v=vYp4LYbnnW8)

<img src="resources/pictures/raft_john_1.png" alt="raft_john_1" width="500"/>  <br/>

**Replicated state machine**  
<img src="resources/pictures/raft_john_2.png" alt="raft_john_2" width="500"/>  <br/>

Paxos, proposers and acceptors, proposer will accept the one with highest proposal number  

**Why Raft**  
<img src="resources/pictures/raft_john_3.png" alt="raft_john_3" width="500"/>  <br/>

**Main topics in Raft**  
<img src="resources/pictures/raft_john_4.png" alt="raft_john_4" width="500"/>  <br/>

<img src="resources/pictures/raft_john_5.png" alt="raft_john_5" width="500"/>  <br/>

<img src="resources/pictures/raft_john_6.png" alt="raft_john_6" width="500"/>  <br/>
Abstract of time, used to determine absolute items  

<img src="resources/pictures/raft_john_7.png" alt="raft_john_7" width="500"/>  <br/>

<img src="resources/pictures/raft_john_8.png" alt="raft_john_8" width="500"/>  <br/>

Safety: nothing bad happens  
Liveness: something good happens  

<img src="resources/pictures/raft_john_9.png" alt="raft_john_9" width="500"/>  <br/>

**Log**  
<img src="resources/pictures/raft_john_10.png" alt="raft_john_10" width="500"/>  <br/>



<img src="resources/pictures/raft_john_11.png" alt="raft_john_11" width="500"/>  <br/>


Two entries in the same position has the same term, could guarantee they have the same command  

<img src="resources/pictures/raft_john_12.png" alt="raft_john_12" width="500"/>  <br/>

Case1: term matched  
Case 2: follower didn't match with server's term, server knows dismatch then he will jump 1 term earlier  

<img src="resources/pictures/raft_john_13.png" alt="raft_john_13" width="500"/>  <br/>

How to guarantee leader holds all entries  
When candidate ask for a vote, it send out its last latest entry.  
**Based on term and index decide for completeness**  