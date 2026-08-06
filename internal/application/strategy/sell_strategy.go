package strategy

import (
	"context"
	"errors"
	"gopher-finance-engine/internal/application/positions/service"
	"gopher-finance-engine/internal/domain/entity"

	"go.uber.org/zap"
)

type SellStrategy struct {
	logger          *zap.Logger
	positionService service.PositionServiceI
}

func NewSellStrategy(
	logger *zap.Logger,
	positionService service.PositionServiceI,
) *SellStrategy {
	return &SellStrategy{
		logger:          logger,
		positionService: positionService,
	}
}

func (s *SellStrategy) Execute(ctx context.Context, position *entity.Positions, order *entity.Order) error {
	if position.TotalAmount < order.Amount {
		s.logger.Error("[SellStrategy] sales quantity less than total quantity")
		return errors.New("sales quantity less than total quantity")
	}
	p := position.UpdatePositionBySellOrder(*order)

	err := s.positionService.UpdatePositionByOrder(ctx, order, p)

	if err != nil {
		s.logger.Error("[SellStrategy] error on update position by order")
		return err
	}

	s.logger.Info("[SellStrategy] update order status for process sucessully")

	return nil
}
