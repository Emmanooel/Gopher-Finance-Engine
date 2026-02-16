package repository

import (
	"context"
	"gopher-finance-engine/internal/domain/entity"
)

type OrdersRepositoryI interface {
	CreateOrders(ctx context.Context, orders *entity.Order) error
}
