package {{ .PackageName }}

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

// Dispatcher 请求分发
func Dispatcher(ctx *gin.Context) {
    fmt.Printf("dispatcher ctrl ...")

	contextHandler := &Handler{}

	// 判断是否需要进行 session 验证
	if contextHandler.LoginVerify() {
		// 自定义处理session信息

		// 业务处理
		contextHandler.Handle(ctx)
	} else {
		// 执行具体的内容
		contextHandler.Handle(ctx)
	}
}
