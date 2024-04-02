package main

import "fmt"

func main() {
	// so basicly array are useless in the golang only useful if we need a fixed
	// size data structure
	// slices are similar to the vectors of c++
	var myarray [5]int
	for i := 0; i < 5; i++ {
		myarray[i] = i
	}
	fmt.Println(myarray)
	myarray2 := [5]int{7, 4, 6, 5, 2}
	fmt.Println(myarray2)

	var courses = []string{"english", "hindi", "maths", "science", "socialscience", "sanskriti"}
	fmt.Println(courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
