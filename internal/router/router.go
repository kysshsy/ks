package router

import (
	"ks/internal/http/controller/hello"

	"github.com/gin-gonic/gin"
)

func RouterInit(e *gin.Engine) {
	e.GET("/hello", hello.Hello)
}
