
## Go concurrency

### Deadlock with go channel

```
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
testing.(*T).Run(0xc4200c40f0, 0x113b276, 0x13, 0x11415d0, 0x1068246)
/usr/local/go/src/testing/testing.go:825 +0x301
testing.runTests.func1(0xc4200c4000)
/usr/local/go/src/testing/testing.go:1063 +0x64
testing.tRunner(0xc4200c4000, 0xc420057df8)
/usr/local/go/src/testing/testing.go:777 +0xd0
testing.runTests(0xc4200a6040, 0x11e3140, 0x1, 0x1, 0x100ea29)
/usr/local/go/src/testing/testing.go:1061 +0x2c4
testing.(*M).Run(0xc4200c0000, 0x0)
/usr/local/go/src/testing/testing.go:978 +0x171
main.main()
_testmain.go:42 +0x151

```
Correct: https://play.golang.org/p/tunBvfqxozY  
Error: https://play.golang.org/p/iMI24ihaJAm  
Fixed: https://play.golang.org/p/eUDS0--ecPv  
If the channel is unbuffered, the sender blocks until the receiver has received the value. If the channel has a buffer, the sender blocks only until the value has been copied to the buffer


### Dead lock with channel

Another issue about dead lock
	- About buffered channel & un buffered channel, before setting to buffered one, seems it hungs there forever
	- Pass waitgroup need pass by reference




## Go Library

### Swift compile issue

```
# github.com/Telenav/osrm-backend/traffic_updater/go/gen-go/proxy
../gen-go/proxy/proxy.go:396:14: too many arguments in call to oprot.Flush
have (context.Context)
want ()
../gen-go/proxy/proxy.go:413:16: too many arguments in call to oprot.Flush
have (context.Context)
want ()
../gen-go/proxy/proxy.go:426:16: too many arguments in call to oprot.Flush
have (context.Context)
want ()
../gen-go/proxy/proxy.go:440:24: too many arguments in call to oprot.Flush
have (context.Context)
want ()
../gen-go/proxy/proxy.go:461:16: too many arguments in call to oprot.Flush
have (context.Context)
want ()
../gen-go/proxy/proxy.go:474:16: too many arguments in call to oprot.Flush
have (context.Context)
want ()
../gen-go/proxy/proxy.go:488:24: too many arguments in call to oprot.Flush
have (context.Context)
want ()
../gen-go/proxy/proxy.go:509:16: too many arguments in call to oprot.Flush
have (context.Context)
want ()
../gen-go/proxy/proxy.go:522:16: too many arguments in call to oprot.Flush
have (context.Context)
want ()
../gen-go/proxy/proxy.go:536:24: too many arguments in call to oprot.Flush
have (context.Context)
want ()
../gen-go/proxy/proxy.go:536:24: too many errors

```
Swift changed its API, remove swift and re- go get
https://github.com/census-instrumentation/opencensus-go/issues/575

Solution: remove swift folder, reinstall swift 
