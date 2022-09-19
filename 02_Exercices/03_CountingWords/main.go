package main

import (
	"io/ioutil"
	"log"
	"strings"
)

const _filePath = "text.txt"

func readFile(filePath string) string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Panic()
	}
	return string(content)
}

func countWords(file string) int {
	words := strings.FieldsFunc(file, split)

	return len(words)
}

func split(s rune) bool {
	return s <= 'A' || s >= 'z'
}

func main() {
	// go test -run TestCountWords
}
