
# Golang Package layout

[Standard package layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)

## Recommended

- Root package is for domain types
- Group subpackages by **dependency**
- **Use a shared mock subpackage**
- Main package ties together dependencies

### Root package is for domain types
Root package defines the domain of entire project, and not depends on other packages in the project
```go
package myapp

type User struct {
	ID      int
	Name    string
	Address Address
}

type UserService interface {
	User(id int) (*User, error)
	Users() ([]*User, error)
	CreateUser(u *User) error
	DeleteUser(id int) error
}

```
<mark>The root package should not depend on any other package in your application!</mark>


### Group subpackages by dependency
If UserService is depend on posgresql as storage layer, there should be a sub package named postgres and providing postgres.UserService implementation.  Most importantly, when the dependencies reference to UserService, they will use a meaningful name of postgres.UserService
```go
package postgres

import (
	"database/sql"

	"github.com/benbjohnson/myapp"
	_ "github.com/lib/pq"
)

// UserService represents a PostgreSQL implementation of myapp.UserService.
type UserService struct {
	DB *sql.DB
}

// User returns a user for a given id.
func (s *UserService) User(id int) (*myapp.User, error) {
	var u myapp.User
	row := db.QueryRow(`SELECT id, name FROM users WHERE id = $1`, id)
	if row.Scan(&u.ID, &u.Name); err != nil {
		return nil, err
	}
	return &u, nil
}

// implement remaining myapp.UserService interface...

```

Benefits:
- Later you could use another database to record data, such as boltdb

- Easy to add a new layer, such as cache for postgresql

```go
package myapp

// UserCache wraps a UserService to provide an in-memory cache.
type UserCache struct {
        cache   map[int]*User
        service UserService
}

// NewUserCache returns a new read-through cache for service.
func NewUserCache(service UserService) *UserCache {
        return &UserCache{
                cache: make(map[int]*User),
                service: service,
        }
}

// User returns a user for a given id.
// Returns the cached instance if available.
func (c *UserCache) User(id int) (*User, error) {
	// Check the local cache first.
        if u := c.cache[id]]; u != nil {
                return u, nil
        }

	// Otherwise fetch from the underlying service.
        u, err := c.service.User(id)
        if err != nil {
        	return nil, err
        } else if u != nil {
        	c.cache[id] = u
        }
        return u, err
}

```
We see this approach in the standard library too. The io.Reader is a domain type for reading bytes and its implementations are grouped by dependency — tar.Reader, gzip.Reader, multipart.Reader. These can be layered as well. It’s common to see an os.File wrapped by a bufio.Reader which is wrapped by a gzip.Reader which is wrapped by a tar.Reader.


- <mark>Separate external dependency with domain logic.</mark>.  
  External dependency likes database or 3rd party email service, domain logic such as user authentication, request data verification, external service calling and trigger email sending.   
  There should be package purely handling domain logic, such as server.UserService;  and we could abstract database related logic into server.Repository.  

<img src="resource/golang-package-layout.png" alt="golang-package-layout" width="400"/>

- The red part is API protocol with external 
- The blue part is domain logic.  They don't contains the implementation of interface, just use them
- The green part is the implementation of those interfaces

```bash
.
├── cmd/
│   └── myapp/
│       └── main.go
├── server/
│   ├── repository/
│   │   ├── mongodb/
│   │   │   ├── repository.go
│   │   │   └── repository_test.go
│   │   ├── postgres/
│   │   │   ├── repository.go
│   │   │   └── repository_test.go
│   │   └── repositorytest/
│   │       └── tester.go
│   ├── service.go
│   └── service_test.go
├── transport/
│   ├── grpc/
│   │   ├── transport.go
│   │   └── transport_test.go
│   └── restful/
│       ├── transport.go
│       └── transport_test.go
└── service.go

```




### Use a shared mock subpackage
Write mock files to support your unit test
```go
package mock

import "github.com/benbjohnson/myapp"

// UserService represents a mock implementation of myapp.UserService.
type UserService struct {
        UserFn      func(id int) (*myapp.User, error)
        UserInvoked bool

        UsersFn     func() ([]*myapp.User, error)
        UsersInvoked bool

        // additional function implementations...
}

// User invokes the mock implementation and marks the function as invoked.
func (s *UserService) User(id int) (*myapp.User, error) {
        s.UserInvoked = true
        return s.UserFn(id)
}

// additional functions: Users(), CreateUser(), DeleteUser()

```

And here is the test

