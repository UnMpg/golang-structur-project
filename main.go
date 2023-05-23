package main

import (
	"gin-api-test/api/v1/middleware"
	"gin-api-test/api/v1/models"
	"gin-api-test/api/v1/routes"
	"gin-api-test/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	var isDev bool

	if config.MyEnv.Env == "PROD" {
		isDev = false
	} else {
		isDev = true
	}

	r := routes.CreateRoute(isDev)

	r.GET("/", func(c *gin.Context) {
		message := "Welcome to golang Gorm and postgress"
		res := models.Response{StatusCode: http.StatusOK, Status: "success", Message: message}
		c.JSON(http.StatusOK, res)
	})

	r.HandleMethodNotAllowed = true
	r.NoMethod(middleware.HandleNoMethod)
	r.NoRoute(middleware.HandleNoRoute)
	routes.UserRouteV1(r)

	timeout := time.Duration(config.MyEnv.Timeout) * time.Second
	newHandler := http.TimeoutHandler(r, timeout, "Timeout!")

	server := &http.Server{
		Addr:         config.MyEnv.ServerPort,
		Handler:      newHandler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 25 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Could not listen on %v : % v \n", config.MyEnv.ServerPort, err)
	}

}
