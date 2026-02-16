package repository

import (
	"context"
	"gopher-finance-engine/internal/domain/entity"
)

type UsersRepositoryI interface {
	CreateUser(ctx context.Context, users entity.Users) error
	FindByEmail(ctx context.Context, email string) (*entity.Users, error)
}
