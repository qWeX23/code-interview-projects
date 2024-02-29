package controllers

import (
	"joynet-assignment/config"
	"joynet-assignment/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReportsInterface interface {
	HandlePostReport(*gin.Context)
}

type reportsController struct {
	e *config.Env
}

func NewReportsController(env *config.Env) ReportsInterface {
	return &reportsController{
		e: env,
	}
}

type postReportRequest struct {
	Device     models.Device               `json:"devinfo"`
	Interfaces map[string]models.Interface `json:"intinfo"`
}
type postReportResponse struct{}

func (rc *reportsController) HandlePostReport(c *gin.Context) {
	payload := postReportRequest{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rc.e.Db.SaveDevice(payload.Device)

	for _, intfc := range payload.Interfaces {
		rc.e.Db.SaveInterfaces(intfc)
	}
	response := postReportResponse{}
	c.JSON(http.StatusCreated, response)
}
