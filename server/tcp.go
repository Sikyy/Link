package server

import (
	"fmt"
	"log"
	"net"
)

func StartServer() {
	// 1. 监听端口
	var (
		listener net.Listener
		err      error
		conn     net.Conn
	)
	if listener, err = net.Listen("tcp", "0.0.0.0:8989"); err != nil {
		log.Fatalf("listen failed, err: %s \n", err.Error())
	}

	// 2. 建立连接并挂起
	for {
		if conn, err = listener.Accept(); err != nil {
			log.Fatalf("create conn failed, err: %s \n", err.Error())
		}

		// 3. 处理消息
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	for {
		var buf [128]byte
		var n, err = conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read from connect failed, err: %v\n", err)
			break
		}
		var str = string(buf[:n])
		fmt.Printf("receive from client, data: %v\n", str)
	}
}
