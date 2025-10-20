package benchmarking

import (
	"testing"
)

// https://pkg.go.dev/testing#hdr-Benchmarks
// https://go.dev/blog/pprof
// https://go.dev/blog/testing-b-loop

var testString = "   \tThis is a string \n with \r many \t spaces and \n newlines.   "

func BenchmarkNaive(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		RemoveWhitespaceThrasher(testString)
	}
}

func BenchmarkBuilderLoop(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		RemoveWhitespaceBuilder(testString)
	}
}

func BenchmarkInPlace(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		RemoveWhitespaceInPlace(testString)
	}
}
