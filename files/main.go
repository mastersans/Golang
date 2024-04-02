package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println()
	content := "this is the test content"
	file, err := os.Create("./myFile.txt")
	checkErr(err)
	length, err := io.WriteString(file, content)
	checkErr(err)
	//fmt.Println(os.Getpid())
	fmt.Println("Length is : ", length)
	//time.Sleep(20 * time.Second)
	readFile("./myFile.txt")
	defer file.Close()
}

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	checkErr(err)
	fmt.Println("The data inside the file is as follows :\n", string(data))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
