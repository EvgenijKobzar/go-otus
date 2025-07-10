package route

import (
	"github.com/gin-gonic/gin"
	"otus/internal/controller"
	"otus/internal/middleware"
	"otus/internal/model/catalog"
)

// http://localhost:8080/v1/otus.serial.get?id=1
// http://localhost:8080/v1/otus.serial.add?fields[title]=test4&fields[sort]=4
// http://localhost:8080/v1/otus.serial.list
// http://localhost:8080/v1/otus.serial.update?id=2&fields[title]=123
// http://localhost:8080/v1/otus.serial.delete?id=2

// http://localhost:8080/v1/otus.episode.get?id=1
// http://localhost:8080/v1/otus.episode.add?fields[title]=test4
// http://localhost:8080/v1/otus.episode.list
// http://localhost:8080/v1/otus.episode.update?id=1&fields[serialId]=1&fields[seasonId]=1&fields[title]=%D0%9F%D1%80%D0%B8%D0%B7%D0%BD%D0%B0%D0%BD%D0%B8%D1%8F
// http://localhost:8080/v1/otus.episode.delete?id=17

// http://localhost:8080/v1/otus.season.get?id=1
// http://localhost:8080/v1/otus.season.add?fields[title]=test4
// http://localhost:8080/v1/otus.season.list
// http://localhost:8080/v1/otus.season.update?id=3&fields[title]=123
// http://localhost:8080/v1/otus.season.delete?id=3

func Init(router *gin.Engine) {

	v1 := router.Group("/v1")
	v1.Use(func(c *gin.Context) { middleware.Process(c) })
	{
		v1.GET("/otus.serial.get", func(c *gin.Context) { controller.GetAction[*catalog.Serial](c) })
		v1.GET("/otus.serial.add", func(c *gin.Context) { controller.AddAction[*catalog.Serial](c) })
		v1.GET("/otus.serial.list", func(c *gin.Context) { controller.GetListAction[*catalog.Serial](c) })
		v1.GET("/otus.serial.update", func(c *gin.Context) { controller.UpdateAction[*catalog.Serial](c) })
		v1.GET("/otus.serial.delete", func(c *gin.Context) { controller.DeleteAction[*catalog.Serial](c) })

		v1.GET("/otus.season.get", func(c *gin.Context) { controller.GetAction[*catalog.Season](c) })
		v1.GET("/otus.season.add", func(c *gin.Context) { controller.AddAction[*catalog.Season](c) })
		v1.GET("/otus.season.list", func(c *gin.Context) { controller.GetListAction[*catalog.Season](c) })
		v1.GET("/otus.season.update", func(c *gin.Context) { controller.UpdateAction[*catalog.Season](c) })
		v1.GET("/otus.season.delete", func(c *gin.Context) { controller.DeleteAction[*catalog.Season](c) })

		v1.GET("/otus.episode.get", func(c *gin.Context) { controller.GetAction[*catalog.Episode](c) })
		v1.GET("/otus.episode.add", func(c *gin.Context) { controller.AddAction[*catalog.Episode](c) })
		v1.GET("/otus.episode.list", func(c *gin.Context) { controller.GetListAction[*catalog.Episode](c) })
		v1.GET("/otus.episode.update", func(c *gin.Context) { controller.UpdateAction[*catalog.Episode](c) })
		v1.GET("/otus.episode.delete", func(c *gin.Context) { controller.DeleteAction[*catalog.Episode](c) })
	}
}
