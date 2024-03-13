package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hazards-emergencies-go/clients"
	"hazards-emergencies-go/configs"
	"hazards-emergencies-go/routes"
)

func main() {

	configs.LoadConfiguration()
	configs.Connect()
	_, err := clients.GetLastFires()
	if err != nil {
		fmt.Println("Error al obtener los últimos incendios")
	}
	//configs.Connect()
	router := gin.Default()
	routes.FireEmergenciesRoute(router)
	routes.EarthquakesRoute(router)

	err = router.Run(":8080")
	if err != nil {
		return
	}
}
