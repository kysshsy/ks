package hello

import (
	"ks/internal/http/service/hello"

	"github.com/gin-gonic/gin"
)

func Hello(ctx *gin.Context) {
	JSON := map[string]interface{}{
		"data": hello.Hello(ctx),
	}
	ctx.AbortWithStatusJSON(200, JSON)
}
