package orders_repository

import (
	"gopher-finance-engine/internal/domain/entity"
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID        uuid.UUID
	UserId    uuid.UUID
	Symbol    string
	Amount    int64
	Price     int64
	Side      string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (o *Order) BuildEntity() *entity.Order {
	return &entity.Order{
		ID:        o.ID.String(),
		UserId:    o.UserId.String(),
		Symbol:    o.Symbol,
		Amount:    o.Amount,
		Price:     o.Price,
		Side:      o.Side,
		Status:    o.Status,
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
	}
}
