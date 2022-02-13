// Package main
// @Author      : lilinzhen
// @Time        : 2022/2/13 14:43:23
// @Description :
package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	/*
		为了解决case_3中无法控制goroutine生命周期的问题，在本示例中使用channel控制goroutine的结束

		channel(done) 用来表示server的结束，当任意一个server产生error时，会将该error推送到该channel
		channel(stop) 用来表示其他server是否需要结束，当任意一个server产生error时，关闭该channel以使其他server关闭

		当运行该程序时：
			1、 会有2个server启动，分别监听8888和8889端口
			2、 会有一个“模拟的server”启动并在10s后返回一个error
			3、 当任意一个server返回error后，其他所有的server都将以友好的方式主动关闭
			4、 你可以通过curl localhost:8888 发送一些请求，这些请求的处理将是异步的（通过goroutine模拟）：在一个随机的时间内打印一行日志

		在此时会存在一个问题：
		服务依旧无法非常友好的关闭，当正常的server没有处理完一些请求时，server关闭，导致丢失了一些请求的处理。
	*/
	done := make(chan error, 3)
	stop := make(chan struct{})
	go func() {
		done <- server4(":8888", stop)
	}()
	go func() {
		done <- server4(":8889", stop)
	}()
	go func() {
		done <- server40() // 为了方便演示，server40会在10s后主动返回一个error
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

func server4(addr string, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler{},
	}

	go func() {
		<-stop // goroutine会在此阻塞，直到stop内有消息，或stop被关闭
		fmt.Println("receive a stop signal or channel(stop) is closed")
		s.Shutdown(context.Background()) // 在此主动关闭server，该goroutine会结束
	}()

	return s.ListenAndServe() // 当s.Shutdown被调用时，此处会返回一个error，该goroutine也被结束
}

func server40() error {
	time.Sleep(time.Second * 10)
	return fmt.Errorf("err")
}

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sleep := rand.Intn(10)
	fmt.Fprintln(w, "Hello World", sleep)
	fmt.Println("Hello World", sleep)
	go h.Event("Event"+strconv.Itoa(sleep), sleep) // 在这里模拟业务的异步处理
}

func (h handler) Event(data string, sleep int) {
	time.Sleep(time.Second * time.Duration(sleep))
	fmt.Println(data)
}
