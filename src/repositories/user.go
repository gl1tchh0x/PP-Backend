package repositories

import (
	"errors"

	"backend.com/go-backend/src/config"
	"backend.com/go-backend/src/models"
	"github.com/alexedwards/argon2id"
)

type CreateUserInput struct {
	Avatar   string `json:"avatar"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	Password string `json:"password" binding:"required"`
	IsStaff  bool   `json:"is_staff"`
}

func CreateUserRepository(data CreateUserInput) error {
	// check if user already exists
	var existingUser models.User
	if err := config.DB.Where("email = ? OR username = ?", data.Email, data.Username).
		First(&existingUser).Error; err == nil {
		return errors.New("user with the given email or username already exists")
	}

	// Create a new user
	hashedPassword, err := argon2id.CreateHash(data.Password, argon2id.DefaultParams)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user := models.User{
		Avatar:   data.Avatar,
		Email:    data.Email,
		Username: data.Username,
		FullName: data.FullName,
		Password: hashedPassword,
		IsStaff:  data.IsStaff,
	}

	tx := config.DB.Begin()
	if err := tx.Create(&user).Error; err != nil {
		return errors.New("failed to create user")
	}
	tx.Commit()
	return nil
}