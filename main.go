package main

import (
	"fmt"
	"reflect"
)

//Reverse reverses a slice.
// var Reverse func(slice interface{})

func Reverse(slice interface{}) {
	fmt.Println("obj is a", reflect.TypeOf(slice).Kind())
	n := reflect.ValueOf(slice).Elem() // 返回slice prt指向的内容

	fmt.Println(n, reflect.TypeOf(n).Kind(), n.Index(0))
	swap := reflect.Swapper(n.Interface()) // 返回n的interface格式传入
	for i, j := 0, n.Len()-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func main() {
	println("Please edit main.go,and complete the 'Reverse' function to pass the test.\nYou should use reflect package to reflect the slice type and make it applly to any type.\nTo run test,please run 'go test'\nIf you pass the test,please run 'git checkout l2' ")
	a := []int{1, 2, 3, 4, 5}
	Reverse(&a)
	fmt.Println(a)
}
