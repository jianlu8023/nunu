package {{ .PackageName }}

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

// Param 请求参数
type Param struct {
}

// IsLegal 参数是否有效
// @return bool 是否有效
func (p Param) IsLegal() bool {
	return false
}

// bindParam 获取参数
// @param ctx *gin.Context 请求上下文
// @return *Param 参数
func bindParam(ctx *gin.Context) *Param {
	var body = &Param{}

	if err := ctx.ShouldBindWith(body, binding.FormMultipart); err != nil {
		fmt.Printf("binding param error %v",err)
		return nil
	}
	return body
}
