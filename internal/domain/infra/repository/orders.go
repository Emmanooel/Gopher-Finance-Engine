package repository

import (
	"context"
	"gopher-finance-engine/internal/domain/entity"
)

type OrdersRepositoryI interface {
	CreateOrders(ctx context.Context, orders *entity.Order) error
	GetAllOrdersInPending(ctx context.Context, limit int) ([]*entity.Order, error)
	UpdateStatusOrders(ctx context.Context, order_id, status string) error
	GetAllOrdersByUserId(ctx context.Context, userId string) ([]*entity.Order, error)
}
