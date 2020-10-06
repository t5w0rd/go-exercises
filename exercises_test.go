package exercises

import (
	"hash/fnv"
	"math/rand"
	"reflect"
	"runtime"
	"sync"
	"testing"
	"time"
)

var rnd *rand.Rand

func init() {
	rnd = rand.New(rand.NewSource(1234567))
}

func Test(t *testing.T) {
	a := new([]byte)
	t.Log(a)
	var b, bb []int
	c := make([]int, 0)
	cc := make([]int, 0)
	d := []int{}
	t.Log(
		len(b),
		len(c),
		reflect.DeepEqual(b, c),
		reflect.DeepEqual(c, d),
		reflect.DeepEqual(b, bb),
		reflect.DeepEqual(c, cc),
	)

	arr := []int{3, 0, 5, 20, 8, 9, 28}
	t.Log(HeapSort(arr))
}

func initStrings() (string, string) {
	var bufA, bufB []byte
	for i := 0; i < 10000000; i++ {
		c := byte(rnd.Int()%26 + int('a'))
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

// 554, 49844, 51939, 41017
func BenchmarkStrcat1(b *testing.B) {
	arr := initStrings2(1000, 16, 32)
	f := Strcat4
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(arr)
	}
}

func BenchmarkConvertString(b *testing.B) {
	s, _ := initStrings()
	var buf []byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf = []byte(s)
	}
	b.StopTimer()
	b.Log(len(buf))
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
	f := GCD
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(293284, 104006)
	}
}

func BenchmarkLCM(b *testing.B) {
	f := LCM
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(293284, 104006)
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

func checkMaxHeap(heap []int) bool {
	n := len(heap)
	m := n >> 1
	for i := 0; i < m; i++ {
		l := (i << 1) + 1
		if heap[l] > heap[i] {
			return false
		}
		if r := l + 1; r < n && heap[r] > heap[i] {
			return false
		}
	}
	return true
}

func TestMaxHeap(t *testing.T) {
	//arr := []int{3, 0, 5, 20, 8, 9, 28}
	arr := randArrayWithSeed(10000, 100)
	f := MaxHeap2
	f(arr)
	if !checkMaxHeap(arr) {
		t.Fail()
	}
}

func BenchmarkMaxHeap(b *testing.B) {
	arr := randArrayWithSeed(10000, 100)
	f := MaxHeap
	b.ResetTimer()
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		heap := make([]int, len(arr))
		copy(heap, arr)

		b.StartTimer()
		f(heap)
		b.StopTimer()

		if !checkMaxHeap(heap) {
			b.Fail()
		}
	}
}

func TestSort(t *testing.T) {
	arr := randArray(10000)
	f := BubbleSort
	f(arr)
	if !checkSort(arr) {
		t.Fail()
	}
}

func BenchmarkSort(b *testing.B) {
	//f := BubbleSort
	//f := SelectSort
	//f := InsertSort
	//f := HeapSort
	//f := QuickSort
	//f := QuickSort2
	//f := QuickSort3
	//f := GoSort
	f := GoStableSort

	b.ResetTimer()
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		arr := randArray(10000)

		b.StartTimer()
		f(arr)
		b.StopTimer()

		if !checkSort(arr) {
			b.Fail()
		}
	}
}

func initStrings2(length int, min, max int) []string {
	ret := make([]string, length)
	buf := make([]byte, max)
	const r = 'z' - 'a'
	for i := 0; i < length; i++ {
		n := rnd.Int()%(max-min) + min
		for j := 0; j < n; j++ {
			buf[j] = byte(rnd.Int31()%r + 'a')
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
		n := rnd.Int()%(maxLen-minLen) + minLen
		for j := 0; j < n; j++ {
			buf[j] = byte(rnd.Int31()%r + 'a')
		}
		ret[i] = string(buf[:n])
	}

	for left := length - len(ret); left > 0; left = length - len(ret) {
		s := ret[rnd.Int()%kinds]
		var n int
		if left <= 1 {
			n = 1
		} else {
			n = rnd.Int()%(left-1) + 1
		}
		for i := 0; i < n; i++ {
			ret = append(ret, s)
		}
	}

	rnd.Shuffle(length, func(i, j int) {
		ret[i], ret[j] = ret[j], ret[i]
	})

	return ret
}

func TestInitString2(t *testing.T) {
	for i, s := range initStrings2(1e7, 16, 32) {
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
	b.ResetTimer()
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
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Put("abcdefg", 10000)
	}
}

func BenchmarkBKDRHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BKDRHash("abcdefg")
	}
}

func initBinTree() *BinTreeNode {
	t := &BinTreeNode{
		2,
		&BinTreeNode{
			7,
			&BinTreeNode{2, nil, nil},
			&BinTreeNode{
				6,
				&BinTreeNode{5, nil, nil},
				&BinTreeNode{11, nil, nil},
			},
		},
		&BinTreeNode{
			5,
			nil,
			&BinTreeNode{
				9,
				&BinTreeNode{4, nil, nil},
				nil,
			},
		},
	}

	return t
}

