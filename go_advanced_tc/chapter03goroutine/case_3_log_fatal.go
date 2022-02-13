// Package main
// @Author      : lilinzhen
// @Time        : 2022/2/13 14:14:04
// @Description :
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	/*
		为了解决case_2中存在的问题，将2个server都使用goroutine，并使用log.Fatal使http产生err后使程序强制退出

		但如下的写法存在的问题：
		任意一个server产生err，程序都会直接退出，那么其他server也会被强制退出，不管其他server的请求处理到什么程度
		即我们无法人为的控制goroutine以友好的方式退出

		Never start a goroutine without knowning when it will stop
	*/
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello World")
		fmt.Println("Hello World")
	})

	go server31()
	fmt.Println("8888 listened success.")
	go server32()
	fmt.Println("8889 listened success.")
	select {}
}

func server31() {
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}

func server32() {
	if err := http.ListenAndServe(":8889", nil); err != nil {
		log.Fatal(err)
	}
}
