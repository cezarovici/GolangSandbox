package main

import (
	"bufio"
	"os"
	"sort"
)

// INPUT strs = ["eat","tea","tan","ate","nat","bat"]
// OUTPUT [["bat"],["nat","tan"],["ate","eat","tea"]]

type sortRuneString []rune

const filePath = "words.txt"

func readFile(filePath string) ([]string, error) {
	fileHandler, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fileHandler.Close()

	var res []string

	scanner := bufio.NewScanner(fileHandler)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	return res, nil
}

func sortString(input string) string {
	runeArray := []rune(input)
	sort.Sort(sortRuneString(runeArray))

	return string(runeArray)
}

func (s sortRuneString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRuneString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRuneString) Len() int {
	return len(s)
}

func main() {
	//log.Print(readFile(filePath))
	sortString("cezar")
}
