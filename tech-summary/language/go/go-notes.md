
- [Go Notes](#Go-Notes)
  - [GOPATH](#GOPATH)
  - [Unit Test](#Unit-Test)
  - [Go package](#Go-package)
  - [Go profile](#Go-profile)
  - [GC](#GC)
  - [Channel management](#Channel-management)
  - [Grammar](#Grammar)
    - [string](#string)
    - [go flags](#go-flags)

# Go Notes

## GOPATH

[Settings of vscode](https://github.com/Microsoft/vscode-go/wiki/GOPATH-in-the-VS-Code-Go-extension)

```bash
"go.testEnvVars": {
"GOPATH": "/Users/xunliu/Desktop/git/distributesystem/6.824-golabs-2018"
},
"go.gopath": "/Users/xunliu/Desktop/git/distributesystem/6.824-golabs-2018",
```

On linux server we could simply set like below for temporary use
```bash
export GOPATH=xxx/go/
```
[<mark>Getting started with Go</mark>](https://medium.com/rungo/working-in-go-workspace-3b0576e0534a)  
[Setting up a Go Development Environment | Mac OS X](https://medium.com/@AkyunaAkish/setting-up-a-golang-development-environment-mac-os-x-d58e5a7ea24f)  
[Setup Go Development Environment with Visual Studio Code](https://rominirani.com/setup-go-development-environment-with-visual-studio-code-7ea5d643a51a)  


## Unit Test

Go test must use "_test" as file's suffix, otherwise might meet error of "no test files".   Function name should start with character with upper case.
For more information please reference to [How to write go code - test](https://golang.org/doc/code.html#Testing)
```
You write a test by creating a file with a name ending in _test.go that contains functions named TestXXX with signature func (t *testing.T). The test framework runs each such function; if the function calls a failure function such as t.Error or t.Fail, the test is considered to have failed.
```
If you want to run test for specific function, you could use *go test -run NameOfTest*, you could refer to here[How to run test cases in a specified file?](https://stackoverflow.com/questions/16935965/how-to-run-test-cases-in-a-specified-file)

More reference: https://golang.org/pkg/testing/, [how to test with go](https://www.calhoun.io/how-to-test-with-go/)



## Go package

My notes about [go package layout](./go-package-layout.md)

Package name should use no under_scores or mixedcaps
https://blog.golang.org/package-names

To get remote packages or dependencies, you could use *go get xxx*
https://golang.org/doc/code.html#remote
https://stackoverflow.com/questions/30295146/how-can-i-install-a-package-with-go-get

[Go包管理机制](https://io-meter.com/2014/07/30/go's-package-management/)

## Go profile

[<mark>Go profile example</mark>](https://flaviocopes.com/golang-profiling/)

[How to profile a program](https://golang.org/pkg/runtime/pprof/)


## GC

[How to clear a map in Go?](https://stackoverflow.com/questions/13812121/how-to-clear-a-map-in-go)
Basically, you could design a map with second level index, after use a set of data, you could assign map with make(map) which could tringle GC
  

## Channel management

[Go Concurrency Patterns: Pipelines and cancellation - goblog](https://blog.golang.org/pipelines)  
[<mark>Anatomy of Channels in Go - Concurrency in Go(2k)</mark>](https://medium.com/rungo/anatomy-of-channels-in-go-concurrency-in-go-1ec336086adb)  
[Testing of functions with channels in Go](https://www.sidorenko.io/post/2019/01/testing-of-functions-with-channels-in-go/)  



## Grammar

### string

Convert slice to string
```
strings.Trim(strings.Join(strings.Fields(fmt.Sprint(v.NodeIDs)), ","), "[]") 
```
More info: [One-liner to transform []int into string](https://stackoverflow.com/questions/37532255/one-liner-to-transform-int-into-string)


### go flags

Golang: parameter's sequence in command line should be the same with the one defined in the code, otherwise will use default value

```go
  var flags struct {
      port          int
      ip            string
      mappingFile   string
      csvFile       string
      highPrecision bool
  }    
       
  func init() {
      flag.IntVar(&flags.port, "p", 6666, "traffic proxy listening port")
      flag.StringVar(&flags.ip, "c", "127.0.0.1", "traffic proxy ip address")
      flag.StringVar(&flags.mappingFile, "m", "wayid2nodeids.csv", "OSRM way id to node ids mapping table")
      flag.StringVar(&flags.csvFile, "f", "traffic.csv", "OSRM traffic csv file")                                                                                                                         
      flag.BoolVar(&flags.highPrecision, "d", false, "use high precision speeds, i.e. decimal")
  }   
```

