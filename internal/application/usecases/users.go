package usecases

import (
	"context"
	"gopher-finance-engine/internal/application/service"
	"gopher-finance-engine/internal/domain/application/usecases"
	"gopher-finance-engine/internal/domain/entity"
	"gopher-finance-engine/internal/domain/infra/repository"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type UsersUsecase struct {
	logger *zap.Logger
	repo   repository.UsersRepositoryI
	auth   service.TokenService
}

func NewUsersUsecase(
	logger *zap.Logger,
	repo repository.UsersRepositoryI,
	auth service.TokenService,
) usecases.UserUsecasesI {
	return &UsersUsecase{
		logger: logger,
		repo:   repo,
		auth:   auth,
	}
}

func (u *UsersUsecase) CreateUser(ctx context.Context, body *entity.Users) error {
	body.Id = uuid.NewString()
	body.CreatedAt = time.Now()

	pass, err := HashPassword(body.Password)
	body.Password = pass

	err = u.repo.CreateUser(ctx, *body)

	if err != nil {
		u.logger.Error("CreateUser", zap.Error(err))
		return err
	}

	return nil
}

func (u *UsersUsecase) Login(ctx context.Context, body entity.UserLogin) (string, error) {
	user, err := u.repo.FindByEmail(ctx, body.Email)

	if err != nil {
		u.logger.Error("Login", zap.Error(err))
		return "", err
	}

	match := CheckPasswordHash(body.Password, user.Password)

	if !match {
		return "", err
	}

	token, err := u.auth.GenToken(user.Id)

	if err != nil {
		u.logger.Error("Login", zap.Error(err))
		return "", err
	}

	return token, nil
}

func (u *UsersUsecase) GetAllUsers(ctx context.Context, page int) (entity.UsersResponse, error) {
	_, err := u.repo.FindByEmail(ctx, "")

	if err != nil {
		u.logger.Error("GetUsers", zap.Error(err))
		return entity.UsersResponse{}, err
	}

	return entity.UsersResponse{}, nil
}
