# Kafka 设计解析 Notes
By 郭俊

## links

https://www.infoq.cn/article/kafka-analysis-part-1/
https://www.infoq.cn/article/kafka-analysis-part-2
https://www.infoq.cn/article/kafka-analysis-part-3
https://www.infoq.cn/article/kafka-analysis-part-4
https://www.infoq.cn/article/kafka-analysis-part-5
https://www.infoq.cn/article/kafka-analysis-part-6
https://www.infoq.cn/article/kafka-analysis-part-7
https://www.infoq.cn/article/kafka-analysis-part-8

## Terminology
- **Broker** Kafka 集群包含一个或多个服务器，这种服务器被称为 broker
- **Topic** 每条发布到 Kafka 集群的消息都有一个类别，这个类别被称为 Topic。（物理上不同 Topic 的消息分开存储，逻辑上一个 Topic 的消息虽然保存于一个或多个 broker 上但用户只需指定消息的 Topic 即可生产或消费数据而不必关心数据存于何处）
- **Partition** Partition 是物理上的概念，每个 Topic 包含一个或多个 Partition.
- **Producer** 负责发布消息到 Kafka broker
- **Consumer** 消息消费者，向 Kafka broker 读取消息的客户端。
- **Consumer Group** 每个 Consumer 属于一个特定的 Consumer Group（可为每个 Consumer 指定 group name，若不指定 group name 则属于默认的 group）

## Kafka 拓扑结构

<img src="../resources/kafka_design_guojun_topo.png" alt="kafka_design_guojun_topo.png" width="600"/>
<br/>

## Topic & partition

Topic 在逻辑上可以被认为是一个 queue，每条消费都必须指定它的 Topic，可以简单理解为必须指明把这条消息放进哪个 queue 里。为了使得 Kafka 的吞吐率可以线性提高，物理上把 Topic 分成一个或多个 Partition，每个 Partition 在物理上对应一个文件夹，该文件夹下存储这个 Partition 的所有消息和索引文件。若创建 topic1 和 topic2 两个 topic，且分别有 13 个和 19 个分区，则整个集群上会相应会生成共 32 个文件夹（本文所用集群共 8 个节点，此处 topic1 和 topic2 replication-factor 均为 1），如下图所示。

<img src="../resources/kafka_design_guojun_partition.png" alt="kafka_design_guojun_partition.png" width="600"/>
<br/>
每个日志文件都是一个 log entrie 序列，每个 log entrie 包含一个 4 字节整型数值（值为 N+5），1 个字节的 "magic value"，4 个字节的 CRC 校验码，其后跟 N 个字节的消息体。每条消息都有一个当前 Partition 下唯一的 64 字节的 offset，它指明了这条消息的起始位置。磁盘上存储的消息格式如下：

<img src="../resources/kafka_design_guojun_pysical.png" alt="kafka_design_guojun_pysical.png" width="600"/>
<br/>
这个 log entries 并非由一个文件构成，而是分成多个 segment，每个 segment 以该 segment 第一条消息的 offset 命名并以“.kafka”为后缀。另外会有一个索引文件，它标明了每个 segment 下包含的 log entry 的 offset 范围，如下图所示。
kafka_design_guojun_segment

因为每条消息都被 append 到该 Partition 中，属于顺序写磁盘，因此效率非常高（经验证，顺序写磁盘效率比随机写内存还要高，这是 Kafka 高吞吐率的一个很重要的保证）。

<img src="../resources/kafka_design_guojun_partition_write.png" alt="kafka_design_guojun_partition_write.png" width="600"/>
<br/>

对于传统的 message queue 而言，一般会删除已经被消费的消息，而 Kafka 集群会保留所有的消息，无论其被消费与否。当然，因为磁盘限制，不可能永久保留所有数据（实际上也没必要），因此 Kafka 提供两种策略删除旧数据。一是基于时间，二是基于 Partition 文件大小


因为 offet 由 Consumer 控制，**所以 Kafka broker 是无状态的**，它不需要标记哪些消息被哪些消费过，也不需要通过 broker 去保证同一个 Consumer Group 只有一个 Consumer 能消费某一条消息，因此也就不需要锁机制

## Producer 消息路由

Producer 发送消息到 broker 时，会根据 Paritition 机制选择将其存储到哪一个 Partition。如果 Partition 机制设置合理，所有消息可以均匀分布到不同的 Partition 里，这样就实现了负载均衡。如果一个 Topic 对应一个文件，那这个文件所在的机器 I/O 将会成为这个 Topic 的性能瓶颈，而有了 Partition 后，不同的消息可以并行写入不同 broker 的不同 Partition 里，极大的提高了吞吐率。可以在 $KAFKA_HOME/config/server.properties 中通过配置项 num.partitions 来指定新建 Topic 的默认 Partition 数量，也可在创建 Topic 时通过参数指定，同时也可以在 Topic 创建之后通过 Kafka 提供的工具修改。

在发送一条消息时，可以指定这条消息的 key，Producer 根据这个 key 和 Partition 机制来判断应该将这条消息发送到哪个 Parition。Paritition 机制可以通过指定 Producer 的 paritition. class 这一参数来指定，该 class 必须实现 kafka.producer.Partitioner 接口

<img src="../resources/kafka_design_guojun_example_code.png" alt="kafka_design_guojun_example_code.png" width="600"/>
<br/>


