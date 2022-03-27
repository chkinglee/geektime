// Package main
// @Author      : lilinzhen
// @Time        : 2022/3/27 20:11:08
// @Description :
package main

import (
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
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("read error: %v\n", err)
			return
		}
		log.Println(conn.RemoteAddr().String(), "receive data length:", n)
		log.Println(conn.RemoteAddr().String(), "receive data string:", string(buffer[:n]))
	}
}

