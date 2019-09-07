package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal)
	// 监听所有信号
	signal.Notify(c)

	fmt.Println("Program started")
	s := <- c
	fmt.Println("Signal received: ", s)
}
