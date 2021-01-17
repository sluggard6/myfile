package util

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"unsafe"

	"github.com/google/uuid"
)

func ShaString(s string) string {
	hash := sha256.Sum256([]byte(s))
	return base64.StdEncoding.EncodeToString(hash[:])
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

func isBlankString(s string) bool {
	return strings.Trim(s, " ") == ""
}
