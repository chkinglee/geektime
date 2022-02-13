// Package main
// @Author      : lilinzhen
// @Time        : 2022/2/13 13:54:55
// @Description :
package main

import (
	"fmt"
	"net/http"
)

func main() {
	/*
		在case_1中，main函数阻塞在http.ListenAndServe的调用中，而使后续代码无法执行
		goroutine使main函数可以在一个http服务启动后，做更多的事，比如再启动一个http服务

		但可能存在的问题：
		1、server22由于某种原因退出后，整个main会退出，即进程退出
		2、server21由于某种原因退出后，由于server2是正常的，main感知不到server1已经退出

		Never start a goroutine without knowning when it will stop
	*/
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello World")
		fmt.Println("Hello World")
	})

	go server21()
	fmt.Println("8888 listened success.")
	server22()
	fmt.Println("8889 listened success.")
}

func server21() {
	http.ListenAndServe(":8888", nil)
}

func server22() {
	http.ListenAndServe(":8889", nil)
}
