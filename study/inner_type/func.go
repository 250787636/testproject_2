package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args"+
		"(%d,%d)", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//func swap(a, b *int) {
//	*a, *b = *b, *a
//}
func swap(a, b int) (int, int) {
	return b, a
}

func mainf() {
	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))
	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println(a, b)
}
