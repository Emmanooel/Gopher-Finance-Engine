package users

import (
	"context"
	"gopher-finance-engine/internal/application/utils"
	"gopher-finance-engine/internal/domain/entity"
	"time"

	"github.com/google/uuid"
)

func (u *UsersUsecase) CreateUser(ctx context.Context, body *entity.Users) error {
	body.Id = uuid.NewString()
	body.CreatedAt = time.Now()

	pass, err := utils.HashPassword(*body.Password)
	body.Password = &pass

	err = u.repo.CreateUser(ctx, *body)

	if err != nil {
		u.logger.Info("CreateUser")
		return err
	}

	return nil
}

func (u *UsersUsecase) GetAllUsers(ctx context.Context, page int) (entity.UsersResponse, error) {
	users, err := u.repo.GetAllUsers(ctx)

	if err != nil {
		u.logger.Info("GetUsers")
		return entity.UsersResponse{}, err
	}

	return entity.UsersResponse{Data: users}, nil
}

func (u *UsersUsecase) GetUserById(ctx context.Context, id string) (*entity.User, error) {
	user, err := u.repo.GetUserById(ctx, id)

	if err != nil {
		u.logger.Info("[GetUserById] error get user by id")
		return nil, err
	}

	return user, nil
}
