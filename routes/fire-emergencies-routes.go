package routes

import (
	"github.com/gin-gonic/gin"
	"hazards-emergencies-go/controllers"
)

func FireEmergenciesRoute(router *gin.Engine) {
	router.GET("api/v1/fires/", controllers.GetLastFires())
}
