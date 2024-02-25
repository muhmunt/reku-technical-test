package server

import (
	"os"
	"reku-technical-test/src/web/routes"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Run(db *gorm.DB, redis *redis.Client) {
	router := gin.Default()

	port := os.Getenv("PORT")

	api := router.Group("/api/v1")

	routes.SetupShortenerRoutes(api, db, redis)

	router.Run(":" + port)
}
