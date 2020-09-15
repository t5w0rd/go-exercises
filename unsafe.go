package exercises

import (
	"reflect"
	"unsafe"
)

func StringBytes(s string) (b []byte) {
	strHdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	sliceHdr := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sliceHdr.Data = strHdr.Data
	sliceHdr.Len = strHdr.Len
	sliceHdr.Cap = strHdr.Len
	return b
}

func StringBytes2(s string) (b []byte) {
	return []byte(s)
}

func BytesString(b []byte) (s string) {
	sliceHdr := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	strHdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	strHdr.Data = sliceHdr.Data
	strHdr.Len = sliceHdr.Len
	return s
}

func BytesString2(b []byte) (s string) {
	return string(b)
}

func CopyString(s string, buf []byte) []byte {
	b := StringBytes(s)
	copy(buf, b)
	return buf
}

func CopyString2(s string, buf []byte) []byte {
	copy(buf, s)
	return buf
}

func CopyBytes(bs []byte, buf []byte) []byte {
	copy(buf, bs)
	return buf
}

func CopyBytes2(bs []byte, buf []byte) []byte {
	buf = append(buf[0:], bs...)
	return buf
}

func CopyBytes3(bs []byte, buf []byte) []byte {
	for i, b := range bs {
		buf[i] = b
	}
	return buf
}
