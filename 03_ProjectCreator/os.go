package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func readFile(filename string) []string {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	fileSlice := strings.Split(string(fileContent), "\n")
	return fileSlice
}

func createDirectory(path string) string {
	errDir := os.MkdirAll(path, 0777)
	if errDir != nil {
		log.Print(errDir)
	}

	newDir := filepath.Join(path)
	newDir = changeDirectory(newDir)

	return newDir
}
func getCurrentDirectory() string {
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Print(err)
	}
	return currentDirectory
}

func changeDirectory(path string) string {
	err := os.Chdir(filepath.Join(getCurrentDirectory(), path))
	if err != nil {
		log.Print(err)
	}

	return getCurrentDirectory()
}

func writeFile(filename string, text []byte) {
	err := ioutil.WriteFile(filename, text, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func checkDirectories(directories []string) bool {
	found := true

	for i, _ := range directories {
		_, err := os.Stat(directories[i])

		if os.IsNotExist(err) {
			log.Fatal("File not found!")
			found = false
		}
	}

	return found
}
