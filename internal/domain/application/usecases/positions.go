package usecases

import (
	"context"
	"gopher-finance-engine/internal/domain/entity"
)

type PositionUsecasesI interface {
	GetPositionByUserId(ctx context.Context, userId string) (*entity.ResponsePositions, error)
	SavePositionByNewOrder(ctx context.Context, order *entity.Order) chan bool
}
