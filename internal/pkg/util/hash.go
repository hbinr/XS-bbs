package util

import (
	"crypto/md5"
	"encoding/hex"

	"xs.bbs/internal/pkg/constant/key"
)

// EncryptPassword 密码加密
func EncryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(key.MD5_SOLT))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
