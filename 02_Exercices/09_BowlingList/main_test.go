package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinMaxMed(t *testing.T) {
	testCases := []struct {
		description string
		input       []int
		want        []int
	}{
		{"shortMax", []int{10, 20, 30}, []int{2, 0, 20}},
		{"shortMin", []int{30, 20, 10}, []int{0, 2, 20}},
		{"medium", []int{20, 40, 50, 10}, []int{2, 1, 30}},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.want, minMaxMed(tc.input))
		})
	}
}
