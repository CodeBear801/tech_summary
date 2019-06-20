
<!-- TOC -->
- [Transections](#transections)
	- [Keywords](#keywords)
	- [Questions](#questions)
	- [Notes](#notes)
		- [Transactions](#transactions)
			- [Consistent](#consistent)
			- [Isolated](#isolated)
				- [Terms](#terms)
		- [Single-Object and Multi-Object Operations](#single-object-and-multi-object-operations)
			- [Single-object writes](#single-object-writes)
			- [Multi-Object Operations](#multi-object-operations)
		- [Weak isolation levels](#weak-isolation-levels)
			- [Read committed](#read-committed)
				- [Implemetation](#implemetation)
			- [Snapshot isolation](#snapshot-isolation)
				- [Implemetation](#implemetation-1)
				- [lost updates](#lost-updates)
				- [write skew](#write-skew)
				- [hantom write skew.](#hantom-write-skew)
		- [Serialization](#serialization)
			- [Actual Serial Execution](#actual-serial-execution)
			- [Pessimistic Lock/Optimistic Lock](#pessimistic-lockoptimistic-lock)
	- [Reference](#reference)



# Transections

## Keywords

## Questions

## Notes

### Transactions
- Transactions provide guarantees about the behavior of data that are fundamental to the old SQL style of operation.
- Transactions were the initial casualty of the NoSQL movement, though they are starting to make a bit of a comeback.
- **Not all applications need transactions. Not all applications want transactions. And not all transactions are the same.**

<img src="resources/pictures/ddia_c7_acid.png" alt="ddia_c7_acid" width="600"/>  
<br/>

#### Consistent
- Consistent in replication means evantural consistent and read your own write consistent
- Consistent in CAP means linearizability, when there is one client successfully write, after that, other client's read must see the value just be writtern
- Consistent in ACID means database is correct between transection, like money transfer between the two

#### Isolated

**Isolation Level**  
<img src="resources/pictures/ansi-sql-isolation-levels.png" alt="ansi-sql-isolation-levels" width="600"/>  
<br/>

**Isolation levels table**  
<img src="resources/pictures/isolation-levels-table.png" alt="isolation-levels-table" width="600"/>  
<br/>

##### Terms
脏写: 写入和提交之间，又别写入了别的数据.  
脏读：一个事务还未提交，另外一个事务访问此事务修改的数据，并使用，读取了事务中间状态数据。  
幻读：一个事务读取2次，得到的记录条数不一致，由于2次读取之间另外一个事务对数据进行了增删。 (insert & delete)
不可重复读：一个事务读取同一条记录2次，得到的结果不一致，由于在2次读取之间另外一个事务对此行数据进行了修改。(update)  
更新丢失: 其含义为T1要更新x的数据，其首先读取x的值，然后再该值上加1再写回数据库。但是在读取x后，T2写入了新的x并成功提交，而T1还是在老的x值的基础上加1。这样，T2的更新对于T1而言就像被丢弃了一样  

### Single-Object and Multi-Object Operations

#### Single-object writes
Atomicity can be implemented using a log for crash recovery and isolation can be implemented using a lock on each object (allowing only one thread to access an object at any one time).  

#### Multi-Object Operations
Foreign key referernce update/Second level index  
Nosql update several document togehter  


### Weak isolation levels

The strongest possible isolation guarantee is serializable isolation: transactions that run concurrently on the same data are guaranteed to perform the same as they would were they to run serially.  However serializable isolation is costly. Systems skimp on it by offering weaker forms of isolation.  As a result, race conditions and failure modes abound. Concurrent failures are really, really hard to debug, because they require lucky timings in your system to occur.

#### Read committed
- The weakest isolation guarantee is read committed.   
  When reading from the database, you will only see data that has been committed (no dirty reads).  
  When writing to the database, you will only overwrite data that has been committed (no dirty writes).  
- This isolation level prevents dirty reads (reads of data that is in-flight) and dirty writes (writes over data that is in-flight).
- Lack of dirty read safety would allow you to read values that then get rolled back. Lack of dirty write safety would allow you to write values that read to something else in-flight (so e.g. the wrong person could get an invoice for a product that they didn't actually get to buy).  

- Read committed does not prevent **the race condition between two counter increments**.

<img src="resources/pictures/ddia_c7_read_commited_example.png" alt="ddia_c7_read_commited_example" width="600"/>  
<br/>
##### Implemetation

  **Hold a row-level lock** on the record you are writing to.  You could do the same with a read lock. However, there is a lower-impact way. Hold the old value in memory, and issue that value in response to reads, until the transaction is finalized.  If a user performs a multi-object write transaction that they believe to be atomic (say, transferring money between two accounts), then performs a read in between the transaction, what they see may seem anomalous (say, one account was deducted but the other wasn't credited).


#### Snapshot isolation

<img src="resources/pictures/ddia_c7_snapshot_example.png" alt="ddia_c7_snapshot_example" width="600"/>  
<br/>
Snapshot isolation could address issue of read committed.  Reads that occur in the middle of a transaction read their data from the version of the data (the snapshot) that preceded the start of the transaction.  This makes it so that multi-object operations look atomic to the end user (assuming they succeed).


##### Implemetation   
  Using write locks and extended read value-holding (sometimes called "multiversion").  A key principle of snapshot isolation is **readers never block writers, and writers never block readers.**  
  <br/>
  * MVCC  
  	* [How MVCC work in postgresql](https://vladmihalcea.com/how-does-mvcc-multi-version-concurrency-control-work/)  
  	* [MVCC in Transactional Systems](https://0x0fff.com/mvcc-in-transactional-systems/)
  		* xmin - which defines the transaction id that inserted the record
  		* xmax - which defines the transaction id that deleted the row
  
		 <img src="resources/pictures/ddia_c7_postgres_mvcc.png" alt="ddia_c7_postgres_mvcc" width="400"/>   <br/>
			* When you insert data, you insert the row with xmin equal to current transaction id and xmax set to null;  
			* When you delete the data, you find visible row that should be deleted, and set its xmax to the current transaction id  
			* When you update the data, for each updated row you first perform “delete” and then “insert”.  


##### lost updates

Concurrent transactions that encapsulate read-modify-write operations will behave poorly on collision. A simple example is a counter that gets updated twice, but only goes up by one. The earlier write operation is said to be lost.  
  Ways to address this problem that live in the wild:
  - Atomic update operation (e.g. UPDATE keyword).
  - Transaction-wide write locking. Expensive!
  - Automatically detecting lost updates at the database level, and bubbling this back up to the application.
  - Atomic compare-and-set (e.g. UPDATE ... SET ... WHERE foo = 'expected_old_value').
  - Delayed application-based conflict resolution. Last resort, and only truly necessary for multi-master architectures.

##### write skew    

  <img src="resources/pictures/ddia_c7_writeskew_example.png" alt="ddia_c7_writeskew_example" width="600"/>   

  - As with lost updates, two transactions perform a read-modify-write, but now they modify two different objects based on the value they read.  
  - Example in the book: two doctors concurrently withdraw from being on-call, when business logic dictates that at least one must always be on call.  This occurs across multiple objects, so atomic operations do not help.  <br/>
  - Automatic detection at the snapshot isolation level and without serializability would require making consistency checks on every write, where is the number of concurrent write-carrying transactions in flight. This is way too high a performance penalty.<br/>
  - Only transaction-wide record locking works. So you have to make this transaction explicitly serialized, using e.g. a FOR UPDATE keyword.<br/>


##### hantom write skew.
  - Materializing conflicts
  - You can theoretically insert a lock on a phantom record, and then stop the second transaction by noting the presence of the lock. This is known as materializing conflicts.  This is ugly because it messes with the application data model, however. Has limited support.  If this issue cannot be mitigated some other way, just give up and go serialized.




### Serialization

#### Actual Serial Execution
- The most literal way is to run transactions on a single CPU in a single thread. This is actual serialization.
- This only became a real possible recently, with the speed-ups of CPUs and the increasing size of RAM. The bottleneck is obviously really low here. But it's looking tenable. Systems that use literal serialization are of a post-2007 vintage.  However, this requires writing application code in a very different way.  
- This method mplemeted in Redis([transactions-redis](https://hub.packtpub.com/transactions-redis/))  
<img src="resources/pictures/ddia_c7_redis_singlethread.png" alt="ddia_c7_redis_singlethread" width="600"/>  
<br/>

#### Pessimistic Lock/Optimistic Lock

Name | Details | Pros | Cons | Comments
---|:---|:---|:---|:---
Two-phase locking | * Transactions that read acquire shared-mode locks on touched records. <br/> * Transactions that write acquire, and transactions that want to write after reading update to, exclusive locks on touched records.<br/> * Transactions hold the highest grade locks they have for as long as the transaction is in play. | * Reads do not block reads or writes, but writes block everything.(Compre snapshot isolation, reads do not block reads or writes and writes do not block reads) | * Because so many locks occur, it's much easier than in snapshot isolation to arrive at a deadlock. Deadlocks occur when a transaction holds a lock on a record that another transaction needs, and that record holds a lock that the other transaction needs. The database has to automatically fail one of the transactions, roll it back, and prompt the application for a retry. <br/> * It has very bad performance implications. Long-blocking writes are a problem. Deadlocks are a disaster. Generally 2PL architectures have very bad performance at high percentiles; this is a main reason why "want-to-be sleek" systems like Amazon have moved away from them. | 
Serializable snapshot isolation | * based on snapshot isolation, but adds an additional algorithmic layer that makes it serialized and isolated. | * Beats 2PL for certain workloads especially read-heavy workloads <br/> * better performance when there is low record competition and when overall load is low| Poorer performance when lock competition is high and overall load is high. |



- 2 Phase Lock  
<img src="resources/pictures/ddia_c7_2pl_example.png" alt="ddia_c7_2pl_example" width="300"/>  
<br/>

- SSI ([Postgis Serializable](https://wiki.postgresql.org/wiki/Serializable))  

<img src="resources/pictures/postgresql-Serialization-Anomalies-in-Snapshot-Isolation.png" alt="postgresql-Serialization-Anomalies-in-Snapshot-Isolation" width="600"/>  
<br/>

- SSI example from book  

<img src="resources/pictures/ddia_c7_ssi_example.png" alt="ddia_c7_ssi_example" width="600"/>  
<br/>
事务42先提交，并且成功了。当事务43提交时，发现事务42已经提交了与自己相冲突的写入，所以必须中止事务43  

- How to avoid dead lock in 2PL(a separate thread checking)  

[CMU Concurrancy control](https://15721.courses.cs.cmu.edu/spring2017/slides/03-cc.pdf)  

- How 2PL addresses phantom skew(implementation detail)  
You can evade phantom skew by using predicate locks. These lock on all data in the database that matches a certain condition, even data that doesn't exist yet.  Predicate locks are expensive because they involve a lot of compare operations, one for each concurrent write. In practice most databases use index-range locks instead, which simplify the predicate to an index of values of some kind instead of a complex condition.  This lock covers a strict superset of the records covered by the predicate lock, so it causes more queries to have to wait for lock releases. But, it spares CPU cycles, which is currently worth it.  

- **SSI is an optimistic concurrency control technique. Two-phase locking is a pessimistic concurrency control technique.**  SSI works by allowing potentially conflicting operations to go through, then making sure that nothing bad happens; 2PL works by preventing potentially conflicting operations from occurring at all.
- SSI detects, at commit time (e.g. at the end of the transaction), whether or not any of the operations (reads or writes) that the transaction is performing are based on outdated premises (values that have changed since the beginning of the transaction). If not, the transaction goes through. If yes, then the transaction fails and a retry is prompted.


## Reference
- [postgresql transaction iso](https://www.postgresql.org/docs/9.5/transaction-iso.html)
- [postgresql high performance tips](https://vladmihalcea.com/9-postgresql-high-performance-performance-tips/)
- [A Critique of ANSI SQL Isolation Levels](https://blog.acolyer.org/2016/02/24/a-critique-of-ansi-sql-isolation-levels/)
- [InnoDB存储引擎MVCC实现原理](https://liuzhengyang.github.io/2017/04/18/innodb-mvcc/)
- [SQL Server Isolation Levels : A Series](https://sqlperformance.com/2014/07/t-sql-queries/isolation-levels)
- [mysql的可重复读隔离级别实现](https://www.jianshu.com/p/69fd2ca17cfd)