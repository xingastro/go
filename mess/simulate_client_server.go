package main

import (
	"fmt"
	"time"
)

///**
//		模拟远程调用RPC：使用通道代替Socket实现RPC的过程
//		客户端与服务器运行在同一个进程，client and server are two independent goroutines
// */
//
//
// // Client
// func RPCClient(ch chan string, req string) (string, error) {
// 	ch <- req  // simulate a request from client to server
// 	select {
// 	case data := <- ch:
// 		return data, nil
//	}
// }
//
// // Server
// func RPCServer(ch chan string) {
// 	// Handle requests from different clients by an infinite loop
// 	for {
// 		data := <- ch
// 		fmt.Println("server received: ", data)
// 		ch <- "roger"
//	}
// }
//func main() {
//	// Simulate starting a server
//	ch := make(chan string)
//	go RPCServer(ch)
//
//	// Simulate sending data
//	receive, err := RPCClient(ch, "hi")
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println("client received:", receive)
//	}
//}

func RPCClient(ch chan string) {
	for {
		select {
		case ch <- "Hi server, I'm client.":
			fmt.Println("")
		case words := <-ch:
			fmt.Println("Server says:", words)
		}
	}
}

func RPCServer(ch chan string) {
	for {
		select {
		case words := <-ch:
			fmt.Println(words)
		case ch <- "I'm server":

		}
	}
}

func main() {
	ch := make(chan string)
	go RPCServer(ch)
	go RPCClient(ch)

	time.Sleep(time.Second)
}
