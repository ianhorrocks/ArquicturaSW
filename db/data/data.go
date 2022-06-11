package data

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	userModel "ArquicturaSW/model"
)

func InsertData(db *gorm.DB) {
	log.Info("Inserting data... ")

	err := db.First(&userModel.User{}).Error

	if err != nil {
		db.Create(&userModel.User{Name: "Ian", LastName: "Horrocks", UserName: "ianho10", Email: "@ucc.edu.ar", Password: "12345"})
		db.Create(&userModel.User{Name: "Nano", LastName: "Nallar", UserName: "lokito", Email: "qseyo@ucc.edu.ar", Password: "12345"})
		db.Create(&userModel.User{Name: "Elias", LastName: "Stuyck", UserName: "loki3", Email: "qseyo@ucc.edu.ar", Password: "12345"})
	}

	log.Info("Data inserted")

}
/*
package database

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	addressModel "github.com/facugarces29/ArquiSoftware/model/address"
	productModel "github.com/facugarces29/ArquiSoftware/model/product"
	userModel "github.com/facugarces29/ArquiSoftware/model/user"
)

func InsertData(db *gorm.DB) {
	// Insert data
	log.Info("Inserting data...")

	//Inserting users
	err := db.First(&userModel.User{}).Error

	if err != nil {
		db.Create(&userModel.User{Name: "lautaro", LastName: "Saenz", UserName: "lautarose", Email: "abcdefg@gmail.com", Pwd: "hola123"})
		db.Create(&userModel.User{Name: "Joaco", LastName: "Reyero", UserName: "jaocoreyero", Email: "12345@gmail.com", Pwd: "hola123"})
		db.Create(&userModel.User{Name: "Facundo", LastName: "Garces", UserName: "Facuelcapo", Email: "asasas@gmail.com", Pwd: "hola123"})
		db.Create(&userModel.User{Name: "Hernan", LastName: "Lachampionliga", UserName: "hernanchampion", Email: "hernan@gmail.com", Pwd: "hola123"})
		db.Create(&userModel.User{Name: "Saul", LastName: "Hudson", UserName: "slash", Email: "slashGNR@gmail.com", Pwd: "hola123"})
	}

	//Inserting Categories

	err = db.First(&productModel.Category{}).Error

	if err != nil {
		db.Create(&productModel.Category{Name: "Electronics"})
		db.Create(&productModel.Category{Name: "Electrodomestics"})
		db.Create(&productModel.Category{Name: "Fitness"})
		db.Create(&productModel.Category{Name: "Toys"})
		db.Create(&productModel.Category{Name: "Electronics"})
		db.Create(&productModel.Category{Name: "Supermarket"})
		db.Create(&productModel.Category{Name: "Babies"})
		db.Create(&productModel.Category{Name: "Fashion"})
	}

	//Inserting products

	err = db.First(&productModel.Product{}).Error

	if err != nil {
		db.Create(&productModel.Product{CategoryID: 1, Name: "Asus Laptop", Description: "Personal laptop, intel i5, 8gb ram, fullhd screen, windows 10", Price: 100000, Stock: 10, Image: ""})
		db.Create(&productModel.Product{CategoryID: 1, Name: "Logitech Mouse", Description: "Gaming Mouse 16000dpi", Price: 8000, Stock: 15, Image: "https://resource.logitechg.com/w_1000,c_limit,q_auto,f_auto,dpr_auto/d_transparent.gif/content/dam/products/gaming/gaming-mice/g305-lightspeed-wireless-gaming-mouse/g304-g305-lightspeed-wireless-gaming-mouse21.png?v=1"})
		db.Create(&productModel.Product{CategoryID: 1, Name: "Hyperx Keyboard", Description: "Keyboard rgb, cherry, us distribution", Price: 12000, Stock: 4, Image: ""})
		db.Create(&productModel.Product{CategoryID: 1, Name: "Samsung Monitor", Description: "Resolution: 1920x1080, 144hz, 32 inches", Price: 125000, Stock: 2, Image: ""})
		db.Create(&productModel.Product{CategoryID: 3, Name: "Randers Multigym", Description: "75kg, more than 30 exercises, aluminium", Price: 85000, Stock: 3, Image: ""})
		db.Create(&productModel.Product{CategoryID: 3, Name: "Smart Tech Fixed Bycicle", Description: "Spinning exercises, high endurance", Price: 35000, Stock: 1, Image: ""})
		db.Create(&productModel.Product{CategoryID: 3, Name: "Pull-Up Bar", Description: "Pull-up bar", Price: 9500, Stock: 6, Image: ""})
		db.Create(&productModel.Product{CategoryID: 3, Name: "Dumbell and Kettlebell kit", Description: "2x10kg dumbell,2x5kg kettlebell", Price: 7000, Stock: 5, Image: ""})
		db.Create(&productModel.Product{CategoryID: 2, Name: "Samsung Refrigerator", Description: "freezer, 555L, no frost", Price: 250000, Stock: 2, Image: ""})
		db.Create(&productModel.Product{CategoryID: 2, Name: "Atma toaster", Description: "220v", Price: 4500, Stock: 8, Image: ""})
		db.Create(&productModel.Product{CategoryID: 2, Name: "samsung Washing Machine", Description: "up to 9kg", Price: 65000, Stock: 4, Image: ""})
		db.Create(&productModel.Product{CategoryID: 2, Name: "BGH Electric Oven", Description: "Electric Oven", Price: 20300, Stock: 9, Image: ""})
		db.Create(&productModel.Product{CategoryID: 4, Name: "Skateboard", Description: "Skateboard", Price: 10000, Stock: 3, Image: ""})
		db.Create(&productModel.Product{CategoryID: 4, Name: "Hot-Wheels Camaro", Description: "1980 Camaro", Price: 2000, Stock: 10, Image: ""})
		db.Create(&productModel.Product{CategoryID: 4, Name: "Nerf n-strike", Description: "nerf n-strike", Price: 5500, Stock: 2, Image: ""})
		db.Create(&productModel.Product{CategoryID: 4, Name: "Lego Spidey", Description: "100 pcs", Price: 7990, Stock: 1, Image: ""})
	}

	//Inserting addresses

	err = db.First(&addressModel.Address{}).Error

	if err != nil {
		db.Create(&addressModel.Address{UserID: 1, State: "Cordoba", City: "Cordoba", Zip: 5000, Addressline: "Eugenio Garzon 467"})
		db.Create(&addressModel.Address{UserID: 2, State: "Cordoba", City: "Cordoba", Zip: 5000, Addressline: "Ituzaingo 200"})
		db.Create(&addressModel.Address{UserID: 3, State: "Cordoba", City: "Cordoba", Zip: 5000, Addressline: "General Paz 344"})
		db.Create(&addressModel.Address{UserID: 4, State: "Cordoba", City: "Cordoba", Zip: 5000, Addressline: "Velez Sarsfield 1500"})
		db.Create(&addressModel.Address{UserID: 5, State: "Cordoba", City: "Cordoba", Zip: 5000, Addressline: "Belgrano 110"})
	}

	//manage errors...

	log.Info("Data inserted")
}
*/