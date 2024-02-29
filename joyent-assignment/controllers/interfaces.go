package controllers

import (
	"joynet-assignment/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InterfacesInterface interface {
	HandleGetInterfaces(*gin.Context)
}
type interfacesController struct {
	e *config.Env
}

func NewInterfacesController(env *config.Env) InterfacesInterface {
	return &interfacesController{
		e: env,
	}
}

func (ic *interfacesController) HandleGetInterfaces(c *gin.Context) {
	interfaces, err := ic.e.Db.GetInterfaces()
	if err != nil {
		ic.e.Log.Error(err)
		response := struct{}{}
		c.JSON(http.StatusInternalServerError, response)
	}
	c.JSON(http.StatusOK, interfaces)
}
