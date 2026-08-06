package strategy

import (
	"context"
	"gopher-finance-engine/internal/application/positions/service"
	"gopher-finance-engine/internal/domain/entity"

	"go.uber.org/zap"
)

type BuyStrategy struct {
	logger          *zap.Logger
	positionService service.PositionServiceI
}

func NewBuyStrategy(
	logger *zap.Logger,
	positionService service.PositionServiceI,
) *BuyStrategy {
	return &BuyStrategy{
		logger:          logger,
		positionService: positionService,
	}
}

func (b *BuyStrategy) Execute(ctx context.Context, position *entity.Positions, order *entity.Order) error {
	p := position.UpdatePositionByPurchaceOrder(*order)
	err := b.positionService.UpdatePositionByOrder(ctx, order, p)

	if err != nil {
		b.logger.Error("[BuyStrategy] error on update position by order")
		return err
	}

	b.logger.Info("[BuyStrategy] update order status for process sucessully")

	return nil
}
