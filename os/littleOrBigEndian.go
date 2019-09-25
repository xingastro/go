package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i uint32 = 1
	ptr := unsafe.Pointer(&i)
	if *(*byte)(ptr) == 1 {
		fmt.Println("little endian")
	} else {
		fmt.Println("big endian")
	}
}