func printNode(node *BinTreeNode) {
	println(node.data)
}

func TestLMR(t *testing.T) {
	tree := initBinTree()
	LMR(tree, printNode)
}

func TestLMR2(t *testing.T) {
	tree := initBinTree()
	LMR2(tree, printNode)
}

func TestLMR3(t *testing.T) {
	tree := initBinTree()
	LMR2(tree, printNode)
}

func BenchmarkLMR(b *testing.B) {
	tree := initBinTree()
	f := LMR
	for i := 0; i < b.N; i++ {
		f(tree, func(node *BinTreeNode) {})
	}
}

func TestMLR(t *testing.T) {
	tree := initBinTree()
	MLR(tree, printNode)
}

func BenchmarkRingIncrease(b *testing.B) {
	f := RingIncrease
	n := 0
	for i := 0; i < b.N; i++ {
		n = f(n, 10)
	}
}

func TestBS(t *testing.T) {
	tree := initBinTree()
	BS(tree, printNode)
}

func TestChangeSlice(t *testing.T) {
	arr := []int{5, 5, 5, 5, 5}
	ChangeSlice(arr)
	t.Log(arr)
}

func BenchmarkStringBytes(b *testing.B) {
	f := StringBytes
	var sum int
	arr := initStrings2(1, 10000, 20000)
	s := arr[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum += len(f(s))
	}
}

func BenchmarkBytesString(b *testing.B) {
	f := BytesString
	arr := initStrings2(1, 10000, 20000)
	bs := []byte(arr[0])
	var sum int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum += len(f(bs))
	}
}

func BenchmarkCopyString(b *testing.B) {
	f := CopyString
	arr := initStrings2(1, 10000, 20000)
	s := arr[0]
	buf := make([]byte, len(s))
	var sum int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum += len(f(s, buf))
	}
}

func BenchmarkCopyBytes(b *testing.B) {
	f := CopyBytes
	arr := initStrings2(1, 10000, 20000)
	s := arr[0]
	bs := []byte(s)
	buf := make([]byte, len(bs))
	var sum int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum += len(f(bs, buf))
	}

	fnv.New32()
}

func BenchmarkSum64String(b *testing.B) {
	//f := BKDRSum64String
	f := XXSum64String
	arr := initStrings2(1000, 10000, 20000)
	s := arr[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(s)
	}
}

func BenchmarkHash(b *testing.B) {
	arr := initStrings2(1000, 10000, 20000)
	bs := []byte(arr[0])
	b.ResetTimer()
	//h := md5.New()  // 2w5 ns
	//h := fnv.New128()  // 4w ns
	//h := fnv.New128()  // 4w ns
	//h := xxhash.New()  // 1344 ns
	h := fnv.New64() // 2w ns
	for i := 0; i < b.N; i++ {
		Hash(h, bs)
	}
}

func TestStringType(t *testing.T) {
	s := "nh 你好"
	t.Log(len(s))

	for i := 0; i < len(s); i++ {
		t.Log(i, reflect.TypeOf(s[0]), s[i])
	}

	for i, c := range s {
		t.Log(i, reflect.TypeOf(s[0]), s[i])
		t.Log(i, reflect.TypeOf(c), c)
	}
}

func TestChan(t *testing.T) {
	ch := make(chan int, 4)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	t.Log(len(ch))
	t.Log(<-ch)
	a, b := <-ch
	t.Log(a, b)
	t.Log(<-ch)
	t.Log(<-ch)
	ch <- 4 // panic
	t.Log(<-ch)
}

func TestSeqBefore(t *testing.T) {
	var seq1 uint8 = 1
	seq2 := seq1 + 128
	if !SeqBefore(seq1, seq2) {
		t.Fail()
	}
}

func TestConcurrencyAdd(t *testing.T) {
	f := ConcurrencyAdd
	threads := runtime.NumCPU() * 10
	const loops = 10000

	var x int64
	target := x + int64(threads)*loops
	f(&x, threads, loops)
	t.Log(x, target)
}

func BenchmarkConcurrencyAdd(b *testing.B) {
	f := ConcurrencyAdd2
	threads := runtime.NumCPU() * 1
	const loops = 10000

	var x int64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(&x, threads, loops)
	}
}

func TestRateLimiter_Acquire(t *testing.T) {
	threads := 10
	loops := 500000
	qps := loops * threads / 5
	tb := NewRateLimiter(qps)
	defer tb.Stop()

	tm := time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < loops; i++ {
				tb.Acquire()
			}
		}()
	}
	wg.Wait()

	realQPS := float64(loops*threads) / time.Now().Sub(tm).Seconds()
	t.Logf("%.1f/%d, %.1f%%", realQPS, qps, realQPS*100/float64(qps))
}

func BenchmarkRateLimiter_Acquire(b *testing.B) {
	b.Log("b.N:", b.N)
	tb := NewRateLimiter(10000)
	defer tb.Stop()
	for i := 0; i < b.N; i++ {
		tb.Acquire()
	}
}
