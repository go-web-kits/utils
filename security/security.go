package security

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(content []byte) string {
	writer := md5.New()
	writer.Write(content)
	return hex.EncodeToString(writer.Sum(nil))
}
