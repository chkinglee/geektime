// Package main
// @Author      : lilinzhen
// @Time        : 2022/2/13 18:59:05
// @Description :
package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

const (
	shutdownTimeout  = time.Second * 5
	forkErrorTimeout = time.Second * 10
)

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

func stopListen(stop chan<- struct{}, done <-chan error) {
	select {
	case <-done:
		close(stop)
	}
}

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

// 模拟一个server，最终返回一个error
func serverFork(done chan<- error) error {
	time.Sleep(forkErrorTimeout)
	err := fmt.Errorf("fork err")
	done <- err
	return err
}

type httpHandler struct {
	wg       sync.WaitGroup // 等价于request group
	shutdown bool           // 判断服务是否正常
}

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
