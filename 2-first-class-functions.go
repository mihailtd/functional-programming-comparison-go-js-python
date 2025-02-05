package main

import "fmt"

func applyOperation(operation func(int, int) int, a int, b int) int {
    return operation(a, b)
}

func add(a, b int) int {
    return a + b
}

func subtract(a, b int) int {
    return a - b
}

func main() {
    result := applyOperation(add, 10, 5)
    fmt.Println(result) // Output: 15

    result = applyOperation(subtract, 10, 5)
    fmt.Println(result) // Output: 5
}