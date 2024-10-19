package main

import (
	"testing"
)

func BenchmarkMain(b *testing.B) {
	var e Empty
	var d Doer = e

	b.Run("Direct", func(b *testing.B) {
		b.ReportAllocs()
		total := uint(0)
		for i := 0; i < b.N; i++ {
			total += e.Do()
		}
	})

	b.Run("Via Interface", func(b *testing.B) {
		b.ReportAllocs()
		sum := uint(0)
		for i := 0; i < b.N; i++ {
			sum += d.Do()
		}
	})
}
