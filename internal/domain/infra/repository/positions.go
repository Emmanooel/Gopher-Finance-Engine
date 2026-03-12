package repository

import (
	"context"
	"gopher-finance-engine/internal/domain/entity"
)

type PositionsRepositoryI interface {
	GetPositionByUserId(ctx context.Context, userId string) ([]*entity.Positions, error)
	SaveNewPosition(ctx context.Context, pos *entity.Positions) error
}
