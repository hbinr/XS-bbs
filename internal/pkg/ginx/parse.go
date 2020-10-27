package ginx

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func QueryInt(param string, c *gin.Context) (intVar int, err error) {
	intStr := c.Query(param)
	if intVar, err = strconv.Atoi(intStr); err != nil {
		return
	}
	return
}
