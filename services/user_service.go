package services

import (
	"github.com/eralves01/payment-ms/models"
	"github.com/eralves01/payment-ms/repository"
)

type UserServices struct {
	repository repository.UserRepositoryInterface
}

func NewUserServices(repository repository.UserRepositoryInterface) *UserServices {
	return &UserServices{repository: repository}
}

func (s *UserServices) CreateUser(user *models.User) error {
	err := s.repository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserServices) GetUsers() ([]models.User, error) {

	return s.repository.GetUsers()
}

func (s *UserServices) GetUserByEmail(email string) (*models.User, error) {

	return s.repository.GetUserByEmail(email)
}

func (s *UserServices) GetUserByID(id string) (*models.User, error) {

	return s.repository.GetUserByID(id)
}

func (s *UserServices) GetUsersByFilters(query map[string][]string) ([]models.User, error) {

	return s.repository.GetUsersByFilters(query)
}
