package usecases

import (
	"context"
	"gopher-finance-engine/internal/domain/entity"
)

type OrdersUsecaseI interface {
	CreateOrders(ctx context.Context, body *entity.Order) error
}
