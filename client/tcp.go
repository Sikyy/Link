package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func StartClient() {
	sendMessage()
	fmt.Println("send over")
}

func sendMessage() {
	// 1. 连接服务器
	var (
		conn net.Conn
		err  error
	)

	if conn, err = net.Dial("tcp", "localhost:8989"); err != nil {
		log.Fatalf("connect failed, err: %s", err.Error())
	}
	defer conn.Close()

	var (
		inputReader  *bufio.Reader
		input        string
		trimmedInput string
	)
	// 2. 读取命令行输入
	inputReader = bufio.NewReader(os.Stdin)

	for {
		// 3. 直到读取到\n
		if input, err = inputReader.ReadString('\n'); err != nil {
			fmt.Printf("read from console failed, err: %s \n", err.Error())
			break
		}

		// 4. 读取Q时停止
		if trimmedInput = strings.TrimSpace(input); trimmedInput == "Q" {
			break
		}

		// 5. 发送消息
		if _, err = conn.Write([]byte(trimmedInput)); err != nil {
			fmt.Printf("write failed , err : %s\n", err.Error())
			break
		}
	}
}
