package main

import (
	"fmt"
)

func main() {
	//resp, err := http.Get("https://www.google.com")
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()

	// html, err := httputil.DumpResponse(resp, true)
	// fmt.Printf("%s\n", html)

	tryRecover()
}


func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			fmt.Println("I don't know what to do")
		}
	}()
	 fmt.Println("This is a error")
}