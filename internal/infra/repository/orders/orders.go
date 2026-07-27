package orders_repository

import (
	"context"
	"errors"
	"gopher-finance-engine/internal/domain/entity"
	"gopher-finance-engine/internal/domain/infra/repository"
	"gopher-finance-engine/pkg/postgres"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type OrdersRepository struct {
	logger *zap.Logger
}

func NewOrdersRepository(
	logger *zap.Logger,
) repository.OrdersRepositoryI {
	return &OrdersRepository{
		logger: logger,
	}
}

func (o *OrdersRepository) CreateOrders(ctx context.Context, order *entity.Order) error {
	db := postgres.Db

	if db == nil {
		o.logger.Error("conn database is null")
		return errors.New("conn database is null")
	}

	const query = ` INSERT INTO orders (id, user_id, symbol, amount, price, side, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := db.Exec(
		ctx,
		query,
		order.ID,
		order.UserId,
		order.Symbol,
		order.Amount,
		order.Price,
		order.Side,
		order.Status,
		order.CreatedAt,
	)

	if err != nil {
		o.logger.Error("error create order", zap.Error(err))
		return errors.New("error create order")
	}

	return nil
}

func (o *OrdersRepository) GetOrdersInPendingByUserId(ctx context.Context, userId string) ([]*entity.Order, error) {
	tx, err := postgres.Db.Begin(ctx)
	if err != nil {
		o.logger.Error("error connect database", zap.Error(err))
		return nil, err
	}

	defer tx.Rollback(ctx)

	const query = `SELECT * FROM orders
		WHERE user_id = $1 AND status = 'PENDING'
	`
	rows, err := tx.Query(ctx, query, userId)

	if err != nil {
		if err == pgx.ErrNoRows {
			o.logger.Error("none rows as returned")
			return nil, err
		}

		o.logger.Error("error on query database:", zap.Error(err))
		return nil, err
	}

	var output []*entity.Order
	var b Order

	for rows.Next() {
		err := rows.Scan(
			&b.ID,
			&b.UserId,
			&b.Symbol,
			&b.Amount,
			&b.Price,
			&b.Side,
			&b.Status,
			&b.CreatedAt,
			&b.UpdatedAt,
		)

		if err != nil {
			o.logger.Error("err on unmarshal database returns for struct: ", zap.Error(err))
		}
		output = append(output, b.BuildEntity())
	}

	o.logger.Info("[orders_repository] search orders pending by userid sucessfully")
	return output, tx.Commit(ctx)
}

func (o *OrdersRepository) UpdateStatusOrders(ctx context.Context, order_id, status string) error {
	tx, err := postgres.Db.Begin(ctx)
	if err != nil {
		o.logger.Error("error connect database", zap.Error(err))
		return err
	}

	defer tx.Rollback(ctx)

	const query = `
	UPDATE orders
	SET
		status = $1
	WHERE id = $2
`

	_, err = tx.Exec(
		ctx,
		query,
		status,
		order_id,
	)

	if err != nil {
		o.logger.Error("error update order status, error:", zap.Error(err))
		return err
	}

	return tx.Commit(ctx)
}
