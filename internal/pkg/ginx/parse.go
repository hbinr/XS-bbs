package ginx

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func QueryInt(param string, c *gin.Context) (intVar int, err error) {
	intStr := c.Query(param)
	if intVar, err = strconv.Atoi(intStr); err != nil {
		zap.L().Error("strconv.Atoi(intStr) 异常", zap.Error(err))
		return
	}
	return
}
