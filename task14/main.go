package main

import (
	"fmt"
	"reflect"
)

 
func main() {
	var num interface{} = 1

	var line interface{} = "line"

	var boolean interface{} = true

	ch := make(chan int)
    chanel := interface{}(ch)

	fmt.Println(reflect.TypeOf(num).String())
	fmt.Println(reflect.TypeOf(line).String())
	fmt.Println(reflect.TypeOf(boolean).String())
	fmt.Println(reflect.TypeOf(chanel).String())
 
}