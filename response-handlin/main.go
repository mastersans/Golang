package main

import (
	"fmt"
	"net/http"
)

func main() {
	const url = "http://localhost:15000"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	//fmt.Println(response)
	defer response.Body.Close()
	fmt.Println(response.ContentLength)
	fmt.Println(response.StatusCode)
}
