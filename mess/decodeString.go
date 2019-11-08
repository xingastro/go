package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)


func main() {
	var s string
	fmt.Printf("input: ")
	fmt.Scan(&s)

	decode(s)
}

func decode(s string) {
	index := 0
	var codePoint uint16
	for i := 0; i < len(s); {
		if s[i] < utf8.RuneSelf {
			// ASCII character
			codePoint = uint16(s[i])
			i++
		} else {
			fb, err := strconv.ParseInt(strconv.FormatInt(int64(s[i]), 2)[4:], 2, 64)
			if err != nil {
				fmt.Printf("Error happened: %v\n", err)
				return
			}
			fb = fb << 12
			sb, err := strconv.ParseInt(strconv.FormatInt(int64(s[i+1]), 2)[2:], 2, 64)
			if err != nil {
				fmt.Printf("Error happened: %v\n", err)
				return
			}
			sb <<= 6
			tb, err := strconv.ParseInt(strconv.FormatInt(int64(s[i+2]), 2)[2:], 2, 64)
			if err != nil {
				fmt.Printf("Error happened: %v\n", err)
			}
			codePoint = uint16(fb + sb + tb)
			i += 3
		}
		fmt.Printf("index: %d; value: %v -- > %c\n", index, codePoint, codePoint)
		index++
	}
}
