package app

import (
	"github.com/gin-gonic/gin" //framework que nos ayuda hacer todas las conecciones marshall unmarshial
	log "github.com/sirupsen/logrus"
)

var (
	router *gin.Engine
)

func init() { //palabra reservada que ejecuta cuando se compila
	router = gin.Default()
}

func StartRoute() { // llama un metodo mapUrls
	mapUrls()

	log.Info("Starting server") // para comprobar que llegamos bien a conectar el servidor
	router.Run(":8090")

}
