package data

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	model "ArquicturaSW/model"
)

func InsertData(db *gorm.DB) {
	log.Info("Inserting data... ")

	err := db.First(&model.User{}).Error

	if err != nil {
		db.Create(&model.User{Name: "Ian", LastName: "Horrocks", UserName: "ianho10", Email: "@ucc.edu.ar", Password: "12345"})
		db.Create(&model.User{Name: "Nano", LastName: "Nallar", UserName: "nanonallar", Email: "qseyo@ucc.edu.ar", Password: "12345"})
		db.Create(&model.User{Name: "Elian", LastName: "Stuyck", UserName: "elian123", Email: "qseyo@ucc.edu.ar", Password: "12345"})
		db.Create(&model.User{Name: "Agus", LastName: "Ferrero", UserName: "agusferrero", Email: "qseyo@ucc.edu.ar", Password: "12345"})
		db.Create(&model.User{Name: "Tomi", LastName: "Calvo", UserName: "tomicalvo", Email: "qseyo@ucc.edu.ar", Password: "12345"})
	}

	err = db.First(&model.Product{}).Error

	if err != nil {
		db.Create(&model.Product{Name: "Asus Laptop", Description: "Personal laptop, intel i5, 8gb ram, fullhd screen, windows 10", Price: 100000, Stock: 10})
		db.Create(&model.Product{Name: "Logitech Mouse", Description: "Gaming Mouse 16000dpi", Price: 8000, Stock: 15})
		db.Create(&model.Product{Name: "Hyperx Keyboard", Description: "Keyboard rgb, cherry, us distribution", Price: 12000, Stock: 4})
		db.Create(&model.Product{Name: "Samsung Monitor", Description: "Resolution: 1920x1080, 144hz, 32 inches", Price: 125000, Stock: 2})
		db.Create(&model.Product{Name: "Randers Multigym", Description: "75kg, more than 30 exercises, aluminium", Price: 85000, Stock: 3})
		db.Create(&model.Product{Name: "Smart Tech Fixed Bycicle", Description: "Spinning exercises, high endurance", Price: 35000, Stock: 1})
		db.Create(&model.Product{Name: "Pull-Up Bar", Description: "Pull-up bar", Price: 9500, Stock: 6})
		db.Create(&model.Product{Name: "Dumbell and Kettlebell kit", Description: "2x10kg dumbell,2x5kg kettlebell", Price: 7000, Stock: 5})
		db.Create(&model.Product{Name: "Samsung Refrigerator", Description: "freezer, 555L, no frost", Price: 250000, Stock: 2})
		db.Create(&model.Product{Name: "Atma toaster", Description: "220v", Price: 4500, Stock: 8})
		db.Create(&model.Product{Name: "samsung Washing Machine", Description: "up to 9kg", Price: 65000, Stock: 4})
		db.Create(&model.Product{Name: "BGH Electric Oven", Description: "Electric Oven", Price: 20300, Stock: 9})
		db.Create(&model.Product{Name: "Skateboard", Description: "Skateboard", Price: 10000, Stock: 3})
		db.Create(&model.Product{Name: "Hot-Wheels Camaro", Description: "1980 Camaro", Price: 2000, Stock: 10})
		db.Create(&model.Product{Name: "Nerf n-strike", Description: "nerf n-strike", Price: 5500, Stock: 2})
		db.Create(&model.Product{Name: "Lego Spidey", Description: "100 pcs", Price: 7990, Stock: 1})
	}

	log.Info("Data inserted")

}