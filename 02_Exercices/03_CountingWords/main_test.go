package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountWords(t *testing.T) {
	numberofWords := countWords(readFile(_filePath))

	assert.Equal(t, numberofWords, 9, "check the number of words")

	log.Print("\n", readFile(_filePath))
}
