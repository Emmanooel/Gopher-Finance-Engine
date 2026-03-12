package entity

import (
	"time"

	"github.com/google/uuid"
)

type Positions struct {
	Id           string    `json:"id"`
	UserId       string    `json:"userId"`
	Symbol       string    `json:"symbol"`
	TotalAmount  int64     `json:"total_amount"`
	AveragePrice int64     `json:"average_price"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ResponsePositions struct {
	Positions []*Positions `json:"positions"`
}

func (p *Positions) BuildPositionByOrder(o Order) {
	p.Id = uuid.NewString()
	p.UserId = o.UserId
	p.Symbol = o.Symbol
	p.TotalAmount = o.Amount
	p.AveragePrice = (o.Price * o.Amount) / o.Amount
	p.UpdatedAt = time.Now()
}
