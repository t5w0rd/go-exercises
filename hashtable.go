package exercises

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
				if i%prime == 0 {
					break
				}
			}
		}
	}
	return primes
}

func BKDRHash(s string) int32 {
	seed := int32(131) // 31 131 113 13131 131313 etc..
	hash := int32(0)
	for _, c := range s {
		hash = hash*seed + c
	}
	return hash & 0x7fffffff
}

type hashNode struct {
	key   string
	value interface{}
}

type emptyValue struct{}

type HashTable struct {
	data       []hashNode
	lineLength []int
	deep       int
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
		data:       data,
		lineLength: lineLength,
		deep:       0,
	}
}

func (h *HashTable) Put(key string, value interface{}) (exist bool, ok bool) {
	hash := int(BKDRHash(key))
	var base int
	var firstEmpty, target *hashNode
	firstEmptyDeep := h.deep + 1
	targetDeep := firstEmptyDeep
	for i := 0; i < h.deep; i++ {
		lineLength := h.lineLength[i]
		off := hash % lineLength
		where := &h.data[base+off]

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
		firstEmpty = &h.data[base+off]
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
		where := &h.data[base+off]

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
		where := &h.data[base+off]

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
			h.data[base+j].value = emptyValue{}
		}

		if deep := i + 1; deep == h.deep {
			break
		}
		base += lineLength
	}
}
