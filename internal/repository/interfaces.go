package repository

import (
	"github.com/MBFG9000/golang-jwt-auth/internal/entity"
	"github.com/google/uuid"
)

type (
	UserRepoInterface interface {
		CreateUser(user *entity.User) error
		GetUserByUsername(username string) (*entity.User, error)
		GetUserByID(id uuid.UUID) (*entity.User, error)
		UpdateUserRole(id uuid.UUID, role string) error
	}
)
