package services

import (
	userCliente "ArquicturaSW/clients/user"
	"ArquicturaSW/dto"
	"ArquicturaSW/model"
	e "ArquicturaSW/utils/errors"

	log "github.com/sirupsen/logrus"

	"crypto/md5"
	"encoding/hex"

	"github.com/dgrijalva/jwt-go"
)

type userService struct{}

// ACA EL SERVICE SE COMUNICA CON EL MODEL(BUSINESS CLASSES) Y CONTROLLER(DTOs)
type userServiceInterface interface { //nombre del metodo, que .. todo lo que va a devolver mi servicio es en DTO (ej: user)
	GetUserById(id int) (dto.UserDto, e.ApiError)
	LoginUser(loginDto dto.LoginDto) (dto.TokenDto, e.ApiError)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserById(id int) (dto.UserDto, e.ApiError) {

	var user model.User = userCliente.GetUserById(id) // crea una var user pidiendole los datos a la bd
	var userDto dto.UserDto                           //Lo que tengo que devolver

	if user.Id == 0 {
		return userDto, e.NewBadRequestApiError("user not found")
	}
	userDto.Id = user.Id
	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.UserName = user.UserName
	userDto.Password = user.Password
	userDto.Address = user.Address

	return userDto, nil
}

var jwtKey = []byte("secret_key")

func (s *userService) LoginUser(loginDto dto.LoginDto) (dto.TokenDto, e.ApiError) {

	log.Debug(loginDto)
	var user model.User = userCliente.GetByUsername(loginDto.Username)

	var tokenDto dto.TokenDto

	if user.Id == 0 {
		return tokenDto, e.NewBadRequestApiError("user not found")
	}

	var pswMd5 = md5.Sum([]byte(loginDto.Password)) //hasheando la passw (to convert string into byte form for processing by md5 hashing algorithm)
	pswMd5String := hex.EncodeToString(pswMd5[:])   // return the hexadecimal encoding

	if pswMd5String == user.Password {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id_user": user.Id,
		})
		tokenString, _ := token.SignedString(jwtKey)
		tokenDto.Token = tokenString //asigna la llave a los tokendto del usuario en particular
		tokenDto.IdUser = user.Id

		return tokenDto, nil
	} else {
		return tokenDto, e.NewBadRequestApiError("contraseña incorrecta")
	}

}

/*
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
*/
