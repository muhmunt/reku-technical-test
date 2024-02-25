package routes

import (
	"reku-technical-test/src/handler"
	"reku-technical-test/src/repository"
	"reku-technical-test/src/shortener"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SetupShortenerRoutes(r *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	routeShortener := r.Group("/shorten")

	shortenerRepository := repository.NewShortener(db, redis)
	shortenerService := shortener.NewService(shortenerRepository)

	shortenerHandler := handler.NewShortener(shortenerService)
	shortenerHandler.Mount(routeShortener)
}
