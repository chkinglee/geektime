// Package main
// @Author      : lilinzhen
// @Time        : 2022/3/27 20:40:24
// @Description :
package main

import (
	"chapter09network/utils"
	"fmt"
	"net"
	"os"
	"time"
)

func sender(conn net.Conn) {
	for i := 0; i < 1000; i++ {
		words := fmt.Sprintf("{\"Id\":%d,\"Name\":\"golang\",\"Message\":\"message\"}", i)
		conn.Write(utils.PacketWithDelimiter([]byte(words)))
	}
	fmt.Println("send over")
}

func main() {
	server := "127.0.0.1:10000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	fmt.Println("connect success")

	go sender(conn)

	for {
		time.Sleep(1 * 1e9)
	}
}

