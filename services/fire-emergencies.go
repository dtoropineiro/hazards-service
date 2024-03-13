package services

import (
	"github.com/gin-gonic/gin"
	"hazards-emergencies-go/clients"
	"net/http"
)

func GetLastFires() gin.HandlerFunc {
	return func(c *gin.Context) {
		fires, err := clients.GetLastFires()
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}
		c.JSON(http.StatusOK, fires)
	}
}
