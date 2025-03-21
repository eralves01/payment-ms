package services

import (
	"github.com/eralves01/payment-ms/database"
	"github.com/eralves01/payment-ms/models"
	"github.com/eralves01/payment-ms/repository"
)

type UserServices struct{}

func NewUserServices() *UserServices {
	return &UserServices{}
}

func (s *UserServices) CreateUser(user *models.User) error {
	repo := getRepository()
	err := repo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserServices) GetUsers() ([]models.User, error) {
	repository := getRepository()
	return repository.GetUsers()
}

func (s *UserServices) GetUserByEmail(email string) (*models.User, error) {
	repository := getRepository()
	return repository.GetUserByEmail(email)
}

func (s *UserServices) GetUserByID(id string) (*models.User, error) {
	repository := getRepository()
	return repository.GetUserByID(id)
}

func (s *UserServices) GetUsersByFilters(query map[string][]string) ([]models.User, error) {
	repository := getRepository()
	return repository.GetUsersByFilters(query)
}

func getRepository() *repository.UserRepository {
	db := database.GetInstance()
	return repository.NewUserRepository(db)
}
