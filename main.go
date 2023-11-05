package main

import (
	"Link/client"
	"Link/server"
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: go run main.go [server|client]")
		return
	}

	mode := args[1]
	if mode == "server" {
		// 启动服务器端
		startServer()
	} else if mode == "client" {
		// 启动客户端
		startClient()
	} else {
		fmt.Println("Invalid mode. Use 'server' or 'client'.")
	}
}

func startServer() {
	// 启动tcp服务器端代码
	server.StartServer()
}

func startClient() {
	client.StartClient()
}

// println("Congratulations! Your are link to the world!")
