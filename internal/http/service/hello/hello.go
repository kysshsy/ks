package hello

import "github.com/gin-gonic/gin"

func Hello(ctx *gin.Context) string {
	return "Hello, world"
}
