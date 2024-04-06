package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int    `json:"paisa"`
	Platform string
	// "-" is for not showing this part in the json
	Password string `json:"-"`
	//omitempty don't show it if its nil
	Tags []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to Json Video")
	EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "learncodeonline.in", "strongpassword", []string{"js", "react"}},
		{"ReactJS Bootcamp", 199, "learncodeonline.in", "strongpassword", []string{"js", "react"}},
		{"ReactJS Bootcamp", 199, "learncodeonline.in", "strongpassword", nil},
	}
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

// decoding of consuming the json data
func DecodeJson() {
	jsonDataFromSomewhere := []byte(`
		{
			"coursename": "ReactJS Bootcamp",
			"paisa": 299,
			"Platform": "learncodeonline.in",
			"tags": ["js","react"]
		}
		`)
	var lco course
	json.Unmarshal(jsonDataFromSomewhere, &lco)
	fmt.Printf("%#v\n", lco)
}
