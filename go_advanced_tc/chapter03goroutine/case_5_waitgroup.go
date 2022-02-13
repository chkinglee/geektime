// Package main
// @Author      : lilinzhen
// @Time        : 2022/2/13 15:52:50
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
		在case_4中，为了保证server同时创建同时关闭，在一个server产生error时，让其他所有server都主动关闭
		但可能导致一些请求没有来得及处理完，正常的server就被关闭了

		在本示例中，使用sync.WaitGroup，来保证server在所有请求处理完成后，再主动关闭

		但依旧存在一个问题：
		如果未处理完成的请求，处理时间过长，此时server又没有关闭，就可能导致又请求一直进，一直长时间处理，server可能永远无法关闭
	*/
	done := make(chan error, 3)
	stop := make(chan struct{})
	go func() {
		done <- server5(":8888", stop)
	}()
	go func() {
		done <- server5(":8889", stop)
	}()
	go func() {
		done <- server50() // 为了方便演示，server40会在10s后主动返回一个error
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

func server5(addr string, stop <-chan struct{}) error {
	h := handler5{wg: sync.WaitGroup{}}
	s := http.Server{
		Addr:    addr,
		Handler: &h,
	}

	go func() {
		<-stop // goroutine会在此阻塞，直到stop内有消息，或stop被关闭
		fmt.Println("receive a stop signal or channel(stop) is closed")

		h.wg.Wait() // 在Shutdown之前等待所有请求处理完成

		s.Shutdown(context.Background()) // 在此主动关闭server，该goroutine会结束
	}()

	return s.ListenAndServe() // 当s.Shutdown被调用时，此处会返回一个error，该goroutine也被结束
}

func server50() error {
	time.Sleep(time.Second * 10)
	return fmt.Errorf("err")
}

type handler5 struct {
	wg sync.WaitGroup // 等价于request group
}

func (h *handler5) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.wg.Add(1) // 当有一个请求来时，Add(1)

	sleep := rand.Intn(20)
	fmt.Fprintln(w, "Hello World", sleep)
	fmt.Println("Hello World", sleep)
	go h.Event("Event"+strconv.Itoa(sleep), sleep) // 在这里模拟业务的异步处理
}

func (h *handler5) Event(data string, sleep int) {
	time.Sleep(time.Second * time.Duration(sleep))
	fmt.Println(data)
	h.wg.Done() // 当有一个请求处理完成时，Done()
}
