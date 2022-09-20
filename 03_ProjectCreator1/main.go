package main

import (
	"bufio"
	"os"
	"strings"
)

type entry struct {
	folderInfo string
	indent     int
}

const _filePath = "folders.txt"

// readFile Helper reading file contents to a slice.
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

func typeofFile(fileName string) string {

	if strings.Contains(fileName, "!") {
		return "path"
	}

	if strings.Contains(fileName, ".") {
		return "file"
	}

	if strings.Contains(fileName, "#") {
		return "pack"
	}

	return "folder"
}

func lineParser(line, pack string) []string {
	var res []string

	files := strings.Split(line, " ")

	for _, file := range files {
		fileTrimmed := strings.TrimLeft(file, " ")
		if fileTrimmed != "" {
			res = append(res, isTestFile(pack, fileTrimmed))
		}
	}

	return res
}

func getPackage(line string) string {
	return line[2:]
}

func isTestFile(isPackage, line string) string {
	if isPackage == "t" {
		return line[:len(line)-3] + "_test.go"
	}

	return line
}
func convertToEntry(line string) *entry {
	trimmed := strings.TrimLeft(line, " ")

	return &entry{
		folderInfo: trimmed,
		indent:     len(line) - len(trimmed),
	}
}

func convertToEntries(content []string) []*entry {
	var res []*entry

	for _, line := range content {
		res = append(res, convertToEntry(line))
	}

	return res
}

func parse(entries []*entry) []string {
	var res []string

	var stackFolders stack
	var stackIndents stack
	var stackPackages stack

	for ix, entry := range entries {

		if typeofFile(entry.folderInfo) == "pack" {
			stackPackages.push(getPackage(entry.folderInfo))

			continue
		}

		if typeofFile(entry.folderInfo) == "file" {
			pack := stackPackages.front()
			files := lineParser(entry.folderInfo, pack)

			for _, file := range files {
				line := stackFolders.String() + "/" + file
				res = append(res, "touch "+line)
			}

			continue
		}

		if ix == 0 {
			stackFolders.push(entry.folderInfo)
			stackIndents.push(0)

			res = append(res, "mkdir "+stackFolders.String())

			continue
		}

		if entry.indent > stackIndents.peek().(int) {
			stackFolders.push(entry.folderInfo)
			stackIndents.push(entry.indent)
			stackPackages.push("")

			res = append(res, "mkdir "+stackFolders.String())

			continue
		}

		if entry.indent == stackIndents.peek().(int) {
			stackFolders.pop()
			stackFolders.push(entry.folderInfo)
			stackIndents.push(entry.indent)
			stackPackages.push("")

			res = append(res, "mkdir "+stackFolders.String())

			continue
		}

		for entry.indent < stackIndents.peek().(int) && len(stackIndents) > 1 {
			if entry.indent == stackIndents.peek().(int) {
				stackFolders.pop()
				stackPackages.pop()

				break
			}

			stackFolders.pop()
			stackIndents.pop()
		}

		stackFolders.push(entry.folderInfo)
		stackIndents.push(entry.indent)

		res = append(res, "mkdir "+stackFolders.String())
	}

	return res
}

func main() {
	// go test -run TestFile
}
