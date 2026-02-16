package repository

import (
	"context"
	"gopher-finance-engine/internal/domain/entity"
)

type PositionsRepositoryI interface {
	GetAllPositions(ctx context.Context, id string) ([]*entity.Positions, error)
}
