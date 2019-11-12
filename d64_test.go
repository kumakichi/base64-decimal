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

func BenchmarkD10ToD64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		D10ToD64(maxNum)
	}
}

func BenchmarkD64ToD10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		D64ToD10(maxNumStr)
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
