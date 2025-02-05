package main

import (
	"fmt"
)

func Map[T, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

func Filter[T any](slice []T, f func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func Reduce[T, U any](slice []T, initial U, f func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = f(result, v)
	}
	return result
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Map: Square each number
	squares := Map(numbers, func(n int) int { return n * n })
	fmt.Println("Squares:", squares)

	// Filter: Get even numbers
	evens := Filter(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Println("Evens:", evens)

	// Reduce: Sum all numbers
	sum := Reduce(numbers, 0, func(acc int, n int) int { return acc + n })
	fmt.Println("Sum:", sum)
}
