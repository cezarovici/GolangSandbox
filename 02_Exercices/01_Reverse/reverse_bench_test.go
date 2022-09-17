package main

import (
	"testing"
)

/*
BenchmarkSwap_the_symmetrical_elements-4         7193155               150 ns/op
BenchmarkSwap_the_symmetrical_elements-4         7019782               156 ns/op
BenchmarkSwap_the_symmetrical_elements-4         7503519               150 ns/op
BenchmarkSwap_the_symmetrical_elements-4         7542765               152 ns/op
BenchmarkSwap_the_symmetrical_elements-4         7562985               154 ns/op
*/
func BenchmarkSwapTheSymmetricalElements(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseAppEnd("cezar")
	}
}

/*
BenchmarkReverseByRune-4                        14136025                88.1 ns/op
BenchmarkReverseByRune-4                        13378930                79.8 ns/op
BenchmarkReverseByRune-4                        15764301                83.0 ns/op
BenchmarkReverseByRune-4                        14178771                85.0 ns/op
BenchmarkReverseByRune-4                        13868714                80.8 ns/op
*/
func BenchmarkReverseByRune(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseByRune("cezar")
	}
}
