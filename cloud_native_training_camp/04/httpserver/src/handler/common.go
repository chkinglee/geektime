// Package handler
// @Author      : lilinzhen
// @Time        : 2022/1/2 10:31:08
// @Description :
package handler

import (
	"cloud_native_training_camp/04/httpserver/src/metrics"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// Func 0. 为所有请求设置公共header
type Func func(writer http.ResponseWriter, request *http.Request, statusCode *int)

// CommonHandler 0. 为所有请求设置公共header
func CommonHandler(f Func) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		start := time.Now()

		// 1. 接收客户端request，并将request中带的header写入response header
		headers := request.Header
		for key, value := range headers {
			writer.Header().Set(key, value[0])
		}

		// 2. 读取当前系统的环境变量中的VERSION配置，并写入response header
		osVersion := os.Getenv("VERSION")
		writer.Header().Set("version", osVersion)

		// 各自请求的业务处理
		var statusCode int
		f(writer, request, &statusCode)

		// 3. Server端记录访问日志包括客户端IP，HTTP返回码，输出到server端的标准输出
		log.Println(request.Host, statusCode)

		// 添加0-2秒随机延时
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		cost := time.Since(start)
		metrics.RequestsCost.WithLabelValues(request.Method, request.RequestURI).Observe(cost.Seconds())
	}
}
