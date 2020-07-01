package main

import "fmt"


func foo() string { // <-- anonymous function 
	return "hello anonymous"
}

func main() {
	fmt.Println(foo())
}
