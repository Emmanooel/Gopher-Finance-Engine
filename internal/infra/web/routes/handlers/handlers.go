package handlers

import "gopher-finance-engine/internal/domain/application/usecases"

type Handlers struct {
	UserUsecase     usecases.UserUsecasesI
	PositionUsecase usecases.PositionUsecasesI
	OrderUsecase    usecases.OrdersUsecaseI
}

func NewHandlers(
	userUsecase usecases.UserUsecasesI,
	positionUsecase usecases.PositionUsecasesI,
	orderUsecase usecases.OrdersUsecaseI,
) *Handlers {
	return &Handlers{
		UserUsecase:     userUsecase,
		PositionUsecase: positionUsecase,
		OrderUsecase:    orderUsecase,
	}
}
