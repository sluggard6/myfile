package util

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"strings"
	"unsafe"

	"github.com/google/uuid"
)

func ShaString(s string) string {
	hash := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", hash[:])
}
func ShaReader(reader io.Reader) ([]byte, error) {
	return SaveAndSha(reader, nil)
}

func SaveAndSha(reader io.Reader, file *os.File) ([]byte, error) {
	defer file.Close()
	hash := sha256.New()
	block := make([]byte, hash.BlockSize())
	for {
		i, err := reader.Read(block)
		if err != nil {
			if err != io.EOF {
				return nil, err
			} else {
				break
			}
		}
		if file != nil {
			file.Write(block[:i])
		}
		hash.Write(block[:i])
	}
	return hash.Sum(nil), nil
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
