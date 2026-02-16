package entity

import "time"

type TaxAlert struct {
	Id           string    `json:"id"`
	UserId       string    `json:"user_id"`
	Symbol       string    `json:"symbol"`
	Profit       int64     `json:"profit"`
	CalculatedAt time.Time `json:"calculated_at"`
}
