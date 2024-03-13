package routes

import (
	"github.com/gin-gonic/gin"
	"hazards-emergencies-go/controllers"
)

func EarthquakesRoute(router *gin.Engine) {
	router.GET("api/v1/earthquakes/", controllers.GetLastEarthquakes())
}
