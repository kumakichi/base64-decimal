package d64

import (
	"math"
	"math/rand"
	"testing"
)

const (
	maxNum    = uint64(math.MaxUint64)
	maxNumStr = "F||||||||||"
)

func BenchmarkDecimalToX64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DecimalToX64(maxNum)
	}
}

func BenchmarkX64ToDecimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X64ToDecimal(maxNumStr)
	}
}

func BenchmarkSetIndexTable(b *testing.B) {
	indexTable := make([]byte, 64)
	a := rand.Perm(64)
	for i, v := range a {
		indexTable[i] = byte(v + '!')
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SetIndexTable(indexTable)
	}
}
