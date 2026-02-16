package repository

import (
	"context"
	"errors"
	"gopher-finance-engine/internal/domain/entity"
	"gopher-finance-engine/internal/domain/infra/repository"
	"gopher-finance-engine/pkg/postgres"

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
