package controllers

import (
	"github.com/gin-gonic/gin"
	"hazards-emergencies-go/services"
	"net/http"
)

func GetLastEarthquakes() gin.HandlerFunc {
	return func(c *gin.Context) {
		earthquakeResponseDto := services.GetEarthquakes()
		c.JSON(http.StatusOK, earthquakeResponseDto)
	}
}
