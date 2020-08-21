
## Problem to solve
<img src="https://user-images.githubusercontent.com/16873751/90936435-04ab9f80-e3ba-11ea-904b-a19955d148db.png" alt="main_operational_flow" width="600"/>

<img src="https://user-images.githubusercontent.com/16873751/90936196-9666dd00-e3b9-11ea-8717-83ec48687bb8.png" alt="story_1" width="600"/>  
<img src="https://user-images.githubusercontent.com/16873751/90936206-9a92fa80-e3b9-11ea-889a-c13acaf544b1.png" alt="story_2" width="600"/>

Explanation

- User could make an `order`, which contains order ref and several `order line`, each `order line` contains `sku`(stock keeping unit) and quantity.  For example, you could make an order contains 10 unit of red-chair and 1 unit of lamp.
- `order line` is allocated on `batches`(inventory, stock).  A `batch` contains `unique ID`, `sku` and `quantity`.  `Purchasing department` is responsible for manage `batches`

Special cases
- When you sell items you need to decrease quantity on batches
- You can't sell more items than existing quantity
- Same order must not be processed for multiple times(refresh page, server down)
- When `batches` running on low quantity should be a notice to purchase more


## Code analysis
- [entrance layer](https://github.com/CodeBear801/pythonArchtectureCode/blob/e3de63c922c2094d63949949b4cdf776fb5462dc/src/allocation/entrypoints/flask_app.py#L12)
  + A wrapper could be implemented for support entrance testing [code](https://github.com/CodeBear801/pythonArchtectureCode/blob/e3de63c922c2094d63949949b4cdf776fb5462dc/tests/e2e/api_client.py#L5)

- [`bootstrap`](https://github.com/CodeBear801/pythonArchtectureCode/blob/e3de63c922c2094d63949949b4cdf776fb5462dc/src/allocation/bootstrap.py#L10) worked as builder or resource manager.  It allocates/constructs all things needed by `messagebus`

- With the help of `sqlachemy`, it could directly operate on database object
   + [mapper()](https://github.com/CodeBear801/pythonArchtectureCode/blob/e3de63c922c2094d63949949b4cdf776fb5462dc/src/allocation/adapters/orm.py#L52) be called at beginning 
 

### A trip about event trigger and handler
Take `allocate` as an example
- The initial entrance is [allocation/entrypoints/flask_app.py](https://github.com/CodeBear801/pythonArchtectureCode/blob/e3de63c922c2094d63949949b4cdf776fb5462dc/src/allocation/entrypoints/flask_app.py#L24)

```py
@app.route("/allocate", methods=['POST'])
def allocate_endpoint():
    try:
        # generates a command and calling evnetbus to handle this command
        cmd = commands.Allocate(
            request.json['orderid'], request.json['sku'], request.json['qty'],
        )
        bus.handle(cmd)
```
- [handler()](https://github.com/CodeBear801/pythonArchtectureCode/blob/e3de63c922c2094d63949949b4cdf776fb5462dc/src/allocation/service_layer/messagebus.py#L27) in messagebus

```py
    def handle(self, message: Message):
        self.queue = [message]
        while self.queue:
            message = self.queue.pop(0)
            if isinstance(message, events.Event):
                self.handle_event(message)
            elif isinstance(message, commands.Command):
                self.handle_command(message)
            else:
                raise Exception(f'{message} was not an Event or Command')

    def handle_command(self, command: commands.Command):
        logger.debug('handling command %s', command)
        try:
            # get registered function 
            handler = self.command_handlers[type(command)]
            handler(command)
            # handler could generate new events, add to queue
            self.queue.extend(self.uow.collect_new_events())
        except Exception:
            logger.exception('Exception handling command %s', command)
            raise
```
-  Here is the logic of [`handler`](https://github.com/CodeBear801/pythonArchtectureCode/blob/e3de63c922c2094d63949949b4cdf776fb5462dc/src/allocation/service_layer/handlers.py#L31)
```py
def allocate(
        cmd: commands.Allocate, uow: unit_of_work.AbstractUnitOfWork
):
    line = OrderLine(cmd.orderid, cmd.sku, cmd.qty)
    with uow:
        product = uow.products.get(sku=line.sku)
        # products' type is repository.AbstractRepository
        # get function will trigger db to load specific product
        # For example, is using SqlAlchemyRepository, self.session.query(model.Product).filter_by(sku=sku).first()

        # product's type is model.Product

        if product is None:
            raise InvalidSku(f'Invalid sku {line.sku}')
        product.allocate(line)
        uow.commit()
```

- So, the [domain model](https://github.com/CodeBear801/pythonArchtectureCode/blob/e3de63c922c2094d63949949b4cdf776fb5462dc/src/allocation/domain/model.py#L17) only need to focus on business logic
```py
class Product:

    def allocate(self, line: OrderLine) -> str:
        try:
            batch = next(
                b for b in sorted(self.batches) if b.can_allocate(line)
            )
            batch.allocate(line)
            self.version_number += 1
            # new events will go to message bus and then trigger related handler
            self.events.append(events.Allocated(
                orderid=line.orderid, sku=line.sku, qty=line.qty,
                batchref=batch.reference,
            ))
            return batch.reference
        except StopIteration:
            self.events.append(events.OutOfStock(line.sku))
            return None

        # self.events = []  # type: List[events.Event]
```
Observer for `Allocated`
```py
EVENT_HANDLERS = {
    events.Allocated: [publish_allocated_event, add_allocation_to_read_model],
}
# update to redis.pubsub
# update view

def publish_allocated_event(
        event: events.Allocated, publish: Callable,
):
    #     publish: Callable = redis_eventpublisher.publish,
    publish('line_allocated', event)
```


## Questions

How to generate unique order ID
- UserID + time + device ID
- distribute lock


why Onion architecture
- Depends on interface, not implementation
- app -> domain <- database

What's the purpose of `replication`
- Usually, it supports get()/set()/list()
- `repo pattern`, abstraction over persist storage, hiding details and pretending all data in memory

What's the difference between `value object` and `entity`
- immutable
- unique identifier


## Reference
- book: https://github.com/cosmicpython/book
- code: https://github.com/cosmicpython/code





