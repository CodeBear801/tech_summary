
# Spark HA

For components in spark, please go to [Spark architecture page](./spark_arch.md)

<img src="https://user-images.githubusercontent.com/16873751/108147035-17f6a980-7083-11eb-8288-612a2c506833.png" alt="spark_arch" width="600"/> 
<br/>

## Driver crash

- 最简单的实现: 所有相关的 executor 和 task 都会被清理，不尝试恢复，
- yarn-cluster 模式下会重启 driver，重跑所有 task.
- 一般说来，尽量保留 driver 已完成的工作，driver 挂了，应该依赖持久化存储和与 executor 的通信，恢复 driver 的状态，重现建立对正在运行的 executor，task 等的管理。而 executor 与 driver 通信失败时应该有重试机制，并且当重试失败时不应该退出，而应该等待 driver 恢复后重连，并把自己的状态告诉 driver，只有在超过 driver 恢复时间时才自行退出

<img src="https://user-images.githubusercontent.com/16873751/108147268-6a37ca80-7083-11eb-844f-4ca3b1e9f781.png" alt="spark_arch" width="600"/> 

(from: https://www.cnblogs.com/juncaoit/p/6542902.html)
<br/>

## Executor crash

<img src="https://user-images.githubusercontent.com/16873751/108147316-76238c80-7083-11eb-83ef-5fba2d049d28.png" alt="spark_arch" width="600"/> 

(from: https://www.cnblogs.com/juncaoit/p/6542902.html)
<br/>

## Task crash

<img src="https://user-images.githubusercontent.com/16873751/108147329-7cb20400-7083-11eb-97d1-a1352093be8a.png" alt="spark_arch" width="600"/> 

(from: https://www.cnblogs.com/juncaoit/p/6542902.html)
<br/>


## Master crashed

<img src="https://user-images.githubusercontent.com/16873751/108148109-fe566180-7084-11eb-9fc2-71a88c5f67a7.png" alt="spark_arch" width="600"/> 

(from: https://blog.csdn.net/zc19921215/article/details/107732908)
<br/>

## More info
- [Spark 容错机制](https://liyichao.github.io/posts/spark-%E5%AE%B9%E9%94%99%E6%9C%BA%E5%88%B6.html)
- [RDD之七：Spark容错机制](https://www.cnblogs.com/duanxz/p/6329675.html)
- 