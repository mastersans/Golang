package main

import (
	"fmt"
	"net/url"
	"reflect"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	result, _ := url.Parse(myUrl)
	fmt.Println(reflect.TypeOf(result))
	//following are the different properties we can use
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	queryParameters := result.Query()
	fmt.Println(reflect.TypeOf(queryParameters))
	//essentially the .values is a key value pair hence we can actually retrive the info
	fmt.Println(queryParameters["paymentid"])
	for key, val := range queryParameters {
		fmt.Printf("the key is %s and the value for the corresponding key is %s\n", key, val)
		fmt.Println(reflect.TypeOf(val))
	}
	partsofUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=sanskar",
	}
	tempUrl := partsofUrl.String()
	fmt.Println(tempUrl)
}
