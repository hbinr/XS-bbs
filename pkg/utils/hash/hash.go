package hash

import (
	"crypto/md5"
	"encoding/hex"

	"xs.bbs/internal/pkg/constant"
)

// MD5String 密码加密
func MD5String(oPassword string) string {
	h := md5.New()
	h.Write([]byte(constant.KeyMySecret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
