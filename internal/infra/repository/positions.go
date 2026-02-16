package repository

import (
	"context"
	"errors"
	"gopher-finance-engine/internal/domain/entity"
	"gopher-finance-engine/internal/domain/infra/repository"
	"gopher-finance-engine/pkg/postgres"

	"go.uber.org/zap"
)

type PositionRepository struct {
	logger *zap.Logger
}

func NewPositionRepository(
	logger *zap.Logger,
) repository.PositionsRepositoryI {
	return &PositionRepository{
		logger: logger,
	}
}

func (p *PositionRepository) GetAllPositions(ctx context.Context, id string) ([]*entity.Positions, error) {
	db := postgres.Db

	if db == nil {
		p.logger.Error("conn database is null")
		return nil, errors.New("conn database is null")
	}
	const query = `SELECT * FROM positions
		WHERE id = $1
		ORDER BY updated_at DESC
	`

	row, err := db.Query(ctx, query, id)

	var positions []*entity.Positions
	var o *entity.Positions

	for row.Next() {
		err = row.Scan(
			&o.Id,
			&o.UserId,
			&o.Symbol,
			&o.TotalAmount,
			&o.AveragePrice,
			&o.UpdatedAt,
		)

		if err != nil {
			p.logger.Error("error unmarshal results database")
			return nil, errors.New("error unmarshal results")
		}

		positions = append(positions, o)
	}

	return positions, nil
}
