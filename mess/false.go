package main

import (
	"bytes"
	"fmt"
	"os"
)

type Duck interface {
	Walk()
	Quack()
}

type Cat struct {
	Name string
}

func (c *Cat) Walk() {
	fmt.Printf("%p\n", &c)
	fmt.Println("cat walk")
}

func (c *Cat) Quack() {
	fmt.Println("cat quack")
	c.Name = "Duck"
}

func main() {
	c := Cat{}
	var d Duck = &c
	fmt.Printf("%p\n", &c)
	d.Walk()
	d.Quack()

	c.Walk()
	c.Quack()
	bytes.NewBuffer
}
