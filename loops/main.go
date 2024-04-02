package main

import "fmt"

func main() {
	days := []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	//here the range return 2 thing one is the index and other is the value
	//hence we use the _, val type thing
	for _, it := range days {
		if it == "friday" {
			goto lco
		}
		fmt.Println(it)
	}
	//this is label
lco:
	fmt.Println("ohh pencho friday hai")
}
