package ginx

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func QueryInt(param string, c *gin.Context) (intVar int64, err error) {
	intStr := c.Query(param)
	if intVar, err = strconv.ParseInt(intStr, 10, 64); err != nil {
		zap.L().Error("strconv.Atoi(intStr) 异常", zap.Error(err))
		return
	}
	return
}
