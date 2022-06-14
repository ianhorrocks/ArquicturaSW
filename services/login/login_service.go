package services

import (
	userCliente "ArquicturaSW/clients/user"
	loginDto "ArquicturaSW/dto"
	userDto "ArquicturaSW/dto"
	"errors"

	log "github.com/sirupsen/logrus"
)

type loginService struct{}

type loginServiceInterface interface {
	Login(loginDto.LoginDto) (userDto.UserDto, error)
}

var (
	LoginService loginServiceInterface
)

func init() {
	LoginService = &loginService{}
}

func (s *loginService) Login(login loginDto.LoginDto) (userDto.UserDto, error) {

	username := login.Username
	pass := login.Password
	log.Debug("pass", pass)
	user, err:= userCliente.GetByUsername(username)
	var userDto userDto.UserDto

	if err != nil {
		log.Println(err)
		return userDto, err
	}

	if user.Password != pass {
		err = errors.New("Credenciales incorrectas")
		log.Println(err)
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
