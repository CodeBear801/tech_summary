
# Golang Event Sourcing

[Gopherfest 2017: Event Sourcing – Architectures and Patterns (Matt Ho)](https://www.youtube.com/watch?v=B-reKkB8L5Q&t=562s)  
[git repo for demo](https://github.com/savaki/eventsource/blob/master/examples/dynamodb/main.go)   
[git repo for reference](https://github.com/altairsix/eventsource)



<img src="https://user-images.githubusercontent.com/16873751/103111386-529f1100-4601-11eb-84d4-06dc6bbccc62.png" alt="dynamodb2" width="600"/><br/>


Below are some good code reference of programming with golang  
## Option

Target  
```go
opts ...Option

// Option provides functional configuration for a *Repository
type Option func(*Repository)
// [perry] Apply a list of operations on Repository

```

Declaration
```go
type Option func(registry *repository)
```

Client code
```go
repo := eventsource.New(&User{}, eventsource.WithStore(store), /*other variables...*/)
```

Logic of `New`
```go
func New(tableName string, opts ...Option) (*Store, error) {

	for _, opt := range opts {
		opt(store)
	}
```

Impl
```go
func WithStore(store Store) Option {
	return func(registry *repository) {
		registry.store = store
	}
}

// WithDebug, WithObservers, WithSerializer
```


## How to define a interface mapping table
Target  
Interface Implementation Name -> Interface Impl Ptr  

client code
```go
// mapping interface name with its implementation ptr
	err := repo.Bind(
		UserCreated{},
		UserNameSet{},
		UserEmailSet{},
	)
```
Impl of Bind
```go

func (r *repository) Bind(events ...Event) error {
	for _, event := range events {
		if event == nil {
			return errors.New("attempt to bind nil event")
		}

		eventType, typ := EventType(event)
		r.logf("Binding %12s => %#v", eventType, event)
		r.types[eventType] = typ
	}

	return nil
}

// EventType is a helper func that extracts the event type of the event along with the reflect.Type of the event.
//
// Primarily useful for serializers that need to understand how marshal and unmarshal instances of Event to a []byte
func EventType(event Event) (string, reflect.Type) {
	t := reflect.TypeOf(event)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if v, ok := event.(EventTyper); ok {
		return v.EventType(), t
	}

	return t.Name(), t
}
```


Event interface
```go
type Event interface {
	// AggregateID returns the aggregate id of the event
	AggregateID() string

	// Version contains version number of aggregate
	EventVersion() int

	// At indicates when the event took place
	EventAt() time.Time
}
```

Interface Impl

```go
// UserCreated defines a user creation event
type UserCreated struct {
	eventsource.Model
}

// UserFirstSet defines an event by simple struct embedding
type UserNameSet struct {
	eventsource.Model
	Name string
}

// UserLastSet implements the eventsource.Event interface directly
type UserEmailSet struct {
	ID      string
	Version int
	At      time.Time
	Email   string
}

func (m UserEmailSet) AggregateID() string {
	return m.ID
}

func (m UserEmailSet) EventVersion() int {
	return m.Version
}

func (m UserEmailSet) EventAt() time.Time {
	return m.At
}
```


## slice as function parameter

```go

func (r *repository) Save(ctx context.Context, events ...Event) error {

	history := make(History, 0, len(events))
	for _, event := range events {
		

		history = append(history, record)
	}

	return r.store.Save(ctx, aggregateID, history...)
}
```
Here is code of calling `Save`

```go
err = repo.Save(ctx, setEmailEvent, setNameEvent)
```

## DynamoDB
- git repo https://github.com/savaki/eventsource/tree/master/provider/dynamodbstore
- https://docs.aws.amazon.com/zh_cn/amazondynamodb/latest/developerguide/DynamoDBLocal.UsageNotes.html
- https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GettingStarted.Java.html
- NoSQL https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/bp-general-nosql-design.html

```java
// create dynamo db's instance
        AmazonDynamoDB client = AmazonDynamoDBClientBuilder.standard()
            .withEndpointConfiguration(new AwsClientBuilder.EndpointConfiguration("http://localhost:8000", "us-west-2"))
            .build();

        DynamoDB dynamoDB = new DynamoDB(client);

```


```java
// create dynamo db's table
// year – The partition key. The ScalarAttributeType is N for number.
// title – The sort key. The ScalarAttributeType is S for string.

            Table table = dynamoDB.createTable(tableName,
            Arrays.asList(new KeySchemaElement("year", KeyType.HASH), // Partition key
                    new KeySchemaElement("title", KeyType.RANGE)), // Sort key
            Arrays.asList(new AttributeDefinition("year", ScalarAttributeType.N),
                         new AttributeDefinition("title", ScalarAttributeType.S)),
                new ProvisionedThroughput(10L, 10L));
            table.waitForActive();


```

```go
func makeUpdateItemInput(tableName, hashKey, rangeKey string, eventsPerItem int, aggregateID string, records ...eventsource.Record) ([]*dynamodb.UpdateItemInput, error) {
    // ...
	for partitionID, partition := range partitions {
		input := &dynamodb.UpdateItemInput{
			TableName: aws.String(tableName),
			Key: map[string]*dynamodb.AttributeValue{
				hashKey:  {S: aws.String(aggregateID)},
				rangeKey: {N: aws.String(strconv.Itoa(partitionID))},
			},
			ExpressionAttributeNames: map[string]*string{
				"#revision": aws.String("revision"),
			},
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":one": {N: aws.String("1")},
			},
        }
        

// dynamodb updateitem: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_UpdateItem.html
/*
partition key 分区键的意义就是将数据进行分区 another name is hash key, PK
官方推荐使用UUID作为分区键，使用UUID能最大限度地保证数据均匀分布在不同地分区中。

sort key为排序key, 用来组织数据, another name is range key, SK

Global Search Index(GSI) 是针对于特定查询而定制化的索引
https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GSI.html#GSI.scenario

from: https://zhuanlan.zhihu.com/p/101965292

DynamoDB 的数据集合是一张表，表里的每一条数据就是一个文档，文档中可以存储基本类型的数据(BOOLEAN, NUMBER, STRING)
或者一些复杂的数据(MAP, LIST, SET)，可以嵌套较为深的层级。在创建一张表的时候，需要指定一个基本类型的属性值作为 PK，
可以再指定一个基本类型的属性值作为 SK。PK 的作用是将数据分散到不同的分区(Partition)里，构建无序的哈希索引，SK的作用
是将同一个分区中的数据按照一定的顺序排列起来，便于查找。SK 中可以使用 ==, <, >=, <=, begins with, between, 
contains, in 等函数来进行较为丰富的查询操作，PK 必须使用准备的值进行查询。
from: https://www.jianshu.com/p/70ae83442584

*/

```
For here, `partition` is not dynamoDB's partition, its a logic to **truck** different version of same item together into dynamoDB's record.

<img src="https://user-images.githubusercontent.com/16873751/85053046-cb9f4680-b14e-11ea-8a1b-85457d1e9574.png" alt="dynamodb1" width="600"/><br/>

put a bounch of events into single dynamodb record.  

<img src="https://user-images.githubusercontent.com/16873751/85053064-d22dbe00-b14e-11ea-9b25-274feedb4a1d.png" alt="dynamodb2" width="600"/><br/>

## Temp notes



CommandHandler

```go
// Command encapsulates the data to mutate an aggregate
type Command interface {
	// AggregateID represents the id of the aggregate to apply to
	AggregateID() string
}

// CommandHandler consumes a command and emits Events
type CommandHandler interface {savaki, 4 years ago: • initial project commit
	// Apply applies a command to an aggregate to generate a new set of events
	Apply(ctx context.Context, command Command) ([]Event, error)
}
```

Aggregate

```go

// Aggregate represents the aggregate root in the domain driven design sense.
// It represents the current state of the domain object and can be thought of
// as a left fold over events.
type Aggregate interface {savaki, 4 years ago: • initial project commit
	// On will be called for each event; returns err if the event could not be
	// applied
	On(event Event) error
}

```

Store

```go

// Store provides an abstraction for the Repository to save data
type Store interface {
	// Save the provided serialized records to the store
	Save(ctx context.Context, aggregateID string, records ...Record) error

	// Load the history of events up to the version specified.
	// When toVersion is 0, all events will be loaded.
	// To start at the beginning, fromVersion should be set to 0
	Load(ctx context.Context, aggregateID string, fromVersion, toVersion int) (History, error)
}

```

Serializer

```go

// Serializer converts between Events and Records
type Serializer interface {
	// MarshalEvent converts an Event to a Record
	MarshalEvent(event Event) (Record, error)

	// UnmarshalEvent converts an Event backed into a Record
	UnmarshalEvent(record Record) (Event, error)
}

```



## keyword
DDD, Event Sourcing, Golang
