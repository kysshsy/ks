package main

import (
	"ks/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()

	router.RouterInit(engine)

	engine.Run(":80")
}
