package main

import (
	"log"
	"strings"
)

func splitter(s string, delimitator byte) (string, string) {
	var (
		value string
		name  string
	)

	for i, element := range s {
		if element == rune(delimitator) {
			value = value + s[i:i+1]

			continue
		}
		name = s[i:]

		break
	}

	return value, name
}

func changeName(filename string) string {
	var newstring []byte
	split := strings.Split(filename, " ")

	for i, name := range split {
		tester := []byte(name)
		if tester[0] == '#' {
			split := strings.Split(filename, "#")
			return split[1]
		}
		for _, elem := range tester {
			if elem != '.' {
				newstring = append(newstring, elem)
			} else {
				break
			}
		}
		if i+1 != len(split) {
			newstring = append(newstring, '.', 'g', 'o', ' ')
		}
	}
	return string(newstring)
}

func verifyPackage(pack, line string) string {
	_, pack = splitter(pack, ' ')

	if pack == "t" {
		return "_test.go"
	}

	return ""
}

func getCase(line string) int {
	options := []byte{'!', '#', '.'}
	byteSlice := []byte(line)

	//		   		0			   1    2
	// cases (point of reference, pack, file)
	for _, elem := range byteSlice {
		for j, opt := range options {
			if elem == opt {
				return j
			}
		}
	}
	// case folder
	return 3
}

func (t *tree) addElement(line string) {
	fileInfo := []string{"Point of reference", "Pack", "File", "Directory"}
	split := strings.Split(line, " ")
	log.Print(fileInfo[getCase(line)])

	x := fileInfo[getCase(line)]
	_, newLine := splitter(line, ' ')

	switch x {
	case "Point of reference":
		t.path = split[1]
		t.delimitator = ""
	case "Pack":
		t.pack = changeName(newLine)
	case "File":
		t.files = append(t.files, changeName(newLine)+verifyPackage(t.pack)+".go")
	case "Directory":
		spaces, name := splitter(newLine, ' ')
		t.folders = append(t.folders, name)
		t.delimitator = spaces
	}
}

func parsing(filename string) {
	var (
		dl         = new(DoubleList)
		t          = new(tree)
		old        = len(t.delimitator)
		slice      = readFile(filename)
		difference int
	)

	for _, name := range slice {
		delimitator, _ := splitter(name, ' ')
		old = len(t.delimitator)
		difference = len(delimitator) - old

		if difference == 0 || getCase(name) != 3 || getCase(name) == 2 {
			t.addElement(name)

		} else if difference > 0 {

			dl.append(t)
			x := new(tree)
			if t.pack != "t" {
				x.pack = t.pack
			}

			t = x
			t.addElement(name)
		}
	}
	dl.append(t)
	dl.displayHead()
}

func main() {
	parsing("filetext.txt")
}
