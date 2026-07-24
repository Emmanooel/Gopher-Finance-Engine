package positions

import (
	"context"
	"gopher-finance-engine/internal/domain/entity"
	"gopher-finance-engine/internal/domain/infra/repository"

	"go.uber.org/zap"
)

type PositionUsecasesI interface {
	GetPositionByUserId(ctx context.Context, userId string) (*entity.ResponsePositions, error)
	SavePositionByNewOrder(ctx context.Context, order *entity.Order) chan bool
}
type PositionUsecase struct {
	Logger       *zap.Logger
	PositionRepo repository.PositionsRepositoryI
	OrdersRepo   repository.OrdersRepositoryI
}

func NewPositionUsecase(
	logger *zap.Logger,
	repo repository.PositionsRepositoryI,
	orders repository.OrdersRepositoryI,
) PositionUsecasesI {
	return &PositionUsecase{
		Logger:       logger,
		PositionRepo: repo,
		OrdersRepo:   orders,
	}
}
