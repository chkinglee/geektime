package main

import (
	"log"
	"net/http"
	"os"
)

/**
模块二作业（第1次作业/第4次课的作业）
1. 接收客户端request，并将request中带的header写入response header
2. 读取当前系统的环境变量中的VERSION配置，并写入response header
3. Server端记录访问日志包括客户端IP，HTTP返回码，输出到server端的标准输出
4. 当访问localhost/healthz时，应返回200
*/

func main() {

	http.HandleFunc("/healthz", CommonHandler(healthz))
	http.HandleFunc("/a", CommonHandler(a))
	err := http.ListenAndServe(":8077", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// HandlerFunc 0. 为所有请求设置公共header
type HandlerFunc func(writer http.ResponseWriter, request *http.Request, statusCode *int)

// CommonHandler 0. 为所有请求设置公共header
func CommonHandler(handler HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
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
		handler(writer, request, &statusCode)

		// 3. Server端记录访问日志包括客户端IP，HTTP返回码，输出到server端的标准输出
		log.Println(request.Host, statusCode)
	}
}

// 4. 当访问localhost/healthz时，应返回200
func healthz(writer http.ResponseWriter, request *http.Request, statusCode *int) {
	*statusCode = http.StatusOK
	writer.WriteHeader(*statusCode)
	_, err := writer.Write([]byte("I am running"))
	if err != nil {
		return
	}
}

func a(writer http.ResponseWriter, request *http.Request, statusCode *int) {
	*statusCode = http.StatusAccepted
	writer.WriteHeader(*statusCode)
}
