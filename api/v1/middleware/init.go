package middleware

import (
	"gin-api-test/api/v1/models"
	"gin-api-test/config"
	"gin-api-test/db/connect"
	"gin-api-test/helpers/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	if config.MyEnv.Env == "PROD" {
		log.Log.Printf("Starting %s on Production Environment", config.MyEnv.AppName)
	} else {
		log.Log.Printf("Starting %s on Production Environment", config.MyEnv.AppName)
	}

	log.SetupLogger()

	err := connect.InitConnectionDB()
	if err != nil {
		log.Log.Error(err.Error())
	}
}

func HandleNoMethod(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, models.CreateResponse(http.StatusMethodNotAllowed, "failed", "Undefined Method", nil))
}

func HandleNoRoute(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, models.CreateResponse(http.StatusMethodNotAllowed, "failed", "route Undefined", nil))
}

func HandlePanic(c *gin.Context, err interface{}) {
	log.Log.Error(err.(error).Error())
	c.JSON(http.StatusInternalServerError, models.CreateResponse(http.StatusInternalServerError, "failed", "Internal server Error", nil))
}
