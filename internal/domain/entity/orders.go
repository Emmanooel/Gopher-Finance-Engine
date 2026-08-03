package entity

import (
	"time"
)

type Order struct {
	ID        string    `json:"id"`
	UserId    string    `json:"user_id"`
	Symbol    string    `json:"Symbol"`
	Amount    int64     `json:"amount" `
	Price     int64     `json:"price"`
	Side      string    `json:"side"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at`
}

type OrderResponse struct {
	Data []*Order `json:"data"`
	Meta Meta     `json:"meta"`
}
