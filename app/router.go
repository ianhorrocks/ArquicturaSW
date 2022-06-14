package app

import (
	"github.com/gin-gonic/gin" //framework que nos ayuda hacer todas las conecciones marshall unmarshial
)

var (
	router *gin.Engine
)

func init() { //palabra reservada que ejecuta cuando se compila
	router = gin.Default()
}

func StartRoute() { // llama un metodo mapUrls
	mapUrls()
	router.Run(":8090")
}
