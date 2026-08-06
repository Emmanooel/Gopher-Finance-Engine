package strategy

import (
	"context"
	"fmt"
	"gopher-finance-engine/internal/application/positions/service"
	"gopher-finance-engine/internal/domain/entity"

	"go.uber.org/zap"
)

type SideStrategyI interface {
	Execute(ctx context.Context, position *entity.Positions, order *entity.Order) error
}

type SideStrategy struct {
	strategies map[string]SideStrategyI
}

func NewSideStrategy(
	logger *zap.Logger,
	positionService service.PositionServiceI,
) *SideStrategy {

	return &SideStrategy{
		strategies: map[string]SideStrategyI{
			entity.BUY:  NewBuyStrategy(logger, positionService),
			entity.SELL: NewSellStrategy(logger, positionService),
		},
	}
}

func (r *SideStrategy) Get(side string) (SideStrategyI, error) {
	strategy, ok := r.strategies[side]

	if !ok {
		return nil, fmt.Errorf("unsupported side: %s", side)
	}

	return strategy, nil
}
