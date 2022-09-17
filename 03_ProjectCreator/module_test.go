package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitter(t *testing.T) {
	text := "   folder"
	spaces, text := splitter(text, ' ')

	assert.Equal(t, spaces, "   ", "check the spaces")
	assert.Equal(t, text, "folder", "check the folder name")
}
