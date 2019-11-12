package d64

const (
	maxStrLen = 11
)

var (
	indexTable = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '{', '|'}
	num2char   map[uint64]byte
	char2num   map[byte]uint64
)

func init() {
	num2char = make(map[uint64]byte, 64)
	char2num = make(map[byte]uint64, 64)
	initMap()
}

func SetIndexTable(idxTable []byte) {
	indexTable = idxTable
	initMap()
}

func DecimalToX64(n uint64) string {
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

func X64ToDecimal(s string) (ret uint64) {
	arr := []byte(s)
	l := len(arr)
	for i, k := l-1, 0; i >= 0; i, k = i-1, k+6 {
		ret += char2num[arr[i]] << k
	}
	return
}

func initMap() {
	for k, v := range indexTable {
		num2char[uint64(k)] = v
		char2num[v] = uint64(k)
	}
}
