package usecases

import (
	"context"
	"gopher-finance-engine/internal/domain/entity"
)

type UserUsecasesI interface {
	CreateUser(ctx context.Context, body *entity.Users) error
	Login(ctx context.Context, body entity.UserLogin) (string, error)
	GetAllUsers(ctx context.Context, page int) (entity.UsersResponse, error)
}
