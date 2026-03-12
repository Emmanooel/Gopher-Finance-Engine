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
		return errors.New("error create order, error:" + err.Error())
	}

	return nil
}

func (o *OrdersRepository) GetOrdersInPendingByUserId(ctx context.Context, userId string) ([]*entity.Order, error) {
	db := postgres.Db

	if db == nil {
		o.logger.Error("conn database is null")
		return nil, errors.New("conn database is null")
	}

	const query = `SELECT * FROM orders
		WHERE user_id = $1 AND status = 'PENDING'
		LIMIT 100
	`
	rows, err := db.Query(ctx, query)

	if err == pgx.ErrNoRows {
		o.logger.Error("none rows as returned")
		return nil, err
	}

	var output []*entity.Order
	var b *entity.Order

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
		)

		if err != nil {
			o.logger.Error("err on unmarshal database returns for struct:" + err.Error())
		}
	}

	return output, nil
}
