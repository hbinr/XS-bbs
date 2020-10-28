package ginx

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
	"go.uber.org/zap"
	"xs.bbs/internal/pkg/constant/e"
)

// BindAndValid 参数绑定及校验
func BindAndValid(c *gin.Context, params interface{}) string {
	if err := c.ShouldBind(params); err != nil {
		zap.L().Error(e.CODE_INVALID_PARAMS.Msg(), zap.Error(err))
		return fmt.Sprintf("%s-err:%s", e.CODE_INVALID_PARAMS.Msg(), err.Error())
	}
	if err := gvalid.CheckStruct(params, nil); err != nil {
		zap.L().Error(e.CODE_INVALID_PARAMS.Msg(), zap.Error(err))
		return err.FirstString()
	}
	return ""
}
