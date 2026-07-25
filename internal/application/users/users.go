package users

import (
	"context"
	"gopher-finance-engine/internal/application/utils"
	"gopher-finance-engine/internal/domain/entity"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (u *UsersUsecase) CreateUser(ctx context.Context, body *entity.Users) error {
	body.Id = uuid.NewString()
	body.CreatedAt = time.Now()

	pass, err := utils.HashPassword(body.Password)
	body.Password = pass

	err = u.repo.CreateUser(ctx, *body)

	if err != nil {
		u.logger.Error("CreateUser", zap.Error(err))
		return err
	}

	return nil
}

func (u *UsersUsecase) GetAllUsers(ctx context.Context, page int) (entity.UsersResponse, error) {
	users, err := u.repo.GetAllUsers(ctx)

	if err != nil {
		u.logger.Error("GetUsers", zap.Error(err))
		return entity.UsersResponse{}, err
	}

	return entity.UsersResponse{Data: users}, nil
}
