package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortString(t *testing.T) {
	got := []string{"acb", "bac", "bca", "cab", "cba"}
	want := "abc"

	for _, word := range got {
		assert.Equal(t, extractLettersFromWord(word), want, "check if the word is sorted")
	}
}
