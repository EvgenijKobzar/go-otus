package main

import (
	"github.com/gin-gonic/gin"
	"otus/internal/route"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	route.Init(r)

	r.Run(":8080")
}
