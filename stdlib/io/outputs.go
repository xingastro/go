package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer
	fmt.Fprintln(&b, "Hello")

	b.Write([]byte("world"))

	b.WriteTo(os.Stdout)
}