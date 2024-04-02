package main

import "fmt"

func main() {
	greeter()
	result := adder(1, 3)
	fmt.Println(result)
	result2, sentence := proAdder(1, 3, 4)
	fmt.Println(result2, sentence)
}

func proAdder(values ...int) (int, string) {
	answer := 0
	for _, val := range values {
		answer += val
	}
	return answer, "mahesh dalla"
}
func adder(a int, b int) int {
	return a + b
}
func greeter() {
	fmt.Println("Jai Shree Ram")
}

//methods
