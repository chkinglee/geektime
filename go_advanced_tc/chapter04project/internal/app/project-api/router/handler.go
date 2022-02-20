// Package router
// @Author      : lilinzhen
// @Time        : 2022/2/20 21:44:20
// @Description :
package router

import (
	"chapter04project/internal/configs"
	"context"
	"fmt"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type httpHandler struct {
	logger   *zap.Logger
	wg       sync.WaitGroup // 等价于request group
	shutdown bool           // 判断服务是否正常
}

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.shutdown {
		fmt.Fprintln(w, "internal server error")
		h.logger.Error("internal server error")
		return
	}
	h.wg.Add(1) // 当有一个请求来时，Add(1)

	sleep := rand.Intn(20)
	fmt.Fprintln(w, "Hello World", sleep)
	h.logger.Info("Hello World", zap.Any("sleep_time", sleep))
	go h.Event("Event"+strconv.Itoa(sleep), sleep) // 在这里模拟业务的异步处理
}

func (h *httpHandler) Event(data string, sleep int) {
	time.Sleep(time.Second * time.Duration(sleep))
	h.logger.Info(data)
	h.wg.Done() // 当有一个请求处理完成时，Done()
}

func (h *httpHandler) BeforeShutdown() {
	// 设置shutdown，不再处理新的http请求
	h.shutdown = true

	// 设置shutdown的超时时间
	ctx, cancel := context.WithTimeout(context.Background(), configs.ShutdownTimeout)
	defer cancel()

	// 等待所有请求处理完成
	ch := make(chan struct{}) // 作用等同于main中的channel(stop)
	go func() {
		h.wg.Wait() // 在Shutdown之前等待所有请求处理完成
		close(ch)
	}()

	select {
	case <-ch:
		h.logger.Info("all request done")
	case <-ctx.Done():
		h.logger.Info("handler request timeout")
	}
}

