package main

import (
	"fmt"
	"io"
	"os"
	"net/http"
	"strings"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: mcurl url")
		os.Exit(-1)
	}
	if !strings.HasPrefix(os.Args[1], "http") {
		os.Args[1] = "https://" + os.Args[1]
	}
}

func main() {
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println("There is an error")
		return
	}
	io.Copy(os.Stdout, r.Body)

	if err := r.Body.Close(); err != nil {
		panic(err)
	}
}
