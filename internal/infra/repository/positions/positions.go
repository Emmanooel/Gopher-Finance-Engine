package positions_repository

import (
	"context"
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

func (p *PositionRepository) SaveNewPosition(ctx context.Context, position *entity.Positions) error {
	tx, err := postgres.Db.Begin(ctx)
	if err != nil {
		p.logger.Error("error connect database", zap.Error(err))
		return err
	}

	defer tx.Rollback(ctx)

	const query = `INSERT INTO positions (id, user_id, symbol, total_amount, average_price, total_cost)
					VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err = tx.Exec(
		ctx,
		query,
		position.Id,
		position.UserId,
		position.Symbol,
		position.TotalAmount,
		position.AveragePrice,
		position.TotalCost,
	)

	if err != nil {
		p.logger.Error("error save new position, error:", zap.Error(err))
		return err
	}

	return tx.Commit(ctx)

}

func (p *PositionRepository) UpdatePosition(ctx context.Context, position *entity.Positions) error {
	tx, err := postgres.Db.Begin(ctx)
	if err != nil {
		p.logger.Error("error connect database", zap.Error(err))
		return err
	}

	defer tx.Rollback(ctx)

	const query = `
	UPDATE positions
	SET
		total_amount = $1,
		average_price = $2,
		total_cost = $3
	WHERE id = $4
`

	_, err = tx.Exec(
		ctx,
		query,
		position.TotalAmount,
		position.AveragePrice,
		position.TotalCost,
		position.Id,
	)

	if err != nil {
		p.logger.Error("error save new position, error:", zap.Error(err))
		return err
	}

	return tx.Commit(ctx)
}

func (p *PositionRepository) GetPositionByUserId(ctx context.Context, userId string) ([]*entity.Positions, error) {
	tx, err := postgres.Db.Begin(ctx)

	if err != nil {
		p.logger.Error("error connect database", zap.Error(err))
		return nil, err
	}

	defer tx.Rollback(ctx)

	const query = `SELECT * FROM positions
	WHERE user_id = $1`

	rows, err := tx.Query(ctx, query, userId)

	if err != nil {
		p.logger.Info("error consult database: ", zap.Error(err))
		return nil, err
	}

	var output []*entity.Positions
	var position Positions
	for rows.Next() {
		err = rows.Scan(
			&position.Id,
			&position.UserId,
			&position.Symbol,
			&position.TotalAmount,
			&position.AveragePrice,
			&position.TotalCost,
			&position.UpdatedAt,
		)

		if err != nil {
			p.logger.Error("errors at unsmarshal return database:", zap.Error(err))
			return nil, err
		}
		output = append(output, position.BuildEntity())
	}

	if len(output) <= 0 {
		p.logger.Info("user not have position activate in wallet")
		return nil, tx.Commit(ctx)
	}

	p.logger.Info("search position by userID sucessfull", zap.Any("output:", output))
	return output, tx.Commit(ctx)
}
