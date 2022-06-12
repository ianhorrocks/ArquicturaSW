package clients

import (
	"ArquicturaSW/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetAddressById(id int) model.Address {
	var address model.Address

	Db.Where("id = ?", id).First(&address)
	log.Debug("Address: ", address)

	return address
}

func GetAddresses() model.Addresses {
	var addresses model.Addresses

	Db.Find(&addresses)

	log.Debug("Addresses: ", addresses)

	return addresses
}