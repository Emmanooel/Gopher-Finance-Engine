package handlers

import (
	"gopher-finance-engine/internal/application/usecases/orders"
	"gopher-finance-engine/internal/application/usecases/positions"
	"gopher-finance-engine/internal/application/usecases/users"
)

type Handlers struct {
	UserUsecase     users.UserUsecasesI
	PositionUsecase positions.PositionUsecasesI
	OrderUsecase    orders.OrdersUsecaseI
}

func NewHandlers(
	userUsecase users.UserUsecasesI,
	positionUsecase positions.PositionUsecasesI,
	orderUsecase orders.OrdersUsecaseI,
) *Handlers {
	return &Handlers{
		UserUsecase:     userUsecase,
		PositionUsecase: positionUsecase,
		OrderUsecase:    orderUsecase,
	}
}
