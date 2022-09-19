package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountWords(t *testing.T) {
	numberofWords := countWords(words)

	assert.Equal(t, numberofWords, 9, "check the number of words")
}
