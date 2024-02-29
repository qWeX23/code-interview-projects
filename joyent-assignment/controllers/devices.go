package controllers

import (
	"joynet-assignment/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DevicesInterface interface {
	HandleGetDevices(c *gin.Context)
}
type devicesController struct {
	e *config.Env
}

func NewDevicesController(env *config.Env) DevicesInterface {
	return &devicesController{
		e: env,
	}
}

func (dc *devicesController) HandleGetDevices(c *gin.Context) {
	devices, err := dc.e.Db.GetDevices()
	if err != nil {
		dc.e.Log.Error(err)
		response := struct{}{}
		c.JSON(http.StatusInternalServerError, response)
	}
	c.JSON(http.StatusOK, devices)
}
