package route

import (
	"github.com/gin-gonic/gin"
	"otus/internal/controller"
	"otus/internal/middleware"
)

func Init(router *gin.Engine) {

	v1 := router.Group("/v1")
	{
		v1.GET("/otus.serial.get", middleware.Process, controller.GetSerialAction)
		v1.GET("/otus.serial.add", middleware.Process, controller.AddSerialAction)
		v1.GET("/otus.serial.list", middleware.Process, controller.GetListSerialAction)
		v1.GET("/otus.serial.update", middleware.Process, controller.UpdateSerialAction)
		v1.GET("/otus.serial.delete", middleware.Process, controller.DeleteSerialAction)
	}
}
