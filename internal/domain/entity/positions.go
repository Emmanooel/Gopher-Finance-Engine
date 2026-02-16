package entity

import "time"

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
