package main

import (
	"bufio"
	"log"
	"os"
	"sort"
)

// INPUT strs = ["eat","tea","tan","ate","nat","bat"]
// OUTPUT [["bat"],["nat","tan"],["ate","eat","tea"]]

type sortRuneString []rune

const filePath = "words.txt"

func readFile(filePath string) []string {
	var res []string
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return res
}

func sortWords(filePath string) map[string][]string {
	matrix := make(map[string][]string)
	slice := readFile(filePath)

	var old string

	log.Print(slice)
	for i, word := range slice {
		for j := i + len(matrix[sortString(word)]); j < len(slice); j++ {
			if sortString(word) == sortString(slice[j]) && slice[j] != old {
				matrix[sortString(word)] = append(matrix[sortString(word)], slice[j])
				old = slice[j]
			}
		}
	}
	return matrix
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
	log.Print(sortWords(filePath))
}
