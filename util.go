package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
)

func Gen(originalStr, prefix string, point, length int) (string, string, error) {
	md5Str := Md5Encrypt(originalStr)

	md5Len := len(md5Str)
	start := point - 1
	end := start + length
	if start >= md5Len || end > md5Len {
		return "", "", errors.New("指定的起点或长度太长了")
	}

	result := prefix + md5Str[start:end]
	return md5Str, result, nil
}

func Md5Encrypt(originalStr string) string {
	originalBytes := []byte(originalStr)
	md5BytesArray := md5.Sum(originalBytes)
	md5Bytes := md5BytesArray[:]
	md5Str := hex.EncodeToString(md5Bytes)
	return md5Str
}
