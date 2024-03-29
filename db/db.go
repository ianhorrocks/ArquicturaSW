package db

import (
	"ArquicturaSW/model"
	userClient "ArquicturaSW/clients/user"
	productClient "ArquicturaSW/clients/product"
	categoryClient "ArquicturaSW/clients/category"
	orderClient "ArquicturaSW/clients/order"
	orderDetailClient "ArquicturaSW/clients/order_detail"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := "prueba"
	DBUser := "root"
	DBPass := "1234"
	DBHost := "127.0.0.1"
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	db.LogMode(true)

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all CLients that we build
	userClient.Db = db
	productClient.Db = db
	categoryClient.Db = db
	orderClient.Db = db
	orderDetailClient.Db = db
	
}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&model.User{}, &model.Product{}, &model.Category{}, &model.Order{}, &model.OrderDetail{})
	log.Info("Finishing Migration Database Tables")
}
