package routes

import (
	"reku-technical-test/src/chef"
	"reku-technical-test/src/handler"
	"reku-technical-test/src/menu"
	"reku-technical-test/src/order"
	"reku-technical-test/src/repository"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SetupPizzaRoutes(r *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	routeShortener := r.Group("/pizza")

	// menu
	menuRepository := repository.NewMenu(db, redis)
	menuService := menu.NewService(menuRepository)
	// chef
	chefRepository := repository.NewChef(db, redis)
	chefService := chef.NewService(chefRepository)
	// order
	orderRepository := repository.NewOrder(db, redis)
	orderService := order.NewService(orderRepository, chefRepository)

	shortenerHandler := handler.NewPizza(menuService, chefService, orderService)
	shortenerHandler.Mount(routeShortener)
}
