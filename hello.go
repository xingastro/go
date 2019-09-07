package main

import (
	"fmt"
	"reflect"
	"time"
)

type User struct {
	Num int
	Status string
	Age uint8
}

func (u User) Hello() {
	fmt.Println("Hello world.")
}

func add(numP *int) {
	for i := 0; i < 10000; i++ {
		*numP++
	}
}

func desc(numP *int) {
	for i := 0; i < 10000; i++ {
		*numP--
	}
}

func put(nums *[]int) {
	for i := 0; i < 100000; i++ {
		*nums = append(*nums, i)
	}
}

func get(nums *[]int) {
	for {
		num := (*nums)[0]
		*nums = (*nums)[1:]
		fmt.Println(num)
	}
}

func main() {
	nums := make([]int, 0)
	go put(&nums)
	go get(&nums)
	time.Sleep(time.Minute)
	// fmt.Println(nums)
}


func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++{
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}