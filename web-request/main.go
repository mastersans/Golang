package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "https://lco.dev"

func PeformPostJsonRequest() {
	const jsonurl = "http://localhost:15000/post"
	requestBody := strings.NewReader(`
	{
		"name": "les goo"
		
	}`)
	response, err := http.Post(jsonurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const formurl = "http://localhost:15000/postform"
	data := url.Values{}
	data.Add("firstname", "Sanskar")
	data.Add("lastname", "Sharma")
	data.Add("email", "sanskarsharma3110@gmail.com")
	response, _ := http.PostForm(formurl, data)
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func main() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	PeformPostJsonRequest()
	PerformPostFormRequest()
	defer response.Body.Close()
	// data, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	panic(err)
	// }
	//fmt.Printf(string(data))
}
