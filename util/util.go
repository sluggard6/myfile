package util

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
	"strings"
	"unsafe"

	"github.com/google/uuid"
)

func ShaString(s string) string {
	hash := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", hash[:])
}
func ShaReader(reader io.Reader) (int, []byte, error) {
	hash := sha256.New()
	block := make([]byte, hash.BlockSize())
	var size int
	for {
		i, err := reader.Read(block)
		if err != nil {
			if err != io.EOF {
				return 0, nil, err
			} else {
				break
			}
		}
		hash.Write(block[0:i])
		size += i
	}
	return size, hash.Sum(nil), nil
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
func Md5String(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
