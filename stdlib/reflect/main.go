package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func main() {
	t := T{33, "skidoo"}
	v := reflect.ValueOf(&t).Elem()
	typeOft := v.Type()

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOft.Field(i), f.Type(), f.Interface())
	}
}