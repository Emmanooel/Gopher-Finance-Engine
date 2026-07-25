package orders

import (
	"context"
	"gopher-finance-engine/internal/domain/entity"
	"gopher-finance-engine/internal/domain/infra/repository"

	"go.uber.org/zap"
)

type OrdersUsecaseI interface {
	CreateOrders(ctx context.Context, body *entity.Order) error
	ProcessOrders(ctx context.Context, userId string) error
}

type OrdersUsecase struct {
	logger *zap.Logger
	repo   repository.OrdersRepositoryI
}

func NewOrdersUsecase(
	logger *zap.Logger,
	repo repository.OrdersRepositoryI,
) OrdersUsecaseI {
	return &OrdersUsecase{
		logger: logger,
		repo:   repo,
	}
}
