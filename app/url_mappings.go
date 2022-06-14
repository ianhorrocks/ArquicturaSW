package app

import (
	//productController "ArquicturaSW/controllers/product"
	userController "ArquicturaSW/controllers/user"
	productController "ArquicturaSW/controllers/product" //igual para category.controller
	loginController "ArquicturaSW/controllers/login"
	searchController "ArquicturaSW/controllers/search"


	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Users
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)
	// Products
	router.GET("/product/:id", productController.GetProductById)
	router.GET("/product", productController.GetProducts)
	// Categories
	router.GET("/category/:id", productController.GetCategoryById)
	router.GET("/category", productController.GetCategories) //productos por categoria HACERLO
	
	// METODOS
	// Login
	router.POST("/login", loginController.Login)
	// Search
	router.GET("/search/:param", searchController.GetProductsBySearchParam)


	log.Info("Finishing mappings configurations") // igual que starting server
}
