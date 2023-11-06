package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func HashMD5(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}
