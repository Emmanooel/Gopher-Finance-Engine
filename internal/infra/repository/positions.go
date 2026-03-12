package repository

import (
	"context"
	"errors"
	"gopher-finance-engine/internal/domain/entity"
	"gopher-finance-engine/internal/domain/infra/repository"
	"gopher-finance-engine/pkg/postgres"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

var (
	ErrUserNotPosition = errors.New("user not have position activate in wallet")
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

func (p *PositionRepository) SaveNewPosition(ctx context.Context, pos *entity.Positions) error {
	db := postgres.Db

	const query = `INSERT INTO positions (id, user_id, symbol, total_amount, average_price, updated_at)
					VALUES ($1, $2, $3, $4, $5, $6)

					ON CONFLICT (user_id, symbol)
					DO UPDATE
					SET total_amount = positions.total_amount + EXCLUDED.total_amount,
						average_price =
						(
							(positions.average_price * positions.total_amount) +
							(EXCLUDED.average_price * EXCLUDED.total_amount)
						)
						/
						NULLIF(positions.total_amount + EXCLUDED.total_amount, 0),
						updated_at = EXCLUDED.updated_at
	`

	_, err := db.Exec(
		ctx,
		query,
		pos.Id,
		pos.UserId,
		pos.Symbol,
		pos.TotalAmount,
		pos.AveragePrice,
		pos.UpdatedAt,
	)

	if err != nil {
		p.logger.Error("error save new position, error:" + err.Error())
		return err
	}

	return nil

}

func (p *PositionRepository) GetPositionByUserId(ctx context.Context, userId string) ([]*entity.Positions, error) {
	db := postgres.Db

	if db == nil {
		p.logger.Error("database is null")
		return nil, errors.New("database is null")
	}

	const query = `SELECT * FROM positions
	WHERE user_id = $1`

	rows, err := db.Query(ctx, query, userId)

	if err == pgx.ErrNoRows {
		p.logger.Info("user not have position")
		return nil, ErrUserNotPosition
	}

	var output []*entity.Positions
	var pp *entity.Positions
	for rows.Next() {
		err = rows.Scan(
			&pp.Id,
			&pp.UserId,
			&pp.Symbol,
			&pp.TotalAmount,
			&pp.AveragePrice,
			pp.UpdatedAt,
		)

		if err != nil {
			p.logger.Error("errors at unsmarshal return database:" + err.Error())
			return nil, err
		}
		output = append(output, pp)
	}

	return output, nil
}
