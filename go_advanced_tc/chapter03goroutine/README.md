# chapter3-并发编程
## 作业内容
1. 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

## 思考分析

### 代码文件与执行方式

本目录下的homework.go

```shell
go mod tidy
go mod vendor
go run homework.go
```

### 代码分析

#### httpserver的执行

```go
type httpHandler struct {
	wg       sync.WaitGroup // 等价于request group
	shutdown bool           // 判断服务是否正常
}
```

结构体`httpHandler`表示http请求的处理器，`wg`用来等待http请求处理，`shutdown`用来标记server是否处于`关闭中`的状态

```go
func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.shutdown {
		fmt.Fprintln(w, "internal server error")
		fmt.Println("internal server error")
		return
	}
	h.wg.Add(1) // 当有一个请求来时，Add(1)

	sleep := rand.Intn(20)
	fmt.Fprintln(w, "Hello World", sleep)
	fmt.Println("Hello World", sleep)
	go h.Event("Event"+strconv.Itoa(sleep), sleep) // 在这里模拟业务的异步处理
}

func (h *httpHandler) Event(data string, sleep int) {
	time.Sleep(time.Second * time.Duration(sleep))
	fmt.Println(data)
	h.wg.Done() // 当有一个请求处理完成时，Done()
}
```

`httpHandler`实现了`ServeHTTP()`方法，接收并处理http请求：
1. 首先判断`shutdown`，为`true`时停止处理http请求
2. `shutdown`为`false`时，调用`wg.Add(1)`，表示接收1个http请求并开始处理
3. 通过`go h.Event()`来模拟异步的业务处理

`httpHandler`实现了`Event()`方法，模拟异步业务处理：随机sleep一段时间，并调用`wg.Done()`表示该请求处理完成

```go
func (h *httpHandler) BeforeShutdown() {
	// 设置shutdown，不再处理新的http请求
	h.shutdown = true

	// 设置shutdown的超时时间
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	// 等待所有请求处理完成
	ch := make(chan struct{}) // 作用等同于main中的channel(stop)
	go func() {
		h.wg.Wait() // 在Shutdown之前等待所有请求处理完成
		close(ch)
	}()

	select {
	case <-ch:
		fmt.Println("all request done")
	case <-ctx.Done():
		fmt.Println("handler request timeout")
	}
}
```
`httpHandler`实现了`BeforeShutdown()`方法，处理server关闭前的动作：
1. 设置`shutdown`为`true`，使handler停止处理新进的http请求
2. 调用`wg.Wait()`等待正在处理的http请求处理完成
3. 通过`context.WithTimeout()`设置shutdown的超时时间，防止正在处理的http请求处理时间过长导致server长时间无法关闭

#### httpserver的创建

```go
func server(addr string, stop <-chan struct{}, done chan<- error) error {
	h := httpHandler{wg: sync.WaitGroup{}}
	s := http.Server{
		Addr:    addr,
		Handler: &h,
	}

	go func() {
		<-stop // goroutine会在此阻塞，直到stop内有消息，或stop被关闭
		fmt.Println("receive a stop signal or channel(stop) is closed")

		h.BeforeShutdown()

		s.Shutdown(context.Background()) // 在此主动关闭server，该goroutine会结束
	}()

	err := s.ListenAndServe()
	done <- err
	return err
}
```

`server()`提供启动一个httpserver的方法，其入参包含`stop <-chan struct{}`表示server主动关闭信号，`done chan<- error`表示server异常信号

在`s.ListenAndServe()`前启动一个goroutine，监听`stop`channel，当收到信号时表示需要主动关闭当前server

当`s.ListenAndServe()`产生error时（`s.Shutdown()`的调用或其他场景产生的error），将该error推入`done`channel，表示当前server已经由于某种原因关闭

#### httpserver的关闭

```go
func stopListen(stop chan<- struct{}, done <-chan error) {
	select {
	case <-done:
		close(stop)
	}
}
```

该方法监听`done`channel，当收到信号时表示有server出于某种状态已经关闭，此时将关闭`stop`channel，以通知其他server主动关闭

#### Linux signal的监听

```go
func signalListen(done chan<- error) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	select {
	case s := <-c:
		fmt.Println("get signal:", s)
	}
	signal.Stop(c)
	err := fmt.Errorf("get signal")
	done <- err
}
```

通过`signal.Notify()`监听interrupt和terminate信号

当收到两者任一信号时，首先停止信号监听，然后向`done`channel推入一个信号，使`stopListen`捕捉到该信号

#### 程序入口

```go
func main() {
	stop := make(chan struct{})
	done := make(chan error, 4)
	go signalListen(done)
	go stopListen(stop, done)
	eg := &errgroup.Group{}
	eg.Go(func() error {
		return server(":8888", stop, done)
	})
	eg.Go(func() error {
		return server(":8889", stop, done)
	})
	eg.Go(func() error {
		return serverFork(done)
	})
	err := eg.Wait()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
```

1. 创建`stop`和`done`两个channel
2. 通过goroutine启动linux信号监听和server关闭信号监听
3. 通过errgroup在后台启动httpserver
4. 等待所有httpserver关闭

#### fork error

```go
// 模拟一个server，最终返回一个error
func serverFork(done chan<- error) error {
	time.Sleep(forkErrorTimeout)
	err := fmt.Errorf("fork err")
	done <- err
	return err
}
```

为方便调试，`serverFork`用来模拟一个httpserver，但只会在sleep一段时间后向`done`channel发送一个信号