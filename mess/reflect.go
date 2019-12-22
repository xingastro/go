package main

import (
	"fmt"
	"log"
	"reflect"
)

type Human struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := Human{Name: "Adam", Age: 20}
	if nameField, ok := reflect.TypeOf(p1).FieldByName("Name"); !ok {
		log.Print("Failed to get 'Name' field.")
		return
	} else {
		fmt.Printf("Tag:name %s", nameField.Tag.Get("name"))
	}
}
