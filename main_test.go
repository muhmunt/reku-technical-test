package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reku-technical-test/src/entity"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestStoreShorten(t *testing.T) {
	short := entity.Shortener{
		TargetURL: "localhost:8082/me/test",
	}

	orderJSON, err := json.Marshal(short)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/shorten", bytes.NewBuffer(orderJSON))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(rr)
	c.Request = req

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
}

func TestGetUrl(t *testing.T) {
	req, err := http.NewRequest("GET", "/pizza/shorten/bda7d7", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(rr)
	c.Request = req

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
}

func TestGetAllUrl(t *testing.T) {
	req, err := http.NewRequest("GET", "/pizza/shorten", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(rr)
	c.Request = req

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
}

func TestPlaceOrder(t *testing.T) {
	order := entity.Order{
		Pizza: "Pizza Cheese",
	}

	orderJSON, err := json.Marshal(order)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/pizza/orders", bytes.NewBuffer(orderJSON))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(rr)
	c.Request = req

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
}

func TestStoreChef(t *testing.T) {
	chef := entity.Chef{
		Name: "John",
	}

	orderJSON, err := json.Marshal(chef)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/pizza/chef", bytes.NewBuffer(orderJSON))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(rr)
	c.Request = req

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
}

func TestGetMenu(t *testing.T) {
	req, err := http.NewRequest("GET", "/pizza/menus", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(rr)
	c.Request = req

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
}
