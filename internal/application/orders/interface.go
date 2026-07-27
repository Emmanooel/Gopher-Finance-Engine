package orders

import (
	"context"
	"gopher-finance-engine/internal/application/positions/service"
	"gopher-finance-engine/internal/domain/entity"
	"gopher-finance-engine/internal/domain/infra/repository"

	"go.uber.org/zap"
)

type OrdersUsecaseI interface {
	CreateOrders(ctx context.Context, body *entity.Order) error
	ProcessOrders(ctx context.Context, userId string) error
	UpdateStatusOrders(ctx context.Context, order_id string) error
}

type OrdersUsecase struct {
	logger *zap.Logger
	repo   repository.OrdersRepositoryI
	pS     service.PositionServiceI
}

func NewOrdersUsecase(
	logger *zap.Logger,
	repo repository.OrdersRepositoryI,
	pS service.PositionServiceI,
) OrdersUsecaseI {
	return &OrdersUsecase{
		logger: logger,
		repo:   repo,
		pS:     pS,
	}
}
