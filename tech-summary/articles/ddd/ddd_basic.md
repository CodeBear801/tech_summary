
所谓领域，即是一个组织的业务开展方式，业务价值便体现在其中.**一个软件系统是否真正可用是通过它所提供的业务价值体现出来的。**

## DDD之战略设计

### 领域和子域（Domain/Subdomain）

领域并不是多么高深的概念，比如，一个保险公司的领域中包含了保险单、理赔和再保险等概念；一个电商网站的领域包含了产品名录、订单、发票、库存和物流的概念。

哪些概念应该建模在哪些子系统里面？

各个子系统之间的应该如何集成？两个系统之间的集成涉及到基础设施和不同领域概念在两个系统之间的翻译

### 限界上下文（Bounded Context）

在一个领域/子域中，我们会创建一个概念上的领域边界，在这个边界中，任何领域对象都只表示特定于该边界内部的确切含义。这样边界便称为限界上下文。限界上下文和领域具有一对一的关系。

将一个限界上下文中的所有概念，包括名词、动词和形容词全部集中在一起，我们便为该限界上下文创建了一套通用语言。通用语言是一个团队所有成员交流时所使用的语言，业务分析人员、编码人员和测试人员都应该直接通过通用语言进行交流。

上下文各个域之间的集成问题，可以用防腐层(ACL,Anti corruption Layer)来隔离。防腐层负责与外部服务提供方打交道，还负责将外部概念翻译成自己的核心领域能够理解的概念。

另外还有共享内核、开放主机服务可以来做隔离


### 架构风格（Architecture）

抽象不应该依赖于细节，细节应该依赖于抽象。

## DDD之战术设计

### 实体 vs 值对象（Entity vs Value Object）

实体表示那些具有生命周期并且会在其生命周期中发生改变的东西

而值对象则表示起描述性作用的并且可以相互替换的概念

实体有unique id, 值对象重载了equals方法

```
• `Entity`: Objects that have a distinct identity that runs through time and 
different representations. You also hear these called "reference objects".
• `Value Object`: Objects that matter only as the combination of their attributes. 
Two value objects with the same values for all their attributes are considered 
equal. I also describe value objects in P of EAA.
• `Service`: A standalone operation within the context of your domain. A Service 
Object collects one or more services into an object. Typically you will have 
only one instance of each service object type within your execution context.

Entities are usually big things like Customer, Ship, Rental Agreement. Values 
are usually little things like Date, Money, Database Query. Services are usually 
accesses to external resources like Database Connection, Messaging Gateway, 
Repository, Product Factory.

One clear division between entities and values is that values override the 
equality method (and thus hash) while entities usually don't. 
```

### 聚合（Aggregate）

一个聚合中可以包含多个实体和值对象，因此聚合也被称为根实体。聚合是持久化的基本单位，它和资源库具有一一对应的关系。
   - 比如一系列订单项组成的订单
   - 比如一辆汽车（Car）包含了引擎（Engine）、车轮（Wheel）和油箱（Tank）等组件

 ```java
public class BlogApplicatioinService {

    @Transactional
    public void createBlog(String blogName, String userId) {
        User user = userRepository.userById(userId);
        Blog blog = user.createBlog(blogName);
        blogRepository.save(blog);
    }
}
```

```
A DDD aggregate is a cluster of domain objects that can be treated as a single unit. 
An example may be an order and its line-items, these will be separate objects, but 
it's useful to treat the order (together with its line items) as a single aggregate.
```

### 领域服务（Domain Service）

- 领域服务和上文中提到的应用服务是不同的，领域服务是领域模型的一部分，而应用服务不是。应用服务是领域服务的客户，它将领域模型变成对外界可用的软件系统
- 要对密码进行加密，我们便可以创建一个 PasswordEncryptService 来专门负责此事。


### 资源库（Repository）

```java
public interface CollectionOrientedUserRepository {
    public void add(User user);
    public User userById(String userId);
    public List allUsers();     
    public void remove(User user); 
   }
```

### 领域事件（Domain Event）

- 领域事件的命名遵循英语中的“名词 + 动词过去分词”格式，即表示的是先前发生过的一件事情。
- 领域事件的额外好处在于它可以记录发生在软件系统中所有的重要修改，这样可以很好地支持程序调试和商业智能化。另外，在 CQRS 架构的软件系统中，领域事件还用于写模型和读模型之间的数据同步。再进一步发展，事件驱动架构可以演变成事件源（Event Sourcing），即对聚合的获取并不是通过加载数据库中的瞬时状态，而是通过重放发生在聚合生命周期中的所有领域事件完成。
- 在DDD中有一条原则：一个业务用例对应一个事务，一个事务对应一个聚合根，也即在一次事务中，只能对一个聚合根进行操作。
- 领域事件给我们带来以下好处： 
    + 解耦微服务（限界上下文）
    + 帮助我们深入理解领域模型
    + 提供审计和报告的数据来源
    + 迈向事件溯源（Event Sourcing）和CQRS等

```java
/**
 * Guava事件发布器实现
 * https://github.com/mymonkey110/event-light
 */
public abstract class GuavaDomainEventPublisher implements DomainEventPublisher {
    private EventBus syncBus = new EventBus(identify());
    private EventBus asyncBus = new AsyncEventBus(identify(), Executors.newFixedThreadPool(1));

    @Override
    public void register(Object listener) {
        syncBus.register(listener);
        asyncBus.register(listener);
    }

    @Override
    public void publish(DomainEvent event) {
        syncBus.post(event);
    }

    @Override
    public void asyncPublish(DomainEvent event) {
        asyncBus.post(event);
    }

}
```

![image](http://blog.didispace.com/content/images/posts/impl-ddd-event-2-4.png)

- 如果一个业务流程需要贯穿几个不同的受限上下文中，那么可以通过以发布领域事件的方式来避免上游系统耦合下游系统。这种解耦方式收益最大，因为其有利于后期系统间的拆分。
- 如果在同一个受限上下文中，也可以通过发布领域事件的方式来达到领域间解耦。


如何发布，接受以及测试event


### [深入浅出Event Sourcing和CQRS](https://zhuanlan.zhihu.com/p/38968012)

CQRS，是 Command Query Responsibility Segregation的缩写，也就是通常所说的读写隔离。在上面，我们说，为了性能考虑，将聚合对象的数据状态用物化视图的形式保存，可以用于数据的查询操作，也就是我们把数据的更新与查询的流程隔离开来。我们通过事件来更新聚合对象的数据状态，同时由另一个处理器处理相同的事件，来更新物化视图的数据。

![image](https://pic1.zhimg.com/80/v2-35249fb2693f44bbe4bf48ea6755c55c_1440w.jpg)


## More info
领域驱动设计实现之路
https://www.infoq.cn/article/implementation-road-of-domain-driven-design/

在微服务中使用领域事件
https://www.cnblogs.com/davenkin/p/microservices-and-domain-events.html

软件架构图的艺术
https://www.infoq.cn/article/crafting-architectural-diagrams/

用于软件架构的 C4 模型
https://www.infoq.cn/article/C4-architecture-model/

Power Use of Value Objects in DDD
https://www.infoq.com/presentations/Value-Objects-Dan-Bergh-Johnsson/

Matin的一系列文章
https://martinfowler.com/tags/domain%20driven%20design.html

一组Ref文章
http://blog.didispace.com/tags/%E9%A2%86%E5%9F%9F%E9%A9%B1%E5%8A%A8%E8%AE%BE%E8%AE%A1/

