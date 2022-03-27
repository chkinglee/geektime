// Package main
// @Author      : lilinzhen
// @Time        : 2022/3/27 20:40:33
// @Description :
package main

import (
	"bufio"
	"chapter09network/utils"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept error: %v\n", err)
			continue
		}
		// 开始goroutine监听连接
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	rd := bufio.NewReader(conn)
	for {
		data, err := rd.ReadSlice(utils.ConstDelimiter)
		if err != nil {
			log.Printf("read error: %v\n", err)
			return
		}
		log.Println(string(data[:len(data)-1]))
	}
}
