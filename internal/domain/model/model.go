package model

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

const sortKeyDelimiter = "--"

func generateHashKey(keys ...string) string {
	str := strings.Join(keys, sortKeyDelimiter)
	buf := sha256.Sum256([]byte(str))
	return hex.EncodeToString(buf[:])
}
