// Package chapter03goroutine
// @Author      : lilinzhen
// @Time        : 2022/2/13 13:46:38
// @Description :
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello World")
		fmt.Println("Hello World")
	})

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}

	// 在本示例中，程序一直run在 http.ListenAndServe，而导致下行日志无法被打印出来
	fmt.Println("8888 listened success.")
}
