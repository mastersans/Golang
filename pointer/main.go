package main

import "fmt"

func main() {
	var temp *int
	fmt.Println(temp)
	number := 69
	ptr := &number
	fmt.Println("this is the address of the number", ptr)
	fmt.Println("this is the value at address", *ptr)
}
