package main

import (
	"testing"
)

func BenchmarkOneOf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		protoPracticeOneOf()
	}
}