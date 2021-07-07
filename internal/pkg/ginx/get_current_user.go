package ginx

import (
	"github.com/gin-gonic/gin"
	"xs.bbs/internal/pkg/constant/e"
	"xs.bbs/internal/pkg/constant/key"
)

// GetCurrentUserID 获取当前登录的用户ID
func GetCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(key.KeyCtxUserID)
	if !ok {
		err = e.ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = e.ErrorUserNotLogin
		return
	}
	return
}
