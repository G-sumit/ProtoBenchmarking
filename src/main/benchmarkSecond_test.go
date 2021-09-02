package main

import (
	"testing"
)

func BenchmarkAny(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		protoPracticeAny()
	}
}