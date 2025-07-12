package main

import (
	"github.com/gin-gonic/gin"
	"otus/internal/route"
)

// @title Serial Catalog API
// @version 1.0
// @description API for managing TV series catalog
// @contact.email evgenij.bx@gmail.com
// @host localhost:8080
// @BasePath /v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	route.Init(r)

	r.Run(":8080")
}
