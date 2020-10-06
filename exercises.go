package exercises

import (
	"bytes"
	"errors"
	"strings"
)

var Base62CharSet = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

type BaseNConverter struct {
	base      int
	encodeMap []byte
	decodeMap [256]int
}

func (c *BaseNConverter) ToBaseN(number int) []byte {
	var baseN []byte
	for number > 0 {
		a := number / c.base
		baseN = append(baseN, c.encodeMap[number-a*c.base])
		number = a
	}
	for i, m, n := 0, len(baseN), len(baseN)/2; i < n; i++ {
		j := m - i - 1
		baseN[i], baseN[j] = baseN[j], baseN[i]
	}
	return baseN
}

func (c *BaseNConverter) ToNumber(baseN []byte) (int, error) {
	number := 0
	for _, b := range baseN {
		if n := c.decodeMap[b]; n < 0 {
			return 0, errors.New("basen: illegal byte")
		} else {
			number = number*c.base + n
		}

	}
	return number, nil
}

func NewBaseNConverter(charSet []byte) *BaseNConverter {
	decodeMap := [256]int{}
	for i, n := 0, len(decodeMap); i < n; i++ {
		decodeMap[i] = -1
	}

	for i, b := range charSet {
		decodeMap[b] = i
	}
	return &BaseNConverter{
		base:      len(charSet),
		encodeMap: charSet,
		decodeMap: decodeMap,
	}
}

func Strcat1(arr []string) (ret string) {
	for _, s := range arr {
		ret += s
	}
	return ret
}

func Strcat2(arr []string) (ret string) {
	var buf []byte
	for _, s := range arr {
		buf = append(buf, s...)
	}
	return string(buf)
}

func Strcat3(arr []string) (ret string) {
	sb := &strings.Builder{}
	for _, s := range arr {
		sb.WriteString(s)
	}
	return sb.String()
}

func Strcat4(arr []string) (ret string) {
	buf := &bytes.Buffer{}
	for _, s := range arr {
		buf.WriteString(s)
	}
	return buf.String()
}

func GCD(a, b int) int {
	for {
		c := a % b
		if c == 0 {
			return b
		}
		a, b = b, c
	}
}

func GCD2(a, b int) int {
	if a == 0 {
		return b
	}
	return GCD2(b%a, a)
}

func LCM(a, b int) int {
	return a / GCD(a, b) * b
}

func LCM2(a, b int) int {
	aa := a
	for ; aa%b != 0; aa += a {
	}
	return aa
}

func StringKinds(arr []string) int {
	c := 0
	m := map[string]struct{}{}
	for _, s := range arr {
		if _, ok := m[s]; !ok {
			c++
			m[s] = struct{}{}
		}
	}
	return c
}

func StringKinds2(arr []string) int {
	c := 0
	t := NewTrieTree()
	for _, s := range arr {
		if exist := t.Set(s); !exist {
			c++
		}
	}
	return c
}

func StringKinds22(arr []string) int {
	c := 0
	t := NewTrieTree()
	for _, s := range arr {
		if exist := t.Set2(s); !exist {
			c++
		}
	}
	return c
}

func StringKinds3(h *HashTable, arr []string) int {
	c := 0
	for _, s := range arr {
		if exist, _ := h.Put(s, struct{}{}); !exist {
			c++
		}
	}
	return c
}

func RingIncrease(n, m int) int {
	return (n + 1) % m
}

func RingIncrease2(n, m int) int {
	if n+1 == m {
		return 0
	} else {
		return n + 1
	}
}

func RingIncrease3(n, m int) int {
	if t := n + 1; t == m {
		return 0
	} else {
		return t
	}
}

func ChangeSlice(arr []int) {
	arr[0] = 1
}

func SeqBefore(seq1, seq2 uint8) bool {
	println(int8(seq1 - seq2))
	return int8(seq1-seq2) < 0
}

func ReturnDefer() (i int) {
	defer func() { i++ }()
	return 1
}
