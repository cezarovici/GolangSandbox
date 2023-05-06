package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringCompression(t *testing.T) {
	type testCase struct {
		test   string
		input  string
		output string
	}

	testCases := []testCase{
		{
			test:   "simple test",
			input:  "aabcccccaaa",
			output: "a2b1c5a3",
		},
		{
			test:   "less input than output",
			input:  "abc",
			output: "abc",
		},
		{
			test:   "Single letter",
			input:  "aaaaaa",
			output: "a6",
		},
		{
			test:   "multiple letters",
			input:  "aaaaaabbbbbb",
			output: "a6b6",
		},
		{
			test:   "combined letters",
			input:  "aaaaaabbbbbccccdddeef",
			output: "a6b5c4d3e2f1",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			require.Equal(t, tc.output, stringCompression(tc.input))
		})
	}
}
