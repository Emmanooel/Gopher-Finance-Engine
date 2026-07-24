package orders

import (
	"context"
	"gopher-finance-engine/internal/application/usecases/positions"
	"gopher-finance-engine/internal/domain/entity"
	"gopher-finance-engine/internal/domain/infra/repository"

	"go.uber.org/zap"
)

type OrdersUsecaseI interface {
	CreateOrders(ctx context.Context, body *entity.Order) error
	ProcessOrders(ctx context.Context, body *entity.Order) error
}

type OrdersUsecase struct {
	logger          *zap.Logger
	repo            repository.OrdersRepositoryI
	positionUsecase positions.PositionUsecasesI
}

func NewOrdersUsecase(
	logger *zap.Logger,
	repo repository.OrdersRepositoryI,
	position positions.PositionUsecasesI,
) OrdersUsecaseI {
	return &OrdersUsecase{
		logger:          logger,
		repo:            repo,
		positionUsecase: position,
	}
}
