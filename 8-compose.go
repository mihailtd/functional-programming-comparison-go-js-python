package main

import "fmt"

func compose[T any](funcs ...func(T) T) func(T) T {
	return func(x T) T {
		result := x
		for _, f := range funcs {
			result = f(result)
		}
		return result
	}
}

func addOne(x int) int {
	return x + 1
}

func square(x int) int {
	return x * x
}

func multiplyByTwo(x int) int {
	return x * 2
}

func main() {
	composed := compose(addOne, multiplyByTwo, square)
	result := composed(5) // ((5 + 1) * 2)^2 = 144
	fmt.Println("Composed result:", result)
}