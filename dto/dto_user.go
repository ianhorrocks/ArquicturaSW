package dto

type UserDto struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UsersDto []UserDto
