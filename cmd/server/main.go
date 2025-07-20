package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
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
	//repo := mongo.NewRepository[*catalog.Serial]()
	//ctx := context.Background()
	//
	//// Проверка соединения
	//_, err := repo.СlientRedis.Ping(ctx).Result()
	//if err != nil {
	//	fmt.Println("Ошибка подключения к Redis:", err)
	//	return
	//}

	// Кеширование данных
	//err = repo.СlientRedis.Set(ctx, "username", "john_doe3", 0).Err()
	//if err != nil {
	//	fmt.Println("Ошибка при кешировании данных:", err)
	//	return
	//}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	route.Init(r)

	r.Run(":8080")
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
