package exercises

import (
	"bytes"
	"errors"
	"strings"
	"sync"
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

func GCD2(a, b int) int {
	if a == 0 {
		return b
	}
	return GCD2(b % a, a)
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
	for i := n-1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
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
		j := i
		k := arr[i]
		for ; j > 0 && arr[j-1] > k; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = k
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

func QuickSort2(arr []int) []int {
	n := len(arr)
	if n <= 100 {
		return InsertSort(arr)
	}

	p := arr[0]
	i := 0
	j := n - 1
LOOP:
	for {
		for ; p <= arr[j]; j-- {
			if i == j {
				break LOOP
			}
		}

		arr[i] = arr[j]
		for ; arr[i] <= p; i++ {
			if i == j {
				break LOOP
			}
		}
		arr[j] = arr[i]
	}
	arr[i] = p

	QuickSort2(arr[:i])
	QuickSort2(arr[i+1:])
	return arr
}

func goQuickSort(arr []int, wg *sync.WaitGroup) {
	defer wg.Done()

	n := len(arr)
	if n <= 100 {
		InsertSort(arr)
		return
	}

	p := arr[0]
	i := 0
	j := n - 1
LOOP:
	for {
		for ; p <= arr[j]; j-- {
			if i == j {
				break LOOP
			}
		}

		arr[i] = arr[j]
		for ; arr[i] <= p; i++ {
			if i == j {
				break LOOP
			}
		}
		arr[j] = arr[i]
	}
	arr[i] = p

	wg.Add(2)
	go goQuickSort(arr[:i], wg)
	go goQuickSort(arr[i+1:], wg)
}

func QuickSort3(arr []int) []int {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go goQuickSort(arr, wg)
	wg.Wait()

	return arr
}

func shiftDown(arr []int, p int) {
	n := len(arr)
	l := (p << 1) + 1
	if l >= n {
		return
	}
	r := l + 1
	largest := p
	if arr[l] > arr[largest] {
		largest = l
	}
	if r < n && arr[r] > arr[largest] {
		largest = r
	}
	if largest != p {
		arr[p], arr[largest] = arr[largest], arr[p]
		shiftDown(arr, largest)
	}
}

func heapBubble(a []int, pos int) {
	if pos == 0 {
		return
	}
	p := (pos - 1) >> 1
	if a[pos] > a[p] {
		a[pos], a[p] = a[p], a[pos]
		heapBubble(a, p)
	}
}

func MaxHeap(arr []int) int {
	heap := make([]int, 0, len(arr))
	for i, e := range arr {
		heap = append(heap, e)
		heapBubble(heap, i)
	}
	return heap[0]
}

func HeapSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	// build heap
	for i := 2; i < n; i++ {
		heapBubble(arr[:i], i - 1)
	}
	println(arr)

	for i := n - 1; i >= 1; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		shiftDown(arr[:i], 0)
	}
	println(arr)
	return arr
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

type trieTreeNode struct {
	children [26]*trieTreeNode
	exists interface{}
}

type TrieTree struct {
	root *trieTreeNode
	size int
	pool []trieTreeNode
	poolUsed int
}

func NewTrieTree() *TrieTree {
	pool := make([]trieTreeNode, 100)
	return &TrieTree{
		root: &pool[0],
		pool: pool,
		poolUsed: 1,
	}
}

func (t *TrieTree) Set(s string) (exist bool) {
	where := &t.root
	for _, c := range s {
		index := c - 'a'
		where = &(*where).children[index]
		if *where == nil {
			*where = &trieTreeNode{}
		}
	}

	if (*where).exists == struct{}{} {
		return true
	}

	(*where).exists = struct{}{}
	t.size++
	return false
}

func (t *TrieTree) Set2(s string) (exist bool) {
	where := &t.root
	for _, c := range s {
		index := c - 'a'
		where = &(*where).children[index]
		if *where == nil {
			*where = &t.pool[t.poolUsed]
			t.poolUsed++
		}
	}

	if (*where).exists == struct{}{} {
		return true
	}

	(*where).exists = struct{}{}
	return false
}

func (t *TrieTree) Get(s string) (exist bool) {
	where := t.root
	for _, c := range s {
		index := c - 'a'
		where = where.children[index]
		if where == nil {
			return false
		}
	}
	return where.exists == struct{}{}
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

func Primes(n int) (primes []int) {
	isNotPrime := make([]bool, n+1)
	//primes = make([]int, 0, n+1)
	for i := 2; i <= n; i++ {
		if isNotPrime[i] == false {
			primes = append(primes, i)
		}
		for _, prime := range primes {
			if cur := i * prime; cur <= n {
				isNotPrime[cur] = true
				if i % prime == 0 {
					break
				}
			}
		}
	}
	return primes
}

func Primes2(n int) (primes []int) {
	isNotPrime := make([]bool, n+1)
	primes = make([]int, n+1)
	primesNum := 0
	for i := 2; i <= n; i++ {
		if isNotPrime[i] == false {
			primes[primesNum] = i
			primesNum++
		}
		for j := 0; j < primesNum; j++ {
			prime := primes[j]
			if cur := i * prime; cur <= n {
				isNotPrime[cur] = true
				if i % prime == 0 {
					break
				}
			}
		}
	}
	return primes
}

func BKDRHash(s string) int32 {
	seed := int32(131)  // 31 131 113 13131 131313 etc..
	hash := int32(0)
	for _, c := range s {
		hash = hash * seed + c
	}
	return hash & 0x7fffffff
}

type hashNode struct {
	key string
	value interface{}
}

type emptyValue struct{}

type HashTable struct {
	data []hashNode
	lineLength []int
	deep int
}

func NewHashTable(firstLineMaxLength int, lines int) *HashTable {
	primes := Primes(firstLineMaxLength)
	lineLength := make([]int, lines)
	size := 0
	for i := 0; i < lines; i++ {
		l := len(primes) - i - 1
		size += l
		lineLength[i] = l

	}
	data := make([]hashNode, size)
	for i := 0; i < size; i++ {
		data[i].value = emptyValue{}
	}

	return &HashTable{
		data: data,
		lineLength: lineLength,
		deep: 0,
	}
}

func (h *HashTable) Put(key string, value interface{}) (exist bool, ok bool) {
	hash := int(BKDRHash(key))
	var base int
	var firstEmpty, target *hashNode
	firstEmptyDeep := h.deep + 1
	targetDeep := firstEmptyDeep
	for i := 0; i < h.deep; i++  {
		lineLength := h.lineLength[i]
		off := hash % lineLength
		where := &h.data[base + off]

		if _, isEmpty := where.value.(emptyValue); isEmpty {
			if firstEmpty == nil {
				firstEmpty = where
				firstEmptyDeep = i + 1
			}
		} else {
			if where.key == key {
				target = where
				targetDeep = i + 1
				break
			}
		}

		base += lineLength
	}

	if target != nil {
		if targetDeep > firstEmptyDeep {
			firstEmpty.key = key
			firstEmpty.value = value
			target.value = emptyValue{}
		} else {
			target.value = value
		}
		return true, true
	}

	if firstEmpty != nil {
		if firstEmptyDeep > h.deep {
			h.deep = firstEmptyDeep
		}
		firstEmpty.key = key
		firstEmpty.value = value
		return false, true
	}

	if h.deep < len(h.lineLength) {
		off := hash % h.lineLength[h.deep]
		h.deep++
		firstEmpty = &h.data[base + off]
		firstEmpty.key = key
		firstEmpty.value = value
		return false, true
	}

	return false, false
}

func (h *HashTable) Get(key string) (value interface{}, ok bool) {
	hash := int(BKDRHash(key))
	var base int
	for i := 0; i < h.deep; i++ {
		lineLength := h.lineLength[i]
		off := hash % lineLength
		where := &h.data[base + off]

		if where.key == key {
			if _, ok := where.value.(emptyValue); !ok {
				return where.value, true
			}
		}

		base += lineLength
	}

	return nil, false
}

func (h *HashTable) Delete(key string) bool {
	hash := int(BKDRHash(key))
	var base int
	for _, lineLength := range h.lineLength {
		off := hash % lineLength
		where := &h.data[base + off]

		if _, ok := where.value.(emptyValue); ok {
			return false
		}

		if where.key == key {
			where.value = emptyValue{}
			return true
		}

		base += lineLength
	}
	return false
}

func (h *HashTable) Reset() {
	var base int
	for i, lineLength := range h.lineLength {
		for j := 0; j < lineLength; j++ {
			h.data[base + j].value = emptyValue{}
		}

		if deep := i + 1; deep == h.deep {
			break
		}
		base += lineLength
	}
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

