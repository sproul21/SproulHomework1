package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func countWords(str string) map[string]int {
	fileasBytes, err := ioutil.ReadFile(str)

	fileContents := string(fileasBytes)

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(fileContents, " ")

	words := strings.Split(processedString, " ")

	counts := make(map[string]int)

	for _, word := range words {
		strings.ToLower(word)
		word = strings.ToLower(word)
		_, ok := counts[word]
		if ok {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}

	return counts

}
func reportResults(str string) {
	for index, element := range countWords(str) {
		fmt.Println(index, ":", element)

	}

}

func main() {
	var str string
	fmt.Println("Enter File Name: ")
	fmt.Scan(&str)
	reportResults(str)

}
