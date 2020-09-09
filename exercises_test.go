package exercises

import (
	"math/rand"
	"testing"
	"time"
)

var rnd *rand.Rand
var rnd2 *rand.Rand

func init() {
	s := rand.NewSource(time.Now().UnixNano())
	rnd = rand.New(s)

	rnd2 = rand.New(rand.NewSource(0))
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

	t.Log(GCD2(56, 72))
	t.Log(GCD2(72, 56))
	t.Log(GCD2(293284, 104006))
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

func BenchmarkGCD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GCD(293284, 104006)
	}
}

func BenchmarkGCD2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GCD2(293284, 104006)
	}
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

func randArrayWithSeed(length int, seed int64) []int {
	s := rand.NewSource(seed)
	rnd := rand.New(s)
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
		InsertSort(randArrayWithSeed(50, 100))
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
		QuickSort(randArrayWithSeed(10000, 100))
	}
}

func BenchmarkQuickSort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort2(randArrayWithSeed(10000, 100))
	}
}

func TestQuickSort3(t *testing.T) {
	arr := randArray(10000)
	QuickSort3(arr)
	if !checkSort(arr) {
		t.Fail()
	}
}

func BenchmarkQuickSort3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort3(randArrayWithSeed(10000, 100))
	}
}

func TestMaxHeap(t *testing.T) {
	arr := []int{3, 0, 5, 20, 8, 9}
	m := HeapSort(arr)
	t.Log(m)
}

func initStrings2(length int) []string {
	ret := make([]string, length)
	var buf [32]byte
	const r = 'z' - 'a'
	for i := 0; i < length; i++ {
		n := rnd2.Int()%16 + 16
		for j := 0; j < n; j++ {
			buf[j] = byte(rnd2.Int31()%r + 'a')
		}
		ret[i] = string(buf[:n])
	}
	return ret
}

func initStrings3(length int, minLen int, maxLen int, kinds int) []string {
	ret := make([]string, kinds, length)
	var buf [32]byte
	const r = 'z' - 'a'
	for i := 0; i < kinds; i++ {
		n := rnd2.Int()%(maxLen-minLen) + minLen
		for j := 0; j < n; j++ {
			buf[j] = byte(rnd2.Int31()%r + 'a')
		}
		ret[i] = string(buf[:n])
	}

	for left := length - len(ret); left > 0; left = length - len(ret) {
		s := ret[rnd2.Int()%kinds]
		var n int
		if left <= 1 {
			n = 1
		} else {
			n = rnd2.Int()%(left-1) + 1
		}
		for i := 0; i < n; i++ {
			ret = append(ret, s)
		}
	}

	rnd2.Shuffle(length, func(i, j int) {
		ret[i], ret[j] = ret[j], ret[i]
	})

	return ret
}

func TestInitString2(t *testing.T) {
	for i, s := range initStrings2(1e7) {
		t.Log(i, s)
	}
}

func TestInitString3(t *testing.T) {
	for i, s := range initStrings3(1e3, 4, 16, 4) {
		t.Log(i, s)
	}
}

func TestStringKinds(t *testing.T) {
	arr := initStrings3(1e3, 4, 16, 500)
	res := StringKinds(arr)

	f := genStringKinds3()

	t.Log(res)
	t.Log(f(arr))
	t.Log(f(arr))
}

func genStringKinds3() func(arr []string) int {
	h := NewHashTable(1000, 5)
	return func(arr []string) int {
		h.Reset()
		return StringKinds3(h, arr)
	}
}

func BenchmarkStringKinds(b *testing.B) {
	arr := initStrings3(1e3, 4, 16, 500)
	res := StringKinds(arr)
	//b.Log(res)
	f := StringKinds
	//f := genStringKinds3()
	//f := StringKinds2
	for i := 0; i < b.N; i++ {
		if f(arr) != res {
			b.Fail()
		}
	}
}

func TestTrieTree_Set(t *testing.T) {
	tried := NewTrieTree()
	t.Log(tried.root)
	t.Log(tried.Set("a"))
	t.Log(tried.Set("a"))
	t.Log(tried.Set("ab"))
}

func TestPrimes(t *testing.T) {
	t.Log(Primes(100))
}

func TestBKDRHash(t *testing.T) {
	t.Log(BKDRHash("abc"))
	t.Log(BKDRHash("abcd"))
	t.Log(BKDRHash("a"))
}

func TestNewHashTable(t *testing.T) {
	h := NewHashTable(100000, 5)
	t.Log(h.Get("abc"))

	t.Log(h.Put("abc", 1234))
	t.Log(h.Get("abc"))

	t.Log(h.Put("abc", 56))
	t.Log(h.Get("abc"))

	t.Log(h.Delete("abc"))
	t.Log(h.Get("abc"))

	t.Log(h.Put("abc", 7))
	t.Log(h.Get("abc"))
}

func BenchmarkPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Primes(100000)
	}
}

func BenchmarkHashTable_Put(b *testing.B) {
	h := NewHashTable(100000, 5)
	for i := 0; i < b.N; i++ {
		h.Put("abcdefg", 10000)
	}
}

func BenchmarkBKDRHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BKDRHash("abcdefg")
	}
}