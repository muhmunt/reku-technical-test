package handler

import (
	"net/http"
	"reku-technical-test/src/entity"
	"reku-technical-test/src/helper"
	"reku-technical-test/src/shortener"

	"github.com/gin-gonic/gin"
)

type shortenerHandler struct {
	shortenerService shortener.Service
}

func NewShortener(shortenerService shortener.Service) *shortenerHandler {
	return &shortenerHandler{shortenerService}
}

func (h *shortenerHandler) Mount(shorten *gin.RouterGroup) {
	shorten.POST("/", h.ShortenURLHandler)
	shorten.GET("/:shortened", h.FetchURLHandler)
	shorten.GET("/", h.FetchAllURL)
}

func (h *shortenerHandler) ShortenURLHandler(c *gin.Context) {
	var request shortener.ShortRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		errors := helper.ValidationFormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed shortened url.", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	generateShortUrl := helper.GenerateShortenedURL(c.Request.Host, request.TargetURL)

	var short entity.Shortener
	short.ID = generateShortUrl
	short.TargetURL = request.TargetURL

	storeUrl, err := h.shortenerService.StoreShortUrl(c, short)

	if err != nil {
		response := helper.APIResponse("Failed store shortened", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := shortener.FormatShortened(storeUrl)
	response := helper.APIResponse("Shorted url successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *shortenerHandler) FetchURLHandler(c *gin.Context) {
	var request shortener.FetchRequest

	err := c.ShouldBindUri(&request)
	if err != nil {
		errors := helper.ValidationFormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed fetch shortened url.", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	getUrl, err := h.shortenerService.GetShortUrl(c, request)

	if err != nil {
		errors := helper.ValidationFormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Shortened url not found.", http.StatusNotFound, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := shortener.FormatShort(request.Shortened, getUrl)
	response := helper.APIResponse("Shorted url successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *shortenerHandler) FetchAllURL(c *gin.Context) {
	sort := c.Query("sort")

	getUrl, err := h.shortenerService.GetAllUrl(sort)

	if err != nil {
		errors := helper.ValidationFormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Shortened url not found.", http.StatusNotFound, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := shortener.FormatShorteneds(getUrl)
	response := helper.APIResponse("Fetch Shorted url successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
