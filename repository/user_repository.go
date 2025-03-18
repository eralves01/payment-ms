package repository

import (
	"database/sql"

	"github.com/eralves01/payment-ms/models"
)

type UserRepositoryInterface interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	_, err := r.db.Exec("INSERT INTO users (name, email, password_hash, user_type) VALUES ($1, $2, $3, $4)", user.Name, user.Email, user.PasswordHash, user.UserType)
	return err
}

func (r *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.UserType, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.UserType, &user.CreatedAt, &user.UpdatedAt)
	return &user, err
}

func (r *UserRepository) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.UserType, &user.CreatedAt, &user.UpdatedAt)
	return &user, err
}
