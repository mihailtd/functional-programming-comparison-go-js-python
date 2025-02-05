package main

import "fmt"


func createGreeting(greeting string) func(string) string {
	return func(name string) string {
			return fmt.Sprintf("%s %s", greeting, name)
	}
}

func main() {
	firstGreeting := createGreeting("Well, hello there ")
	secondGreeting := createGreeting("Hola ")
	fmt.Println(firstGreeting("Remi"))
	fmt.Println(secondGreeting("Ana"))
}
