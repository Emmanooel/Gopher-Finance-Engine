package positions_repository

import (
	"gopher-finance-engine/internal/domain/entity"
	"time"

	"github.com/google/uuid"
)

type Positions struct {
	Id           uuid.UUID
	UserId       uuid.UUID
	Symbol       string
	TotalAmount  int64
	AveragePrice int64
	TotalCost    int64
	UpdatedAt    time.Time
}

func (p *Positions) BuildEntity() *entity.Positions {
	return &entity.Positions{
		Id:           p.Id.String(),
		UserId:       p.UserId.String(),
		Symbol:       p.Symbol,
		TotalAmount:  p.TotalAmount,
		AveragePrice: p.AveragePrice,
		TotalCost:    p.TotalCost,
		UpdatedAt:    p.UpdatedAt,
	}
}
