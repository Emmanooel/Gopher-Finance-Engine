package positions

import (
	"context"
	"gopher-finance-engine/internal/domain/entity"
	"gopher-finance-engine/internal/domain/infra/repository"

	"go.uber.org/zap"
)

type PositionUsecasesI interface {
	SavePositionByNewOrder(ctx context.Context, order *entity.Order) (*entity.Positions, error)
	GetPositionByUserId(ctx context.Context, id string) (*entity.ResponsePositions, error)
	SearchPositionByUserIdAndSymbol(ctx context.Context, userId, symbol string) (*entity.Positions, error)
	ListAllPositionByUserId(ctx context.Context, id string) (*entity.ResponsePositions, error)
	UpdatePositionByUserIdAndSymbol(ctx context.Context, userId string, symbol string, position *entity.Positions) error
	DeletePositionByUserId(ctx context.Context, userId string) error
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
