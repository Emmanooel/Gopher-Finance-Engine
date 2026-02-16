package usecases

import (
	"context"
	"gopher-finance-engine/internal/domain/application/usecases"
	"gopher-finance-engine/internal/domain/entity"
	"gopher-finance-engine/internal/domain/infra/repository"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type OrdersUsecase struct {
	logger *zap.Logger
	repo   repository.OrdersRepositoryI
}

func NewOrdersUsecase(
	logger *zap.Logger,
	repo repository.OrdersRepositoryI,
) usecases.OrdersUsecaseI {
	return &OrdersUsecase{
		logger: logger,
		repo:   repo,
	}
}

func (u *OrdersUsecase) CreateOrders(ctx context.Context, body *entity.Order) error {
	body.ID = uuid.NewString()
	body.Status = "PENDING"

	err := u.repo.CreateOrders(ctx, body)

	if err != nil {
		u.logger.Error("error create orders")
		return err
	}

	return nil
}
