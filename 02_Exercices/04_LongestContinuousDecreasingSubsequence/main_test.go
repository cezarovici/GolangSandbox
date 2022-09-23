package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseArray(t *testing.T) {
	testCases := []struct {
		description string
		input       []int
		want        []int
	}{
		{"short", []int{1, 2, 3}, []int{1}},
		{"medium", []int{3, 2, 1}, []int{3, 2, 1}},
		{"long", []int{12, 9, 10, 8, 6, 5, 3, 4, 2, 1}, []int{10, 8, 6, 5, 3}},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.want, parseArray(tc.input))
		})
	}
}
