package main

import (
	"log"
	"sukasaair/models"
	"sukasaair/repository/mongodb"
	"sukasaair/repository/redis"
	"sukasaair/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "sukasaair/docs"
)

var redisClient redis.RedisClientInterface

func main() {

	mongoDB, err := mongodb.NewMongoDB()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	log.Println("✅ MongoDB connected")

	redisClient = redis.NewRedisClient("localhost:6379", "", 0)
	log.Println("✅ Redis client initialized")

	models.SetDatabase(mongoDB)

	models.SetRedisClient(redisClient)

	r := gin.Default()

	routes.SetupAuthRouter(r)
	routes.SetupSeatRouter(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Println("🚀 Server running on port 8080")

	r.Run(":8080")
}
