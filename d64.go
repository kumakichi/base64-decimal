package d64

import (
	"errors"
	"fmt"
)

const (
	maxStrLen = 11
	charNum   = 64
)

var (
	indexTable = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '{', '|'}
	num2char   map[uint64]byte
	char2num   map[byte]uint64
)

func init() {
	num2char = make(map[uint64]byte, charNum)
	char2num = make(map[byte]uint64, charNum)
	initMap()
}

// set index table which will be used in convertion between decimal base64 and base10
func SetIndexTable(idxTable []byte) error {
	indexTable = idxTable
	return initMap()
}

// base 10 to base 64
func D10ToD64(n uint64) string {
	var numSlice [maxStrLen]byte
	idx := maxStrLen
	for n > 0 {
		idx--
		rem := n & 63
		numSlice[idx] = num2char[rem]
		n = n >> 6
	}

	return string(numSlice[idx:])
}

// base 64 to base 10
func D64ToD10(s string) (ret uint64) {
	arr := []byte(s)
	l := len(arr)
	for i, k := l-1, 0; i >= 0; i, k = i-1, k+6 {
		ret += char2num[arr[i]] << k
	}
	return
}

func initMap() error {
	for k, v := range indexTable {
		if v < '!' || v > '~' {
			return errors.New("index table supports only ascii")
		}

		num2char[uint64(k)] = v
		char2num[v] = uint64(k)
	}

	if len(num2char) != charNum {
		return errors.New(fmt.Sprintf("index table length error, should be %d", charNum))
	}

	return nil
}
