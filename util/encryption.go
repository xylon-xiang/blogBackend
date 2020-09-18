package util

import (
	"crypto"
	"fmt"
)

func EncryptPassword(password string) string {
	h := crypto.SHA256.New()
	h.Write([]byte(password))
	byteStream := h.Sum(nil)
	return fmt.Sprintf("%x", byteStream)
}
