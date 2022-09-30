package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	testCases := []struct {
		description string
		input       []int
		want        []int
	}{
		{"short & equal ", []int{3, 2, 1}, []int{3, 2, 1}},
		{"short to be sorted", []int{1, 2, 3}, []int{3, 2, 1}},
		{"medium to be sorted", []int{5, 1, 2, 3, 5, 7}, []int{7, 5, 5, 3, 2, 1}},
		{"long", []int{1, 1, 1, 1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1, 1, 1, 1}},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.want, sort(tc.input))
		})
	}
}

func TestSortMinMax(t *testing.T) {
	testCases := []struct {
		description string
		input       []int
		want        []int
	}{
		{"sorted and even", []int{2, 3, 5, 7, 8, 9}, []int{9, 2, 8, 3, 7, 5}},
		{"unsorted and odd", []int{5, 3, 1, 7}, []int{7, 1, 5, 3}},
		{"short", []int{2, 3}, []int{3, 2}},
		{"equal", []int{9, 2, 8, 3, 7, 5}, []int{9, 2, 8, 3, 7, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.want, sortMaxMin(tc.input))
		})
	}
}
