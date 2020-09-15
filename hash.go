package exercises

import (
	"github.com/cespare/xxhash"
	"hash"
)

func BKDRHash(s string) int32 {
	seed := int32(131) // 31 131 113 13131 131313 etc..
	hash := int32(0)
	for _, c := range s {
		hash = hash*seed + c
	}
	return hash & 0x7fffffff
}

func BKDRSum64String(s string) uint64 {
	seed := uint64(131313) // 31 131 113 13131 131313 etc..
	hash := uint64(0)
	for _, c := range s {
		hash = hash*seed + uint64(c)
	}
	return hash & 0x7fffffffffffffff
}

func XXSum64String(s string) uint64 {
	return xxhash.Sum64String(s)
}

func Hash(h hash.Hash, bs []byte) []byte {
	h.Reset()
	h.Write(bs)
	return h.Sum(nil)
}
