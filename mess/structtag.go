package main

import (
	"fmt"
	"encoding/json"
)

type Cat struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type int    `json:"type"`
}

func main() {
	c := &Cat{Id: 12432413, Name: "Adam", Type: 1}
	res, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}
