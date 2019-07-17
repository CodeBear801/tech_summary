

##GOPATH

Settings of vscode
https://github.com/Microsoft/vscode-go/wiki/GOPATH-in-the-VS-Code-Go-extension

"go.testEnvVars": {
"GOPATH": "/Users/xunliu/Desktop/git/distributesystem/6.824-golabs-2018"
},
"go.gopath": "/Users/xunliu/Desktop/git/distributesystem/6.824-golabs-2018",


Export GOPATH=xxx/go/


https://medium.com/@AkyunaAkish/setting-up-a-golang-development-environment-mac-os-x-d58e5a7ea24f  
https://rominirani.com/setup-go-development-environment-with-visual-studio-code-7ea5d643a51a  
<br/>
Setting go env +++ 
https://medium.com/rungo/working-in-go-workspace-3b0576e0534a

## Unit Test

Go test must use "_test" as file's suffix, otherwise might meet error of "no test files".   Function name should start with character with upper case.
For more information please reference to [How to write go code - test](https://golang.org/doc/code.html#Testing)
```
You write a test by creating a file with a name ending in _test.go that contains functions named TestXXX with signature func (t *testing.T). The test framework runs each such function; if the function calls a failure function such as t.Error or t.Fail, the test is considered to have failed.
```
If you want to run test for specific function, you could use *go test -run NameOfTest*, you could refer to here[How to run test cases in a specified file?](https://stackoverflow.com/questions/16935965/how-to-run-test-cases-in-a-specified-file)

More reference: https://golang.org/pkg/testing/, [how to test with go](https://www.calhoun.io/how-to-test-with-go/).



## Go package
Package name should use no under_scores or mixedcaps
https://blog.golang.org/package-names

To get remote packages or dependencies, you could use *go get xxx*
https://golang.org/doc/code.html#remote
https://stackoverflow.com/questions/30295146/how-can-i-install-a-package-with-go-get


## GC

How to clear a map in Go?
https://stackoverflow.com/questions/13812121/how-to-clear-a-map-in-go
Basically, you could design a map with second level index, after use a set of data, you could assign map with make(map) which could tringle GC
  

##Grammar

### string

Convert slice to string
```
strings.Trim(strings.Join(strings.Fields(fmt.Sprint(v.NodeIDs)), ","), "[]") 
```


### go flags

Golang: parameter's sequence in command line should be the same with the one defined in the code, otherwise will use default value

```
 13 var flags struct {
 14     port          int
 15     ip            string
 16     mappingFile   string
 17     csvFile       string
 18     highPrecision bool
 19 }    
 20      
 21 func init() {
 22     flag.IntVar(&flags.port, "p", 6666, "traffic proxy listening port")
 23     flag.StringVar(&flags.ip, "c", "127.0.0.1", "traffic proxy ip address")
 24     flag.StringVar(&flags.mappingFile, "m", "wayid2nodeids.csv", "OSRM way id to node ids mapping table")
 25     flag.StringVar(&flags.csvFile, "f", "traffic.csv", "OSRM traffic csv file")                                                                                                                         
 26     flag.BoolVar(&flags.highPrecision, "d", false, "use high precision speeds, i.e. decimal")
 27 }   
```


## Reference Links

Golang 
Anatomy of Channels in Go - Concurrency in Go(2k)
https://medium.com/rungo/anatomy-of-channels-in-go-concurrency-in-go-1ec336086adb

concurrency for beginners(wait-group)
https://medium.com/@matryer/very-basic-concurrency-for-beginners-in-go-663e63c6ba07

Channel test
https://www.sidorenko.io/post/2019/01/testing-of-functions-with-channels-in-go/

Merge channel
https://medium.com/justforfunc/two-ways-of-merging-n-channels-in-go-43c0b57cd1de

Functional programming in go(1.5K)
https://medium.com/@geisonfgfg/functional-go-bc116f4c96a4


How to profile a program
https://golang.org/pkg/runtime/pprof/

Go profile+++
https://flaviocopes.com/golang-profiling/

go依赖包管理
https://io-meter.com/2014/07/30/go's-package-management/


Golang, covert vector of int to string
https://stackoverflow.com/questions/37532255/one-liner-to-transform-int-into-string


Golang work with large csv
https://www.reddit.com/r/golang/comments/9qyvlo/golang_for_etl_data_pipelines_like_airflow_luigi/


Golang for ETL / Data pipelines (like Airflow, Luigi)?
https://www.reddit.com/r/golang/comments/9qyvlo/golang_for_etl_data_pipelines_like_airflow_luigi/
Could use bash operators in airflowhttps://github.com/apache/airflow/blob/f5e3b03aa6b902898148e88fb9d90ecebcadc226/docs/howto/operator.rst#bashoperator
Or https://github.com/jondot/crunch

Go Concurrency Patterns: Pipelines and cancellation
https://blog.golang.org/pipelines

Reading a file concurrently in Golang
https://stackoverflow.com/questions/27217428/reading-a-file-concurrently-in-golang


Tips to working with large CSV files?  -> uses a bufio.Reader
https://www.reddit.com/r/golang/comments/9wo7bd/tips_to_working_with_large_csv_files/


