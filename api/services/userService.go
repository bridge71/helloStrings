package services

import (
	"github.com/bridge71/helloStrings/api/repositories"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (s *UserService) CheckUser(c *gin.Context) {
	s.UserRepository.CreaterUser(c)
}
