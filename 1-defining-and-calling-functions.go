package main

import (
    "fmt"
)

func square(i int) int {
	return i * i
}

func main() {
	result := square(5)
	fmt.Println(result)
}