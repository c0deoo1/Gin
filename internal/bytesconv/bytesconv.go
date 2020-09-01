package bytesconv

import (
	"reflect"
	"unsafe"
)

// StringToBytes converts string to byte slice without a memory allocation.
// 为了提升性能，这里利用了string和slice的内部接口，避免了对string的data的拷贝，仅仅改动了指针的指向。
func StringToBytes(s string) (b []byte) {
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Len
	return b
}

// BytesToString converts byte slice to string without a memory allocation.
// 避免内存拷贝
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
