package route

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "otus/docs"
	"otus/internal/core"
	"otus/internal/handler"
	"otus/internal/middleware"
	"otus/internal/model"
	"otus/internal/model/catalog"
	"otus/internal/repository/postgres/sqlc"
)

func Init(router *gin.Engine) {

	v1 := router.Group("/v1")

	url := ginSwagger.URL("/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	v1.Use(func(c *gin.Context) { middleware.Process(c) })
	{
		v1.GET("/otus.serial.get/:id", func(context *gin.Context) { getHandler[*catalog.Serial]().GetSerial(context) })
		v1.POST("/otus.serial.add", middleware.Auth, func(context *gin.Context) { getHandler[*catalog.Serial]().AddSerial(context) })
		v1.GET("/otus.serial.list", func(context *gin.Context) { getHandler[*catalog.Serial]().GetListSerial(context) })
		v1.PUT("/otus.serial.update/:id", middleware.Auth, func(context *gin.Context) { getHandler[*catalog.Serial]().UpdateSerial(context) })
		v1.DELETE("/otus.serial.delete/:id", middleware.Auth, func(context *gin.Context) { getHandler[*catalog.Serial]().DeleteSerial(context) })

		v1.GET("/otus.season.get/:id", func(context *gin.Context) { getHandler[*catalog.Season]().GetSeason(context) })
		v1.POST("/otus.season.add", middleware.Auth, func(context *gin.Context) { getHandler[*catalog.Season]().AddSeason(context) })
		v1.GET("/otus.season.list", func(context *gin.Context) { getHandler[*catalog.Season]().GetListSeason(context) })
		v1.PUT("/otus.season.update/:id", middleware.Auth, func(context *gin.Context) { getHandler[*catalog.Season]().UpdateSeason(context) })
		v1.DELETE("/otus.season.delete/:id", middleware.Auth, func(context *gin.Context) { getHandler[*catalog.Season]().DeleteSeason(context) })

		v1.GET("/otus.episode.get/:id", func(context *gin.Context) { getHandler[*catalog.Episode]().GetEpisode(context) })
		v1.POST("/otus.episode.add", middleware.Auth, func(context *gin.Context) { getHandler[*catalog.Episode]().AddEpisode(context) })
		v1.GET("/otus.episode.list", func(context *gin.Context) { getHandler[*catalog.Episode]().GetListEpisode(context) })
		v1.PUT("/otus.episode.update/:id", middleware.Auth, func(context *gin.Context) { getHandler[*catalog.Episode]().UpdateEpisode(context) })
		v1.DELETE("/otus.episode.delete/:id", middleware.Auth, func(context *gin.Context) { getHandler[*catalog.Episode]().DeleteEpisode(context) })

		v1.GET("/otus.account.get/:id", func(context *gin.Context) { getHandler[*model.Account]().GetAccount(context) })
		v1.GET("/otus.account.list", func(context *gin.Context) { getHandler[*model.Account]().GetListAccount(context) })
		v1.DELETE("/otus.account.delete/:id", middleware.Auth, func(context *gin.Context) { getHandler[*model.Account]().DeleteAccount(context) })
		v1.POST("/otus.account.register", handler.RegisterAccount)
		v1.POST("/otus.account.login/", handler.LoginAccount)
	}
}

func getHandler[T catalog.HasId]() *handler.Handler[T] {
	repo := sqlc.NewRepository[T]()
	service := core.New(repo)
	return handler.New(service)
}
