
## Problem
- When graph is too large to feed into single machine.  
- Parallel graph model like BGM, lack of fault handling ability

## Overview

The high-level organization of Pregel programs is inspired by Valiant’s Bulk Synchronous Parallel model([wiki](https://en.wikipedia.org/wiki/Bulk_synchronous_parallel), [quora](https://www.quora.com/What-is-BSP-Bulk-Synchronous-Parallel)). Pregel computations consist of a sequence of iterations, called supersteps. During a superstep the framework invokes a userdefined function for each vertex, conceptually in parallel.  
The function specifies behavior at a single vertex V and a single superstep S. It can read messages sent to V in superstep S − 1, send messages to other vertices that will be
received at superstep S + 1, and modify the state of V and its outgoing edges. Messages are typically sent along outgoing edges, but a message may be sent to any vertex whose identifier is known.  
The vertex-centric approach is reminiscent of MapReduce in that users focus on a local action, processing each item independently, and the system composes these actions to lift
computation to a large dataset. By design the model is well suited for distributed implementations: it doesn’t expose any mechanism for detecting order of execution within a superstep, and all communication is from superstep S to superstep S + 1.

## C++ API

<img src="https://user-images.githubusercontent.com/16873751/95389451-b3436b00-08a8-11eb-98e7-1a17563e2be2.png" alt="pregel_paper_pic3" width="500"/>  <br/>

- 顶点类Vertex包含vertices、edges、messges三种相关数据，采用protocol buffer实现易变类型
- 用户通过重写Compute()函数（C++中的虚函数）定义每个superstep中顶点进行的操作；GetValue()和MutableValue()函数分别得到和修改顶点关联值
- Message Passing：发送消息时根据目标顶点是否在本地使用不同的方式；master用barrier机制实现
- Combiners：（某些应用）将收到的消息进行合并，减少信息发送量，默认不启动
- Aggregators：（如min、max、sum）每个superstep钟每个顶点提供一个值给aggregator使用，系统通过reduce操作得到一个全局值，该值可被下一个superstep中的所有顶点使用
- Topology Mutations:在算法执行过程中可以改变拓扑结构，使用lazy机制；
- Input and Output：Pregel提供常见格式文件的读写，通过继承Reader和Wirter类实现特别的需求



## More Info
- Paper [EN](https://kowshik.github.io/JPregel/pregel_paper.pdf) [CN](https://developer.aliyun.com/article/4761)
- [Pregel In Graphs - Models and Instances](https://www.slideshare.net/ChaseZhang3/pregel-in-graphs-models-and-instances)
 

