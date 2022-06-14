package services

import (
	userCliente "ArquicturaSW/clients/user"
	 "ArquicturaSW/dto"
	 "ArquicturaSW/model"
	e "ArquicturaSW/utils/errors"
)

type userService struct{}

// ACA EL SERVICE SE COMUNICA CON EL MODEL(BUSINESS CLASSES) Y CONTROLLER(DTOs)
type userServiceInterface interface { //nombre del metodo, que .. todo lo que va a devolver mi servicio es en DTO (ej: user)
	GetUserById(id int) (dto.UserDto, e.ApiError) 
	GetUsers() (dto.UsersDto, e.ApiError)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserById(id int) (dto.UserDto, e.ApiError) {

	var user model.User = userCliente.GetUserById(id)// crea una var user pidiendole los datos a la bd
	var userDto dto.UserDto //Lo que tengo que devolver

	if user.UserID == 0 {
		return userDto, e.NewBadRequestApiError("user not found")
	}
	userDto.Id = user.UserID
	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.UserName = user.UserName
	userDto.Email = user.Email
	userDto.Password = user.Password
	
	return userDto, nil
}

func (s *userService) GetUsers() (dto.UsersDto, e.ApiError) {

	var users model.Users = userCliente.GetUsers()
	var usersDto dto.UsersDto

	for _, user := range users {
		var userDto dto.UserDto
		userDto.Id = user.UserID
		userDto.Name = user.Name
		userDto.LastName = user.LastName
		userDto.UserName = user.UserName
		userDto.Email = user.Email
		userDto.Password = user.Password
		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *userService) GetByUsername(username string) (dto.UserDto, error) {

	user, err := userCliente.GetByUsername(username)
	var userDto dto.UserDto

	if err != nil {
		return userDto, err
	}

	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.UserName = user.UserName
	userDto.Id = user.UserID
	userDto.Email = user.Email
	userDto.Password = user.Password
	return userDto, nil
}