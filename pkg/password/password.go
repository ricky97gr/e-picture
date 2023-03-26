package password

import (
	"crypto/md5"
	"encoding/hex"
)

func GetPassword(passwd string) string {
	sum := md5.Sum([]byte(passwd))
	return hex.EncodeToString(sum[:])
}
