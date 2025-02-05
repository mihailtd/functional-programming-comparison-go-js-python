package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	points1 := []Point{{1, 2}, {3, 4}}
	points2 := points1 // Creates a copy of the slice header and underlying array

	// Appending creates a new slice, doesn't modify points1
	points2 = append(points2, Point{5, 6})

	// Modifying an element in points2 creates a copy of that element.
	points2[0].X = 10

	fmt.Println(points1) // Output: [{1 2} {3 4}] (original unaffected)
	fmt.Println(points2) // Output: [{10 2} {3 4} {5 6}]
}