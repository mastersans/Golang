package main

import "fmt"

func main() {
	language := make(map[string]string)
	language[".js"] = "javascript"
	language[".py"] = "python"
	language[".java"] = "java"
	language[".cpp"] = "cplusplus"
	//here as we can see printf is only used to get single string printed
	fmt.Printf(language[".py"])
	delete(language, ".py")
	fmt.Println(language)
	//for loop using the range key word
	for key, value := range language {
		fmt.Printf("for the key : %v, the value is %v\n", key, value)
	}
	//here we can also observe that the key:value pair
	// is stored in the order of insertion
}
