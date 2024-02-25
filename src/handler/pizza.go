package handler

import (
	"fmt"
	"net/http"
	"reku-technical-test/src/chef"
	"reku-technical-test/src/helper"
	"reku-technical-test/src/menu"
	"reku-technical-test/src/order"

	"github.com/gin-gonic/gin"
)

type pizzaHandler struct {
	menuService  menu.Service
	chefService  chef.Service
	orderService order.Service
}

func NewPizza(menuService menu.Service, chefService chef.Service, orderService order.Service) *pizzaHandler {
	return &pizzaHandler{menuService, chefService, orderService}
}

func (h *pizzaHandler) Mount(pizza *gin.RouterGroup) {
	pizza.GET("/menus", h.FetchMenusHandler)
	pizza.POST("/chef", h.StoreChef)
	pizza.POST("/order", h.PlaceOrder)
}

func (h *pizzaHandler) FetchMenusHandler(c *gin.Context) {

	allMenu, err := h.menuService.GetAllMenu()

	if err != nil {
		response := helper.APIResponse("Failed fetch menus", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := menu.FormatMenus(allMenu)
	response := helper.APIResponse("Fetch menu success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *pizzaHandler) StoreChef(c *gin.Context) {
	var request chef.CheftRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		errors := helper.ValidationFormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Request failed store chef.", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newChef, err := h.chefService.StoreChef(request)

	if err != nil {
		response := helper.APIResponse("Failed store chef", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := chef.FormatChef(newChef)
	response := helper.APIResponse("Store chef successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *pizzaHandler) PlaceOrder(c *gin.Context) {
	var request order.OrderRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		errors := helper.ValidationFormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Request failed place order.", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	getMenu, err := h.menuService.GetMenuById(request.PizzaID)

	if err != nil {
		response := helper.APIResponse("Failed get menu", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newOrder, err := h.orderService.PlaceOrder(getMenu)

	if err != nil {
		fmt.Println(err)
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := order.FormatOrder(newOrder)
	response := helper.APIResponse("Store order successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
