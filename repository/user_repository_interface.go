package repository

import "github.com/eralves01/payment-ms/models"

type UserRepositoryInterface interface {
	CreateUser(user *models.User) error
	GetUsers() ([]models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	GetUsersByFilters(query map[string][]string) ([]models.User, error)
}
