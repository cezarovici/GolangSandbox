package main

import (
	"fmt"
	"sort"
	"strings"
)

// INPUT strs = ["eat","tea","tan","ate","nat","bat"]
// OUTPUT [["bat"],["nat","tan"],["ate","eat","tea"]]

var _words = []string{"eat", "tea", "tan", "ate", "nat", "bat"}

func main() {
	parsed := parseWords(_words)

	for _, words := range parsed {
		fmt.Println(words)
	}
}

func extractLettersFromWord(word string) string {
	res := strings.Split(word, "")

	sort.Strings(res)

	return strings.Join(res, "")
}

func parseWords(words []string) map[string][]string {
	res := make(map[string][]string)

	for _, word := range words {
		letters := extractLettersFromWord((word))

		res[letters] = append(res[letters], word)
	}

	return res
}
