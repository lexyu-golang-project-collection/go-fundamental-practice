package main

import (
	"testing"
	"unique"
)

var (
	s1 = "just a really long string that's really long"
	s2 = "just a really long string that's really long"
)

func BenchmarkStringComparison(b *testing.B) {
	// for i := 0; i < b.N; i++ {
	// 	_ = (s1 == s2)
	// }

	for range b.N {
		_ = (s1 == s2)
	}
}

func BenchmarkUniqueStringComparison(b *testing.B) {
	h1 := unique.Make(s1)
	h2 := unique.Make(s2)
	for range b.N {
		_ = (h1 == h2)
	}
}
