package main

import (
	"io"
	"log"
	"net/http"
)

type MyHandler struct {

}

func (mh *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Accepted a request from: ", r.RemoteAddr)
	indexPage(w, r)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello index page")
}

func main() {
	http.Handle("/", &MyHandler{})
	log.Println("Start listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}