package utils

import (
	"encoding/base64"
	"encoding/hex"
	"crypto/md5"
)

func Md5(s string) string {
	m := md5.New()
	m.Write([]byte(s))
	result := hex.EncodeToString(m.Sum(nil))
	return result
}

func Base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

/*
str := utils.Base64Encode([]byte("Hello, playground"))
 */

func Base64Decode(src string) (string, err) {
	code, err := base64.StdEncoding.DecodeString(src)
	return string(code), err
}

func APHash(key string) uint32 {
	var hash uint32
	for i, char := range key {
		if i&1 == 0 {
			hash ^= ((hash << 7) ^ uint32(char) ^ (hash >> 3))
		} else {
			hash ^= (^((hash << 11) ^ uint32(char) ^ (hash >> 5)))
		}
	}
	return (hash & 0x7FFFFFFF)
}
