package util

import (
	"crypto/sha256"
)

func shaToString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return ""
	// return *(*string)(unsafe.Pointer(&h.Sum([]byte(s))))
	// return string(h.Sum(nil)[:])
}
