package main

import (
	"io"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	b.Run("Direct", func(b *testing.B) {
		b.ReportAllocs()
		var e Empty
		for i := 0; i < b.N; i++ {
			var buf [50]byte
			e.Read(buf[:])
		}
	})

	b.Run("Via Interface", func(b *testing.B) {
		b.ReportAllocs()
		var d io.Reader = Empty{}
		for i := 0; i < b.N; i++ {
			var buf [50]byte
			d.Read(buf[:])
		}
	})
}
