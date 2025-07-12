package route

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "otus/docs"
	"otus/internal/handler"
	"otus/internal/middleware"
)

func Init(router *gin.Engine) {

	v1 := router.Group("/v1")

	url := ginSwagger.URL("/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	v1.Use(func(c *gin.Context) { middleware.Process(c) })
	{
		v1.GET("/otus.serial.get/:id", handler.GetSerial)
		v1.POST("/otus.serial.add", handler.AddSerial)
		v1.GET("/otus.serial.list", handler.GetListSerial)
		v1.PUT("/otus.serial.update/:id", handler.UpdateSerial)
		v1.DELETE("/otus.serial.delete/:id", handler.DeleteSerial)

		v1.GET("/otus.season.get/:id", handler.GetSeason)
		v1.POST("/otus.season.add", handler.AddSeason)
		v1.GET("/otus.season.list", handler.GetListSeason)
		v1.PUT("/otus.season.update/:id", handler.UpdateSeason)
		v1.DELETE("/otus.season.delete/:id", handler.DeleteSeason)

		v1.GET("/otus.episode.get/:id", handler.GetEpisode)
		v1.POST("/otus.episode.add", handler.AddEpisode)
		v1.GET("/otus.episode.list", handler.GetListEpisode)
		v1.PUT("/otus.episode.update/:id", handler.UpdateEpisode)
		v1.DELETE("/otus.episode.delete/:id", handler.DeleteEpisode)
	}
}
