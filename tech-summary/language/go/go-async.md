# Golang Async

## What is Channel

unbuffered_channel

<img src="resource/unbuffered_channel.png" alt="unbuffered_channel" width="400"/>

<br/>

buffered_channel

<img src="resource/buffered_channel.png" alt="buffered_channel" width="400"/>
<br/>
Image from [The Nature Of Channels In Go](https://www.ardanlabs.com/blog/2014/02/the-nature-of-channels-in-go.html)


## Chan

- **When use channel, must be aware of who will create and when will close**
- Channels are both communicable and synchronize

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
