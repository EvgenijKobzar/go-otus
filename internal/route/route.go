package route

import (
	"github.com/gin-gonic/gin"
	"otus/internal/controller"
	"otus/internal/middleware"
	"otus/internal/model/catalog"
)

func Init(router *gin.Engine) {

	v1 := router.Group("/v1")
	v1.Use(func(c *gin.Context) { middleware.Process(c) })
	{
		v1.GET("/otus.serial.get/:id", func(c *gin.Context) { controller.GetAction[*catalog.Serial](c) })
		v1.POST("/otus.serial.add", func(c *gin.Context) { controller.AddAction[*catalog.Serial](c) })
		v1.GET("/otus.serial.list", func(c *gin.Context) { controller.GetListAction[*catalog.Serial](c) })
		v1.PUT("/otus.serial.update/:id", func(c *gin.Context) { controller.UpdateAction[*catalog.Serial](c) })
		v1.DELETE("/otus.serial.delete/:id", func(c *gin.Context) { controller.DeleteAction[*catalog.Serial](c) })

		v1.GET("/otus.season.get/:id", func(c *gin.Context) { controller.GetAction[*catalog.Season](c) })
		v1.POST("/otus.season.add", func(c *gin.Context) { controller.AddAction[*catalog.Season](c) })
		v1.GET("/otus.season.list", func(c *gin.Context) { controller.GetListAction[*catalog.Season](c) })
		v1.PUT("/otus.season.update/:id", func(c *gin.Context) { controller.UpdateAction[*catalog.Season](c) })
		v1.DELETE("/otus.season.delete/:id", func(c *gin.Context) { controller.DeleteAction[*catalog.Season](c) })

		v1.GET("/otus.episode.get/:id", func(c *gin.Context) { controller.GetAction[*catalog.Episode](c) })
		v1.POST("/otus.episode.add", func(c *gin.Context) { controller.AddAction[*catalog.Episode](c) })
		v1.GET("/otus.episode.list", func(c *gin.Context) { controller.GetListAction[*catalog.Episode](c) })
		v1.PUT("/otus.episode.update/:id", func(c *gin.Context) { controller.UpdateAction[*catalog.Episode](c) })
		v1.DELETE("/otus.episode.delete/:id", func(c *gin.Context) { controller.DeleteAction[*catalog.Episode](c) })
	}
}
