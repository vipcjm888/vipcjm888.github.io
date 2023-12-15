# 设计

#### 设计规范

好的设计需要注意“对齐、对比、重复、亲密性”四大设计原则，注意颜色运用与字体选择。

#写给大家的一本设计书

# 编程

## Golang

#### Service launch time

在线服务应避免在高峰期、节假日前上线，以免造成损失。

#### Line of sight

向下扫描看到一个缩进的某行，可以了解到这个函数做了什么；向下扫描到两个缩进的某行，可以了解这个函数的错误处理。这也是为什么if/else在一些golang代码中比较少去使用的可能原因？

```go
func (c *context) Render(code int, name string, data interface{}) (err error) {
	if c.echo.Renderer == nil {
		return ErrRendererNotRegistered
	}
	buf := new(bytes.Buffer)
	if err = c.echo.Renderer.Render(buf, name, data, c); err != nil {
		return
	}
	return c.HTMLBlob(code, buf.Bytes())
}

//bad
func doSomething() error {
    if something.OK() {
        return nil
    }else{
        return errors.New("something")
    }
}

//good
func doSomething() error {
    if !something.OK() {
        return errors.New("something")
    }
    return nil
}
```

#### Single method interfaces

使用仅有一个方法组成的接口，更简单，更易于实现，实现一个方法即可构建handler。常用于一些标准库。

```go
type Reader interface{
    Read(p []byte) (n int,err error)
}

type Handler interface{
    ServeHttp(ResponseWriter, *Request)
}

type HandlerFunc func(ResponseWriter, *Request)
```

#### Log Blocks

easier to see logs you care about

```go
func something() {
   log.Println("---------")
   defer log.Println("--------")
}
```

#### Return teardown function

return a teardown function from setup function

clean-up code is encapsulated

caller doesnt need to worry about cleaning up

if setup changes teardown code is with it

```go
func setup(t *testing.T)(*os.File,func(),error){
    teardonw:=func(){}
    // make a test file
    f,err:=ioutil.TempFile(os.TempDir,"test")
    if err!=nil {
        return nil,teardown,err
    }
    teardown = func(){
        // close f
        err := f.Close()
        if err != nil {
            t.Error("setup: Close:", err)
        }
        // delete the test file
        err = os.RemoveAll(f.Name())
        if err != nil {
            t.Error("setup: RemoveAll:", err)
        }
    }
    return f, teardown, nil
}

func TestSomething(t *testing.T){
    f,teardown,err:=setup(t)
    defer teardown()
    if err != nil {
        t.Error("setup:", err)
    }
    //do something with f
}
```

#### Discovering interfaces

添加一个类型`type Sizers []Sizer`,可以将多种对象视为一种

```go
type Sizer interface {
  Size() int64
}
```

can write an ad-hoc size calculator that makes use of the same methods, can just be the type itself

```go
type SizeFunc func() int64
func (s SizeFun) Size() int64 {
    return s()
}
type Size int64
func (s Size) Size() int64 {
    return int64(s)
}
```

#### Optional features

看某个对象是否存在某一个方法实现

function that calls options functions if that method exists. obj, ok := v.(Valid)

```go
type Valid interface {
    OK() error
}
func(p Person)OK()error{
    if p.Name=="" {
        return errors.New("name required")
    }
    return nil
}
func Decode(r io.Reader, v interface{})error {
    err:=json.NewDecoder(r).Decode(v)
    if err!=nil{
        return err
    }
    obj,ok:=v.(Valid)
    if !ok {
        return nil//no OK method
    }
    err = obj.OK()
    if err!=nil {
        return err
    }
    return nil
}
```

#### Retrying

from https://github.com/matryer/try

```go
package try

import "errors"
// MaxRetries is the maximum number of retries before bailing.
var MaxRetries = 10
var errMaxRetriesReached = errors.New("exceeded retry limit")
// Func represents functions that can be retried.
type Func func(attempt int) (retry bool, err error)
// Do keeps trying the function until the second argument
// returns false, or no error is returned.
func Do(fn Func) error {
	var err error
	var cont bool
	attempt := 1
	for {
		cont, err = fn(attempt)
		if !cont || err == nil {
			break
		}
		attempt++
		if attempt > MaxRetries {
			return errMaxRetriesReached
		}
	}
	return err
}
// IsMaxRetries checks whether the error is due to hitting the
// maximum number of retries or not.
func IsMaxRetries(err error) bool {
	return err == errMaxRetriesReached
}
```

#### Write obvious code not clever

```go
//bad
func some(){
    defer StartTimer("some")()
    //...
}
//good
func some(){
    stop := StartTimer("some")()
    defer stop()
    //...
}
```


#### Semaphores

limit number of go routines

```go
var (
    concurrent = 5
    semaphoreChan = make(chan struct{},concurrent)
)
func dowork(item int){
    semaphoreChan <- struct{}{}// block while full
    gofunc(){
        defer func(){
            <-semaphoreChan // read to release a slot
        }()
        log.Println(item)
        time.Sleep(1*time.Second)
    }()
}
func main(){
    for i:= 0; i< 10000;i++{
        doWork(i)
    }
}
```