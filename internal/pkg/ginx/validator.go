package ginx

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
	"xs.bbs/internal/pkg/constant/e"
)

func BindAndValid(c *gin.Context, params interface{}) string {
	if err := c.ShouldBind(params); err != nil {
		return fmt.Sprintf("%s-err:%s", e.GetMsg(e.INVALID_PARAMS), err.Error())
	}
	if err := gvalid.CheckStruct(params, nil); err != nil {
		return err.FirstString()
	}
	return ""
}
