package main

import (
	"log"

	"joynet-assignment/config"
	"joynet-assignment/controllers"

	"github.com/gin-gonic/gin"
	ginlogrus "github.com/toorop/gin-logrus"
)

var env config.Env

func init() {
	env = config.NewEnv()
	env.Db.Migrate()
}

func main() {
	router := gin.Default()
	router.Use(ginlogrus.Logger(env.Log))

	router.GET("/devices", controllers.NewDevicesController(&env).HandleGetDevices)
	router.GET("/interfaces", controllers.NewInterfacesController(&env).HandleGetInterfaces)
	router.POST("/reports", controllers.NewReportsController(&env).HandlePostReport)
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
