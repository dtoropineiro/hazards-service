package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"hazards-emergencies-go/dto"
	"hazards-emergencies-go/services/composite"
	"net/http"
	_ "net/http"
)

func GetLastFires() gin.HandlerFunc {
	return func(c *gin.Context) {
		central123 := &composite.Central123FireEmergenciesCompositeService{}
		cvb := &composite.CvbFireEmergenciesCompositeService{}

		compositeService := composite.NewCompositeFireEmergencyService(central123, cvb)
		allFires, err := compositeService.GetAllFires()
		if err != nil {
			logrus.Error(err)
			return
		}
		fireResponseDto := dto.FireResponseDto{
			FireDto: allFires,
		}
		c.JSON(http.StatusOK, fireResponseDto)
	}
}
