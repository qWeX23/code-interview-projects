package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	conStr := "mongodb://root:example@mongo:27017"

	db := newMongoDb(conStr)
	ps := NewPatientService(db)

	r := gin.Default()
	r.GET("/patient/:id", ps.GetPatientHandler)
	r.POST("/patient", ps.PostPatientHandler)
	r.PUT("/patient/:id", ps.PutPatientHandler)
	r.DELETE("/patient/:id", ps.DeletePatientHandler)
	r.GET("/search", ps.SearchPatientHandler)
	r.Run(":8080")
}
