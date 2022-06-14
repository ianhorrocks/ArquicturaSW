package services

import (
	userDto "ArquicturaSW/dto"
	model "ArquicturaSW/model"
	loginDto "ArquicturaSW/dto"
	userCliente "ArquicturaSW/clients/user"

	log "github.com/sirupsen/logrus"
)

type loginService struct{}

type loginServiceInterface interface {
	Login(loginDto.LoginDto) userDto.UserDto
}

var (
	LoginService loginServiceInterface
)

func init() {
	LoginService = &loginService{}
}

func (s *loginService) Login(login loginDto.LoginDto) userDto.UserDto {

	id := login.Id
	pass := login.Password
	log.Debug("pass", pass)
	var user model.User = userCliente.GetUserById(id)
	var userDto userDto.UserDto

	if user.UserID == 0 {
		return userDto
	}

	if user.Password != pass {
		return userDto
	}

	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.UserName = user.UserName
	userDto.Id = user.UserID
	userDto.Email = user.Email
	userDto.Password = user.Password

	return userDto

}
