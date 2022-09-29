package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertStringToIndex(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		want        int
	}{
		{"1 1 2 3", "aabc", 1123},
		{"11 2 3", "kbc", 1123},
		{"1 12 3", "alc", 1123},
		{"1 1 23", "aaw", 1123},
		{"11 23", "kw", 1123},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.want, convertStringToIndex(tc.input))
		})
	}
}

func TestToBinary(t *testing.T) {
	testCases := []struct {
		description string
		input       int
		want        int
	}{
		{"1", 1, 1},
		{"2", 2, 10},
		{"5", 5, 101},
		{"16", 16, 10000},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.want, toBinary(tc.input))
		})
	}
}
