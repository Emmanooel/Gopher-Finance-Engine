package handlers

import (
	"gopher-finance-engine/internal/application/orders"
	"gopher-finance-engine/internal/application/positions"
	"gopher-finance-engine/internal/application/users"

	"go.uber.org/zap"
)

type Handlers struct {
	logger          *zap.Logger
	UserUsecase     users.UserUsecasesI
	PositionUsecase positions.PositionUsecasesI
	OrderUsecase    orders.OrdersUsecaseI
}

func NewHandlers(
	logger *zap.Logger,
	userUsecase users.UserUsecasesI,
	positionUsecase positions.PositionUsecasesI,
	orderUsecase orders.OrdersUsecaseI,
) *Handlers {
	return &Handlers{
		logger:          logger,
		UserUsecase:     userUsecase,
		PositionUsecase: positionUsecase,
		OrderUsecase:    orderUsecase,
	}
}
