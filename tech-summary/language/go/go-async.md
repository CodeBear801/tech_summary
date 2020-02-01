
# Golang Async

## What is Channel

unbuffered_channel

<img src="resource/unbuffered_channel.png" alt="unbuffered_channel" width="400"/>  


buffered_channel

<img src="resource/buffered_channel.png" alt="buffered_channel" width="400"/>  

Image from [The Nature Of Channels In Go](https://www.ardanlabs.com/blog/2014/02/the-nature-of-channels-in-go.html)


## Chan

- **When use channel, must be aware of who will create and when will close**
- Channels are both communicable and synchronize


### Nil chan

Nil channel: nil channel always blocks, which means that trying to read or write from a nil channel will block.  
**You could always read from closed channel, but you can never write to closed channel.**

### Chan of Chan

When using a `chan chan`, the thing you want to send down the transport is another transport to send things back.  
They are useful when you want to get a response to something, and you don’t want to setup two channels (it’s generally considered bad practice to have data moving bidirectionally on a single channel)  

<img src="resource/chan-chan.png" alt="chan-chan" width="400"/>  

Image from [Understanding Chan Chan's in Go](http://tleyden.github.io/blog/2013/11/23/understanding-chan-chans-in-go/)

[example](https://play.golang.org/p/chi6P2XGTO)

### Return chan

```go
func gen(nums []int) <-chan int {
}
```

[example](https://play.golang.org/p/Qh30wzo4m0)

## Select

- select likes switch statement for multiple channels.  
   + All channels are evaluated
   + Selection blocks until one communication can proceed, which then does
   + If multiple can proceed, select chooses pseudo-randomly
   + A default clause, if present, executes immediately if no channel is ready


### for-select pattern
- create a loop runs in its own goroutine
- select lets loop avoid blocking indefinitely in any state

```go
func (s * sub) loop() {
    ... declare mutable state...
    select {
        case <- c1:
            ... read/write state ...
        case c2 <- x:
            ... read/write state ...
        case y := <-timer:
            ... read/write state ...
    }
}

```
- The cases interact via local state in loop


### Stop goroutine
- A goroutine could be stopped by timer/notify by waitgroup/channel

```go
func timeout(w *sync.WaitGroup, t time.Duration) bool {
	temp := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		defer close(temp)
		w.Wait()
	}()

	select {
	case <-temp:
		return false
	case <-time.After(t):
		return true
	}
}
```

## Fan-in
Multiplexing (fan-in): function that takes multiple channels and pipes to one channel, so that the returned channel receives both outputs

<img src="resource/fan-in.png" alt="fan-in" width="400"/> 

## Pipeline

- [go-osmium](https://github.com/Telenav/open-source-spec/blob/master/osmium/doc/pbf_golang.md) 
- [OSRM speed table dumper](https://github.com/Telenav/osrm-backend/blob/037a545659cef519f976360ae5c90dffcbafc145/integration/cmd/osrm-traffic-updater/speed_table_dumper.go#L26)


## Great examples

### LoadBalancer

[code](https://github.com/CodeBear801/tech_summary/blob/master/tech-summary/language/go/code/sync/balance.go#L86) [slides](https://talks.golang.org/2012/waza.slide#41) [video](https://www.youtube.com/watch?v=jgVhBThJdXc)

<img src="resource/balancer.png" alt="balancer" width="400"/> 

The code simulate how to write a load balancer  

Input
- Simulate 100 goroutine and each goroutine keep on sending requests
- each goroutine execute the function of `requester`, for { send request to work, get response}
- `work` is a unbuffered channel, means there is just one request get in at one time
- `request`'s definition contains fn() and a channel for response

```go
// [Perry] requester simulate the call from external users
func requester(work chan Request) {
	c := make(chan int)
	for {
		time.Sleep(time.Duration(rand.Int63n(nWorker * 2e9)))
		// [Perry] send request
		work <- Request{op, c}
		// [Perry] wait for answer, result:= <-c, then further process on result
		<-c
	}
}
```

Worker
- worker contains a buffered channel for request
- worker contains `pending` which means how many requests on him
- worker has an index in pool/heap
- worker function keep on executing requests, send response to channel, and notify loadbalancer

```go
func (w *Worker) work(done chan *Worker) {
	for {
		req := <-w.requests  // get request from load balancer
		req.c <- req.fn()    // call fn and send result to requester
		done <- w            // tell balancer we finished the job, send worker's pointer 
	}
}

```

Balancer
- Most important task for balancer is picking who is next worker, by balance()
- When there is a work(request), balancer will call `dispatch()` then pick a worker, put request in his queue
- When there is a worker done his task(via done), balancer will call `complete()` to adjust status
- Use `worker`'s pointer as trigger is interesting

```go

func (b *Balancer) balance(work chan Request) {
	for {
		select {
		case req := <-work:
			b.dispatch(req)
		case w := <-b.done:
			b.completed(w)
		}
		b.print()
	}
}
```

### Google search Example

[code]() [video](https://www.youtube.com/watch?v=f6kdp27TYZs&t=1021s)
