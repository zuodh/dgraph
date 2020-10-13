package fbx

import (
	"unsafe"
)

// BufSize is the default buffer size for builders
const bufSize = 1024

// BytesToString represents a byte buffer as a string.
// The buffer should not be modified after this.
func BytesToString(buf []byte) string {
	return *(*string)(unsafe.Pointer(&buf))
}
