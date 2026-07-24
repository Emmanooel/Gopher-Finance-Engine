package users

import (
	"context"
	"gopher-finance-engine/internal/domain/entity"
	"gopher-finance-engine/internal/domain/infra/repository"
	"gopher-finance-engine/internal/infra/auth"

	"go.uber.org/zap"
)

type UserUsecasesI interface {
	CreateUser(ctx context.Context, body *entity.Users) error
	Login(ctx context.Context, body entity.UserLogin) (string, error)
	GetAllUsers(ctx context.Context, page int) (entity.UsersResponse, error)
}

type UsersUsecase struct {
	logger *zap.Logger
	repo   repository.UsersRepositoryI
	auth   auth.TokenService
}

func NewUsersUsecase(
	logger *zap.Logger,
	repo repository.UsersRepositoryI,
	auth auth.TokenService,
) UserUsecasesI {
	return &UsersUsecase{
		logger: logger,
		repo:   repo,
		auth:   auth,
	}
}
