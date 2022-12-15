package app

import (
	categoryController "ArquicturaSW/controllers/category"
	productController "ArquicturaSW/controllers/product" //igual para category.controller
	userController "ArquicturaSW/controllers/user"
	orderDetailController "ArquicturaSW/controllers/order_detail"
	orderController "ArquicturaSW/controllers/order"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Users
	router.GET("/user/:id", userController.GetUserById)
	router.POST("/login", userController.UserLogin)
	
	// Products
	router.GET("/product/:id", productController.GetProductById)
	router.GET("/product", productController.GetProducts)
	router.GET("/productRandom/:cantidad", productController.GetProductsByNumber)

	// Search product by text
	router.GET("/productText/:texto", productController.GetProductsByText)

	// Products by category
	router.GET("/productCategory/:idCategory", productController.GetProductsByCategory)

	// Categories
	router.GET("/category", categoryController.GetCategories)

	// OrdersDetails by order ID
	router.GET("/ordersDetails/:idOrder", orderDetailController.GetOrderDetailByIdOrder)

	// Insert order
	router.POST("/order", orderController.OrderInsert)

	// Orders by user ID
	router.GET("/ordersByUser/:idUser", orderController.GetOrdersByIdUser)


	log.Info("Finishing mappings configurations")
}
