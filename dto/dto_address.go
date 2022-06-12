package dto

type AddressDto struct {
	Id          uint   `json:"id"`
	UserId      uint   `json:"user_id"`
	State       string `json:"state"`
	City        string `json:"city"`
	Zip         int    `json:"zip"`
	AddressLine string `json:"address_line"`
}

type AddressesDto []AddressDto