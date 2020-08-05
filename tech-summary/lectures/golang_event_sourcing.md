# Golang Event Sourcing

[Gopherfest 2017: Event Sourcing – Architectures and Patterns (Matt Ho)](https://www.youtube.com/watch?v=B-reKkB8L5Q&t=562s)  
[git repo for demo](https://github.com/savaki/eventsource/blob/master/examples/dynamodb/main.go)   
[git repo for reference](https://github.com/altairsix/eventsource)

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
```
For here, `partition` is not dynamoDB's partition, its a logic to **truck** different version of same item together into dynamoDB's record.

<img src="https://user-images.githubusercontent.com/16873751/85053046-cb9f4680-b14e-11ea-8a1b-85457d1e9574.png" alt="dynamodb1" width="400"/><br/>
<img src="https://user-images.githubusercontent.com/16873751/85053064-d22dbe00-b14e-11ea-9b25-274feedb4a1d.png" alt="dynamodb2" width="400"/><br/>


## keyword
DDD, Event Sourcing, Golang
