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

func Strcat1(a, b string) string {
	return a + b
}

func Strcat2(a, b string) string {
	return string(append([]byte(a), b...))
}

func Strcat3(a, b string) string {
	sb := &strings.Builder{}
	sb.WriteString(a)
	sb.WriteString(b)
	return sb.String()
}

func Strcat4(a, b string) string {
	buf := bytes.NewBufferString(a)
	buf.WriteString(b)
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

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func LCM2(a, b int) int {
	aa := a
	for ; aa%b != 0; aa += a {
	}
	return aa
}

func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := n; i >= 1; i-- {
		for j := 1; j < i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}

func SelectSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		k := i - 1
		for j := i; j < n; j++ {
			if arr[j] < arr[k] {
				k = j
			}
		}
		if k != i-1 {
			arr[k], arr[i-1] = arr[i-1], arr[k]
		}
	}
	return arr
}

func InsertSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		for j, k := i, arr[i]; ; j-- {
			if j > 0 {
				if arr[j-1] < k {
					break
				}
				arr[j] = arr[j-1]
			} else {
				arr[0] = k
				break
			}
		}
	}
	return arr
}

func Partition(arr []int, k int) ([]int, int) {
	n := len(arr)
	if k >= n {
		return arr, -1
	}
	p := arr[k]
	if k != 0 {
		arr[k], arr[0] = arr[0], arr[k]
	}

	i, j := 0, n-1
	for i < j {
		for ; i < j && p <= arr[j]; j-- {
		}
		arr[i] = arr[j]

		for ; i < j && arr[i] <= p; i++ {
		}
		arr[j] = arr[i]
	}
	arr[i] = p
	return arr, i
}

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	_, k := Partition(arr, 0)
	QuickSort(arr[:k])
	QuickSort(arr[k+1:])
	return arr
}
