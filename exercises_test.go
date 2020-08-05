package exercises

import (
	"math/rand"
	"testing"
	"time"
)

var rnd *rand.Rand

func init() {
	s := rand.NewSource(time.Now().UnixNano())
	rnd = rand.New(s)
}

func initStrings() (string, string) {
	var bufA, bufB []byte
	for i := 0; i < 10000000; i++ {
		c := byte(rand.Int()%26 + int('a'))
		bufA = append(bufA, c)
		bufB = append(bufB, c+byte('A')-byte('a'))
	}
	return string(bufA), string(bufB)
}

func TestBaseNConverter_ToBaseN(t *testing.T) {
	c := NewBaseNConverter(Base62CharSet)
	t.Log(string(c.ToBaseN(4354536116143)))
}

func TestBaseNConverter_ToNumber(t *testing.T) {
	c := NewBaseNConverter(Base62CharSet)
	n, err := c.ToNumber([]byte("BOpKlv4L"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(n)
}

func TestStrcat1(t *testing.T) {
	println(Strcat1("abc", "def"))
}

func BenchmarkStrcat1(b *testing.B) {
	stringA, stringB := initStrings()
	var n int
	for i := 0; i < b.N; i++ {
		n = len(Strcat1(stringA, stringB))
	}
	b.Log(n)
}

func BenchmarkStrcat2(b *testing.B) {
	stringA, stringB := initStrings()
	var n int
	for i := 0; i < b.N; i++ {
		n = len(Strcat1(stringA, stringB))
	}
	b.Log(n)
}

func BenchmarkStrcat3(b *testing.B) {
	stringA, stringB := initStrings()
	var n int
	for i := 0; i < b.N; i++ {
		n = len(Strcat1(stringA, stringB))
	}
	b.Log(n)
}

func BenchmarkStrcat4(b *testing.B) {
	stringA, stringB := initStrings()
	var n int
	for i := 0; i < b.N; i++ {
		n = len(Strcat1(stringA, stringB))
	}
	b.Log(n)
}

func TestGCD(t *testing.T) {
	t.Log(GCD(56, 72))
	t.Log(GCD(72, 56))
	t.Log(GCD(293284, 104006))
}

func TestLCM(t *testing.T) {
	t.Log(LCM(3, 6))
	t.Log(LCM(34, 51))
	t.Log(LCM(293284, 104006))
}

func TestLCM2(t *testing.T) {
	t.Log(LCM2(3, 6))
	t.Log(LCM2(34, 51))
	t.Log(LCM2(293284, 104006))
}

func BenchmarkLCM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LCM(293284, 104006)
	}
}

func BenchmarkLCM2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LCM2(293284, 104006)
	}
}

func randArray(length int) []int {
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rnd.Intn(10000)
	}
	return arr
}

func checkSort(arr []int) bool {
	for i, n := 1, len(arr); i < n; i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func TestBubbleSort(t *testing.T) {
	arr := randArray(10000)
	BubbleSort(arr)
	if !checkSort(arr) {
		t.Fail()
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(randArray(10000))
	}
}

func TestSelectSort(t *testing.T) {
	arr := randArray(10000)
	SelectSort(arr)
	if !checkSort(arr) {
		t.Fail()
	}
}

func BenchmarkSelectSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectSort(randArray(10000))
	}
}

func TestInsertSort(t *testing.T) {
	arr := randArray(10000)
	InsertSort(arr)
	if !checkSort(arr) {
		t.Fail()
	}
}

func BenchmarkInsertSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertSort(randArray(10000))
	}
}

func TestPatition(t *testing.T) {
	arr := randArray(10)
	t.Log(arr)
	arr, k := Partition(arr, 6)
	t.Log(arr, k)
}

func TestQuickSort(t *testing.T) {
	arr := randArray(10000)
	QuickSort(arr)
	if !checkSort(arr) {
		t.Fail()
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(randArray(10000))
	}
}
