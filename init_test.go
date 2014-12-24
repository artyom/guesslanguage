package guesslanguage

import "testing"

func BenchmarkPrepareBlocks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = prepareBlocks()
	}
}
