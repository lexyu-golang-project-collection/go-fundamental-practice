package main

import (
	"io"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	var e Empty
	var d io.Reader = e

	b.Run("Direct", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var buf [100]byte
			e.Read(buf[:])
		}
	})

	b.Run("Via Interface", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var buf [100]byte
			d.Read(buf[:])
		}
	})
}
