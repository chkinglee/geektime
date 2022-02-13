// Package main
// @Author      : lilinzhen
// @Time        : 2022/2/13 16:16:26
// @Description :
package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func main() {
	/*
		在case_5中，使用了sync.WaitGroup来保证server在关闭前处理完所有正在处理的请求，但可能因请求源源不断或请求长时间处理而导致server永远无法退出

		在本示例中
			1、使用context.WithTimeout来使server在一定时间内退出
			2、设定一个bool值，当server收到stop信号时，在正式关闭前，停止新请求的处理
	*/
	done := make(chan error, 3)
	stop := make(chan struct{})
	go func() {
		done <- server6(":8888", stop)
	}()
	go func() {
		done <- server6(":8889", stop)
	}()
	go func() {
		done <- server60() // 为了方便演示，server40会在10s后主动返回一个error
	}()

	var stopped bool
	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil { // 程序会在此阻塞，接收done中的消息
			fmt.Printf("error: %v\n", err)
		}
		if !stopped { // 当有server为done时，关闭其他server，此处通过close channel来“告知”其他server
			stopped = true
			close(stop)
		}
	}
}

func server6(addr string, stop <-chan struct{}) error {
	h := handler6{wg: sync.WaitGroup{}}
	s := http.Server{
		Addr:    addr,
		Handler: &h,
	}

	go func() {
		<-stop // goroutine会在此阻塞，直到stop内有消息，或stop被关闭
		fmt.Println("receive a stop signal or channel(stop) is closed")

		h.BeforeShutdown6()

		s.Shutdown(context.Background()) // 在此主动关闭server，该goroutine会结束
	}()

	return s.ListenAndServe() // 当s.Shutdown被调用时，此处会返回一个error，该goroutine也被结束
}

func (h *handler6) BeforeShutdown6() {
	h.shutdown = true
	const timeout = 15 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

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

func server60() error {
	time.Sleep(time.Second * 10)
	return fmt.Errorf("err")
}

type handler6 struct {
	wg sync.WaitGroup // 等价于request group
	shutdown bool
}

func (h *handler6) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (h *handler6) Event(data string, sleep int) {
	time.Sleep(time.Second * time.Duration(sleep))
	fmt.Println(data)
	h.wg.Done() // 当有一个请求处理完成时，Done()
}
