package main

import (
	"fmt"
	"net"
	"net/http"
	_ "net/http/pprof"
)

func fibornacci(n int) int {
	a, b := 0, 1
	for a < n {
		a, b = b, a+b
	}
	return a
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
}

func createFBS(w http.ResponseWriter, r *http.Request) {
	var fbs []int
	for i := 0; i < 1000; i++ {
		fbs = append(fbs, fibornacci(i))
	}
	w.Write([]byte(fmt.Sprintf("%v", fbs)))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/createFBS", createFBS)

	net.Dial("tcp", "baidu.com:80")
	http.ListenAndServe(":8080", nil)
}