package main

import (
	"log"
	"sort"
	"strings"
)

// INPUT strs = ["eat","tea","tan","ate","nat","bat"]
// OUTPUT [["bat"],["nat","tan"],["ate","eat","tea"]]

var _words = []string{"eat", "tea", "tan", "ate", "nat", "bat"}

func sortWords(words []string) map[string][]string {
	matrix := make(map[string][]string)

	var old string

	log.Print(words)
	for i, word := range words {
		for j := i + len(matrix[sortString(word)]); j < len(words); j++ {
			if sortString(word) == sortString(words[j]) && words[j] != old {
				matrix[sortString(word)] = append(matrix[sortString(word)], words[j])
				old = words[j]
			}
		}
	}
	return matrix
}
func sortString(word string) string {
	res := strings.Split(word, "")

	sort.Strings(res)

	return strings.Join(res, "")
}
