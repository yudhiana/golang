package main

import "fmt"

func main() {

	hello := func() string { // <-- anonymous function
		return "hello anonymous"
	}()

	fmt.Println(hello)
}
