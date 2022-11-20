package dto

type UserDto struct {
	Id       int   `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

type UsersDto []UserDto
