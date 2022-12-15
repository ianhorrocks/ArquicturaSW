package app

import (

	"time"

	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin" //framework que nos ayuda hacer todas las conecciones marshall unmarshial
	log "github.com/sirupsen/logrus"
)

var (
	router *gin.Engine
)

func init() { //palabra reservada que ejecuta cuando se compila
	log.Info("inittt")
	router = gin.Default()
	router.Use(cors.New(cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
			AllowHeaders: []string{"Access-Control-Allow-Origin", "*"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
	}))
}

func StartRoute() { // llama un metodo mapUrls
	mapUrls()

	log.Info("Starting Server")
	router.Run(":8090")
}
