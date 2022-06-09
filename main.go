package main

import (
	"ArquicturaSW/app"
	"ArquicturaSW/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
// prueba




// siempre una app en go tiene que tener esto
// Hay distintos paquetes (app)
// Controllers
// Paquete de clases que viajan entre el controlador y el servicio
//service no lo vamos a usar