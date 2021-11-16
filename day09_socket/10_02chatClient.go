/*
E:\golang\HelloGolang\day09_socket>go run 10_02chatClient.go :9080
1
127.0.0.1:11767say:1

*/

package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main()  {
	startClient(os.Args[1])
	//startClient(":9080")
}

func startClient(tcpAddrStr string)  {

	// 解析
	tcpAddr, err := net.ResolveTCPAddr("tcp4", tcpAddrStr)
	if err != nil{
		log.Printf("Resovle tcp addr failed: %v\n", err.Error())
		return
	}

	// 向服务器拨号
	conn, err := net.DialTCP("tcp", nil , tcpAddr)
	if err != nil{
		log.Printf("Dial to server failed: %v\n", err.Error())
		return
	}

	buf := make([]byte, 1024)
	// 向服务器端发送消息
	go sendMsg(conn)

	// 接受来自服务器端的广播消息
	for{
		length, err := conn.Read(buf)
		if err != nil{
			log.Printf("recv server msg failed: %v\n", err.Error())
			conn.Close()
			os.Exit(0)
			break
		}
		fmt.Println(string(buf[0:length]))
	}
}

func sendMsg(conn net.Conn)  {
	username := conn.LocalAddr().String()

	for{
		var input string
		fmt.Scanln(&input)

		if input == "/q" || input == "/quit"{
			fmt.Println("byebye....")
			conn.Close()
			os.Exit(0)
		}

		// 处理消息
		if len(input) > 0 {
			msg := username + "say:" + input
			_, err := conn.Write([]byte(msg))
			if err != nil{
				conn.Close()
				break
			}
		}
	}

}


