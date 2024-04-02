package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func saveHTMLToFile(url, filename string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	fmt.Printf("HTML body for %s appended to: %s\n", url, filename)
	return nil
}

func main() {
	startPage := 1
	endPage := 10
	baseURL := "https://leetcode.com/contest/biweekly-contest-113/ranking/"
	filename := "leetcode_contest_combined_html.html"

	// Create or truncate the file
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	for page := startPage; page <= endPage; page++ {
		url := fmt.Sprintf("%s%d", baseURL, page)

		err := saveHTMLToFile(url, filename)
		if err != nil {
			log.Fatal(err)
		}
	}
}
