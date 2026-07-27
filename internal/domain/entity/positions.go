package entity

import (
	"time"

	"github.com/google/uuid"
)

type Positions struct {
	Id           string    `json:"id"`
	UserId       string    `json:"userId"`
	Symbol       string    `json:"symbol"`
	TotalCost    int64     `json:"total_cost"`
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
	p.TotalCost = o.Price * o.Amount
	p.TotalAmount = o.Amount
	p.AveragePrice = p.TotalCost / p.TotalAmount
}

func (p *Positions) UpdatePositionByPurchaceOrder(o Order) *Positions {
	totalCost := p.TotalCost + (o.Amount * o.Price)
	totalAmount := p.TotalAmount + o.Amount
	averagePrice := totalCost / totalAmount
	return &Positions{
		Id:           p.Id,
		UserId:       p.UserId,
		Symbol:       p.Symbol,
		TotalCost:    totalCost,
		TotalAmount:  totalAmount,
		AveragePrice: averagePrice,
	}
}

func (p *Positions) UpdatePositionBySellOrder(o Order) *Positions {
	totalCost := p.TotalCost - (p.AveragePrice * o.Amount)
	totalAmount := p.TotalAmount - o.Amount
	return &Positions{
		Id:           p.Id,
		UserId:       p.UserId,
		Symbol:       p.Symbol,
		TotalCost:    totalCost,
		TotalAmount:  totalAmount,
		AveragePrice: p.AveragePrice,
	}
}
