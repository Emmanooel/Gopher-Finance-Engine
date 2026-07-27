package repository

import (
	"context"
	"gopher-finance-engine/internal/domain/entity"
)

type PositionsRepositoryI interface {
	GetPositionByUserId(ctx context.Context, userId string) ([]*entity.Positions, error)
	UpdatePosition(ctx context.Context, position *entity.Positions) error
	SaveNewPosition(ctx context.Context, position *entity.Positions) error
}
