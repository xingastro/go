package main

import (
	"net"
	"log"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// Create a tcp listener
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			for {
				msg, err := ioutil.ReadAll(c)
				if err != nil {
					log.Fatal(err)
				}
				io.Copy(os.Stdout, msg)
				
			}
			defer c.Close()
		}(conn)
	}
}
