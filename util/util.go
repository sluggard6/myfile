package util

import (
	"bytes"
	"crypto/sha256"
	"unsafe"

	"github.com/google/uuid"
)

func ShaString(s string) string {
	hash := sha256.Sum256([]byte(s))
	return ByteArrayToString(hash[:])
	// h := sha256.New()
	// h.Write([]byte(s))
	// return ""
	// return *(*string)(unsafe.Pointer(&h.Sum([]byte(s))))
	// return string(h.Sum(nil)[:])
}
func UUID() string {
	var buffer bytes.Buffer
	for _, chr := range uuid.New().String() {
		if "-" == string(chr) {
			continue
		}
		buffer.WriteString(string(chr))
	}
	return buffer.String()
}
func ByteArrayToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
