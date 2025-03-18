package services

import (
	"payment-ms/database"
	"payment-ms/models"
	"payment-ms/repository"
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

func (s *UserServices) GetUserByEmail(email string) (*models.User, error) {
	repository := getRepository()
	return repository.GetUserByEmail(email)
}

func (s *UserServices) GetUserByID(id string) (*models.User, error) {
	repository := getRepository()
	return repository.GetUserByID(id)
}

func getRepository() *repository.UserRepository {
	db := database.GetInstance()
	return repository.NewUserRepository(db)
}
