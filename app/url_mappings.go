package app

import (
	//productController "ArquicturaSW/controllers/product"
	userController "ArquicturaSW/controllers/user"
	productController "ArquicturaSW/controllers/product"


	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Products Mapping
	//router.GET("/product/:product_id", productController.GetProductByEan)
	//router.GET("/product", productController.GetProducts)

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)
	// Products Mapping
	router.GET("/product", productController.GetProducts)

	log.Info("Finishing mappings configurations") // igual que starting server
}
