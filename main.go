package main

import (
	"fmt"
	"reflect"
)

func MakeMap(fpt interface{}) {
	fnV := reflect.ValueOf(fpt).Elem()
	fmt.Println(fnV)
	fnI := reflect.MakeFunc(fnV.Type(), implMap)
	fmt.Println(fnI)
	fnV.Set(fnI)
}

//TODO:completes implMap function.
// var implMap func([]reflect.Value) []reflect.Value

func implMap(args []reflect.Value) (results []reflect.Value) {
	if len(args) == 0 {
		return nil
	}

	var ret reflect.Value

	switch args[1].Kind() {
	case reflect.Slice:
		data := args[1].Interface().([]int)
		for i, v := range data {
			println(v)
			params := make([]reflect.Value, 1)
			params[0] = reflect.ValueOf(v)
			data[i] = int(args[0].Call(params)[0].Int())
		}
		println(data)
		ret = reflect.ValueOf(data)

	case reflect.Map:
		data := args[1].Interface().(map[int]int)
		for k, v := range data {
			println(k, v)
			params := make([]reflect.Value, 1)
			params[0] = reflect.ValueOf(v)
			data[k] = int(args[0].Call(params)[0].Int())
		}
		println(data)
		ret = reflect.ValueOf(data)

	}
	results = append(results, ret)
	return
}

func main() {

	println("It is said that Go has no generics.\nHowever we have many other ways to implement a generics like library if less smoothly,one is reflect.MakeFunc.\nUnderscore is a very useful js library,and now let's implement part of it-map,it will help you to understand how reflect works.\nPlease finish the 'implMap' function and pass the test.")
	var f func(func(e int) int, []int) []int
	MakeMap(&f)

	var a = []int{1, 2, 3}

	b := f(func(e int) int {
		return e + 1
	}, a)

	if b[0] != 2 || b[1] != 3 || b[2] != 4 || len(b) != 3 {
		fmt.Println("fail!")
	}

	var f2 func(func(e int) int, map[int]int) map[int]int

	MakeMap(&f2)

	var c = map[int]int{0: 0, 1: 1, 2: 3}

	d := f2(func(e int) int {
		return e + 1
	}, c)

	if d[0] != 1 || d[1] != 2 || d[2] != 4 || len(d) != 3 {
		fmt.Println("fail!")
	}
}
