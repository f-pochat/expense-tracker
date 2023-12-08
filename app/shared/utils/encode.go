package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func EncodePassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	sha1Hash := hex.EncodeToString(hash.Sum(nil))
	return sha1Hash
}
