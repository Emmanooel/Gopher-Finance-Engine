package users

import (
	"context"
	"gopher-finance-engine/internal/domain/entity"
	"gopher-finance-engine/internal/utils"

	"go.uber.org/zap"
)

func (u *UsersUsecase) Login(ctx context.Context, body entity.UserLogin) (string, error) {
	user, err := u.repo.FindByEmail(ctx, body.Email)

	if err != nil {
		u.logger.Error("Login", zap.Error(err))
		return "", err
	}

	match := utils.CheckPasswordHash(body.Password, *user.Password)

	if !match {
		return "", err
	}

	u.logger.Info("login", zap.Any("user_id", user.Id), zap.Any("auth", u.auth))
	token, err := u.auth.GenToken(user.Id)

	if err != nil {
		u.logger.Error("Login", zap.Error(err))
		return "", err
	}

	return token, nil
}
