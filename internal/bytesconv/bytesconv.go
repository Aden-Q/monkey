package bytesconv

import "unsafe"

func ByteToString(b byte) string {
	return unsafe.String(&b, 1)
}

func BytesToString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
