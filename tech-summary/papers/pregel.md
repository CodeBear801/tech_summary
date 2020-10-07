
## Problem
- When graph is too large to feed into single machine, e.g. social graph, how to apply algorithm on it, how to scale  
- Parallel graph model like BGM, lack of fault handling ability

## Overview

The high-level organization of Pregel programs is inspired by Valiant’s Bulk Synchronous Parallel model([wiki](https://en.wikipedia.org/wiki/Bulk_synchronous_parallel), [quora](https://www.quora.com/What-is-BSP-Bulk-Synchronous-Parallel)). Pregel computations consist of a sequence of iterations, called supersteps. During a superstep the framework invokes a userdefined function for each vertex, conceptually in parallel.  
The function specifies behavior at a single vertex V and a single superstep S. It can read messages sent to V in superstep S − 1, send messages to other vertices that will be
received at superstep S + 1, and modify the state of V and its outgoing edges. Messages are typically sent along outgoing edges, but a message may be sent to any vertex whose identifier is known.  
The vertex-centric approach is reminiscent of MapReduce in that users focus on a local action, processing each item independently, and the system composes these actions to lift
computation to a large dataset. By design the model is well suited for distributed implementations: it doesn’t expose any mechanism for detecting order of execution within a superstep, and all communication is from superstep S to superstep S + 1.

### Pregel vs mapreduce

Pregel 的图计算过程与 MapReduce 非常接近:在迭代的每一个步骤当中， 将会以图的节点为中心进行 Map 操作，这意味着在 Map 函数当中我们只有以
某一节点为中心的局部信息， 这包括:
1. 一个 Vertex 和它当前 Attach 的 Message
2. 当前 Vertex 向外指向的 Edge 和 Edge 上的属性
3. 当前 Vertex 在上一步计算当中所接收到的全部 Message
对于图中的每一个 Vertex，在 Pregel 当中的一次 Superstep 包括接收消息、合并消息和发送消息三个步骤。

在大多数算法当中，所有的 Vertex 都进入 Inactive 状态就意味着算法结束。

## C++ API

<img src="https://user-images.githubusercontent.com/16873751/95389451-b3436b00-08a8-11eb-98e7-1a17563e2be2.png" alt="pregel_paper_pic3" width="500"/>  <br/>

- 顶点类Vertex包含vertices、edges、messges三种相关数据，采用protocol buffer实现易变类型
- 用户通过重写Compute()函数（C++中的虚函数）定义每个superstep中顶点进行的操作；GetValue()和MutableValue()函数分别得到和修改顶点关联值
- Message Passing：发送消息时根据目标顶点是否在本地使用不同的方式；master用barrier机制实现
- Combiners：（某些应用）将收到的消息进行合并，减少信息发送量，默认不启动
- Aggregators：（如min、max、sum）每个superstep钟每个顶点提供一个值给aggregator使用，系统通过reduce操作得到一个全局值，该值可被下一个superstep中的所有顶点使用
- Topology Mutations:在算法执行过程中可以改变拓扑结构，使用lazy机制；
- Input and Output：Pregel提供常见格式文件的读写，通过继承Reader和Wirter类实现特别的需求





## System design
1. 将图分区到不同机器进行计算
2. 使用主从模型进行任务调度和管理
3. 使用 Message 缓冲近一步提高通讯吞吐量
   - Message 缓冲是在计算节点(Worker)的层面上提高吞吐量的一个优化。 Message 在 Worker 之间传递时并不是来一个发一个，而是通过缓冲积攒一些
Message，之后以 Batch 的形式批量发送。 这一优化可以减少网络请求的 Overhead。
4. 使用 Checkpoint 和 Confined Recovery 实现容错性
   - Pregel 使用两种方法来实现容错性:
       - Checkpoint 在 Superstep 执行前进行，用来保存当前系统的状态。当某一图分区计算失败但 Worker 仍然可用时， 可以从 Checkpoint 执行快速恢复
       - 当某一 Worker 整体失败当机使得它所记录的全部状态丢失时，新启动的 Worker 可能要重新接收上一步发送出来的消息。 为了避免无限制的递归重
新计算之前的步骤，Pregel 将 Worker 将要发送的每一条消息写入 Write Ahead Log。 这种方法被称为 Confined Recovery


## Examples

### Find connected component 1
[link](pregel_connected_component_example.md)


### Find connected component 2

<img src="https://user-images.githubusercontent.com/16873751/95391641-4205b700-08ac-11eb-9d18-667e9f4aca22.png" alt="pregel_paper_pic3" width="600"/>  <br/>

1. 为每个节点初始化一个唯一的 Message 值作为初始值
2. 在每一个步骤当中，一个 Vertex 将其本身和接收到的 Message 聚合为它们之中的最大值(最小值)
3. 如果 Attach 在某一个 Vertex 上的 Message 在上一步当中变大(变小)了， 它就会把新的值发送给所有相邻的节点，否则它会执行 Vote to halt 来
Inactivate 自己
4. 算法一直执行直到所有节点都 Inactive 为止

上述连通分量的算法假定边都是双向的(可以通过两条相反的边实现)。可以想像， 由于同一连通分量当中的节点都可以互相传播消息，因此最终在同一个 连通分量里的 Vertex， 必定都会拥有这一连通分量内 Message 的最大值(最小值)。这个最后的值就可以作为这一连通分量的 Identifier。

### Single source shortest path

bellman-ford

<img src="https://user-images.githubusercontent.com/16873751/95392311-62824100-08ad-11eb-854a-5bedf4de9e9b.png" alt="shortest_path_1" width="600"/>  <br/>
N2、N3 接收到 N1 发送的消息， 从而更新了自己的消息为更小的值。执行结束后，N1因为没有变化而 Inactive

<img src="https://user-images.githubusercontent.com/16873751/95392327-6ca43f80-08ad-11eb-9848-356d90bbfe8b.png" alt="hortest_path_2" width="600"/>  <br/>
N2 和 N3 户想发送消息，由于 N1 -> N3 -> N2 的路径更短，N2 的 Message 被更新， N3 则变为 Inactive。

<img src="https://user-images.githubusercontent.com/16873751/95392343-72018a00-08ad-11eb-8b8c-bc14db42c906.png" alt="hortest_path_3" width="600"/>  <br/>
N2 在上一步仍然是 Active 的状态，它将会向 N3 发送最后一次消息。 由于 N3 没有更新自己的值，此时图中三个节点都变味 Inactive，算法结束。




## More Info
- Paper [EN](https://kowshik.github.io/JPregel/pregel_paper.pdf) [CN](https://developer.aliyun.com/article/4761)
- Optimization on large scale graph [Connected Components in MapReduce and Beyond](https://research.google/pubs/pub43122/)
- [Pregel In Graphs - Models and Instances](https://www.slideshare.net/ChaseZhang3/pregel-in-graphs-models-and-instances)
- [Pregel（图计算）技术原理](https://cshihong.github.io/2018/05/30/Pregel%EF%BC%88%E5%9B%BE%E8%AE%A1%E7%AE%97%EF%BC%89%E6%8A%80%E6%9C%AF%E5%8E%9F%E7%90%86/)
- [Google's hama](https://github.com/apache/hama)
