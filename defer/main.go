package main

import "fmt"

func main() {
	defer fmt.Println("Ram Ram")
	fmt.Println("na Bhai")
	fmt.Println("sarayane")
	myDefer()
}

func myDefer() {
	for i := 0; i < 6; i++ {
		defer fmt.Println(i)
	}
}
