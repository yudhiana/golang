package main

import "fmt"

func foo() func(int) int { 
	i := 0
	return func(j int) int {
		i += j
		return i
	}
}

func main() {
	bar := foo()
	fmt.Println(bar(10)) // Prints 10
	fmt.Println(bar(10)) // Prints 20
}
