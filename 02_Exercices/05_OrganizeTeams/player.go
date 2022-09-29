package main

import (
	"log"
	"regexp"
)

type player struct {
	grade float32
	name  string
}

func (p *player) setName(name string) {
	p.name = name
}

func (p *player) setGrade(grade float32) {
	p.grade = grade
}

func (p *player) setInfoPlayer(name string, grade float32) {
	p.setGrade(grade)
	p.setName(name)
}

func getInfo(line string) (string, int) {
	rm := regexp.MustCompile("[0-9]+")
	grade := rm.FindAllString(line, -1)

	return line[:(len(line) - 3)], 

}

func readPlayers(filepath string) {
	var players []player

	fileLines, err := readFile(filepath)
	if err != nil {
		log.Print(err)
	}

	for number, player := range fileLines {

		players[number].setInfoPlayer(infos[0])
	}
}