```go
package http_test

import (
	"testing"
	"net/http"
	"net/http/httptest"

	"github.com/benbjohnson/myapp/mock"
)

func TestHandler(t *testing.T) {
	// Inject our mock into our handler.
	var us mock.UserService
	var h Handler
	h.UserService = &us

	// Mock our User() call.
	us.UserFn = func(id int) (*myapp.User, error) {
		if id != 100 {
			t.Fatalf("unexpected id: %d", id)
		}
		return &myapp.User{ID: 100, Name: "susy"}, nil
	}

	// Invoke the handler.
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/users/100", nil)
	h.ServeHTTP(w, r)
	
	// Validate mock.
	if !us.UserInvoked {
		t.Fatal("expected User() to be invoked")
	}
}

```


### Main package ties together dependencies
Use the Go convention of placing our main package as a subdirectory of the cmd package.

```bash
myapp/
	cmd/
		myapp/
			main.go
		myappctl/
			main.go
```

Injecting dependencies at compile time

```go
package main

import (
	"log"
	"os"
	
	"github.com/benbjohnson/myapp"
	"github.com/benbjohnson/myapp/postgres"
	"github.com/benbjohnson/myapp/http"
)

func main() {
	// Connect to database.
	db, err := postgres.Open(os.Getenv("DB"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create services.
	us := &postgres.UserService{DB: db}

	// Attach to HTTP handler.
	var h http.Handler
	h.UserService = us
	
	// start http server...
}

```


## MVC(Beego)
Example project is [Beego](https://github.com/astaxie/beego)

```bash
├── conf
│   └── app.conf
├── controllers
│   ├── admin
│   └── default.go
├── main.go
├── models
│   └── models.go
├── static
│   ├── css
│   ├── ico
│   ├── img
│   └── js
└── views
    ├── admin
    └── index.tpl
```

Example
[controllers](https://github.com/beego/samples/blob/master/todo/controllers/task.go) could only visit public interface in [model](https://github.com/beego/samples/blob/master/todo/models/task.go) layer

If you use microservice and in your code just have few model, you could use such strategy to orgnize your code.


## 知乎

[知乎社区核心业务 Golang 化实践](https://zhuanlan.zhihu.com/p/48039838?hmsr=toutiao.io&utm_medium=toutiao.io&utm_source=toutiao.io)

```bash
.
├── bin              	--> 构建生成的可执行文件
├── cmd              	--> 各种服务的 main 函数入口（ RPC、Web 等）
│   ├── service 
│   │    └── main.go
│   ├── web
│   └── worker
├── gen-go           	--> 根据 RPC thrift 接口自动生成
├── pkg              	--> 真正的实现部分（下面详细介绍）
│   ├── controller
│   ├── dao
│   ├── rpc
│   ├── service
│   └── web
│   	├── controller
│   	├── handler
│   	├── model
│   	└── router
├── thrift_files     	--> thrift 接口定义
│   └── interface.thrift
├── vendor           	--> 依赖的第三方库（ dep ensure 自动拉取）
├── Gopkg.lock       	--> 第三方依赖版本控制
├── Gopkg.toml
├── joker.yml        	--> 应用构建配置
├── Makefile         	--> 本项目下常用的构建命令
└── README.md
```
- bin：构建生成的可执行文件，一般线上启动就是 `bin/xxxx-service`
- cmd：各种服务（RPC、Web、离线任务等）的 main 函数入口，一般从这里开始执行
- gen-go：thrift 编译自动生成的代码，一般会配置 Makefile，直接 `make thrift` 即可生成（这种方式有一个弊端：很难升级 thrift 版本）
- pkg：真正的业务实现（下面详细介绍）
- thrift_files：定义 RPC 接口协议
- vendor：依赖的第三方库

pkg 下放置着项目的真正逻辑实现

```bash
pkg/
├── controller    	
│   ├── ctl.go       	--> 接口
│   ├── impl         	--> 接口的业务实现
│   │	└── ctl.go
│   └── mock         	--> 接口的 mock 实现
│     	└── mock_ctl.go
├── dao           	
│   ├── impl
│   └── mock
├── rpc           	
│   ├── impl
│   └── mock
├── service       	--> 本项目 RPC 服务接口入口
│   ├── impl
│   └── mock
└── web           	--> Web 层（提供 HTTP 服务）
    ├── controller    	--> Web 层 controller 逻辑
    │   ├── impl
    │   └── mock
    ├── handler       	--> 各种 HTTP 接口实现
    ├── model         	-->
    ├── formatter     	--> 把 model 转换成输出给外部的格式
    └── router        	--> 路由
```





## Reference
- [Golang UK Conference 2017 | Mat Ryer - Writing Beautiful Packages in Go](https://www.youtube.com/watch?v=cAWlv2SeQus&t=18s)

