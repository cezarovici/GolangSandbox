package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortSwapping(t *testing.T) {
	testCases := []struct {
		description string
		input       []int
		want        []int
	}{
		{"short", []int{2, 1, 0}, []int{0, 1, 2}},
		{"medium", []int{2, 0, 2, 1, 2, 0}, []int{0, 0, 1, 2, 2, 2}},
		{"long", []int{2, 2, 2, 0, 0, 0, 1, 1, 1}, []int{0, 0, 0, 1, 1, 1, 2, 2, 2}},
		{"equal2", []int{2, 2, 2}, []int{2, 2, 2}},
		{"equal1", []int{1, 1, 1}, []int{1, 1, 1}},
		{"nil", []int{}, []int{}},
		{"increasing", []int{0, 1, 2}, []int{0, 1, 2}},
		{"decreasing", []int{2, 1, 0}, []int{0, 1, 2}},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.want, sortBySwapping(tc.input))
		})
	}
}
