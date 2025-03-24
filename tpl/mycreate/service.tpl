package {{ .PackageName }}

import (
	"github.com/gin-gonic/gin"
)

// Handler 处理器
type Handler struct {
}

// LoginVerify 登录验证
// @return bool 是否验证
func (Handler) LoginVerify() bool {
	return true
}

// Handle 处理请求
// @param ctx *gin.Context 请求上下文
func (Handler) Handle(ctx *gin.Context) {
    params := bindParam(ctx)
	if params == nil || !params.IsLegal() {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "params error",
			"data": nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
    		"code": http.StatusOK,
    		"msg":  "method not implemented",
    		"data": nil,
    	})
	return
}
