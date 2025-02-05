package main

import (
	"fmt"
	"reflect"
)


func curry(f interface{}) interface{} {
	ft := reflect.TypeOf(f)
	if ft.Kind() != reflect.Func {
		panic("curry: argument must be a function")
	}

	numArgs := ft.NumIn()
	args := make([]reflect.Value, 0, numArgs)

	var curried func(x reflect.Value) interface{}
	curried = func(x reflect.Value) interface{} {
		args = append(args, x)
		if len(args) == numArgs {
			return reflect.ValueOf(f).Call(args)[0].Interface()
		} else {
			return curried
		}
	}

	return curried
}

func add(a, b, c int) int {
	return a + b + c
}

func main() {
	curriedAdd := curry(add)

	add1 := curriedAdd.(func(reflect.Value) interface{})(reflect.ValueOf(1))
	add12 := add1.(func(reflect.Value) interface{})(reflect.ValueOf(2))
	result := add12.(func(reflect.Value) interface{})(reflect.ValueOf(3)).(int)

	fmt.Println(result) // Output: 6
}