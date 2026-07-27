package service

import (
	"context"
	"gopher-finance-engine/internal/application/positions"
	"gopher-finance-engine/internal/domain/entity"

	"go.uber.org/zap"
)

type PositionServiceI interface {
	SaveNewPositionByOrder(ctx context.Context, order *entity.Order) (*entity.Positions, error)
	SearchPositionByUserIdAndSymbol(ctx context.Context, userId, symbol string) (*entity.Positions, error)
	UpdatePositionByOrder(ctx context.Context, order *entity.Order, position *entity.Positions) error
}

type PositionService struct {
	logger *zap.Logger
	PU     positions.PositionUsecasesI
}

func NewPositionService(
	logger *zap.Logger,
	PUsecase positions.PositionUsecasesI,
) PositionServiceI {
	return &PositionService{
		logger: logger,
		PU:     PUsecase,
	}
}

func (p *PositionService) SaveNewPositionByOrder(ctx context.Context, order *entity.Order) (*entity.Positions, error) {
	p.logger.Info("save position by new order")
	positions, err := p.PU.SavePositionByNewOrder(ctx, order)

	if err != nil {
		p.logger.Info("error save positions by orders")
		return nil, err
	}

	p.logger.Info("save position by order successfully", zap.Any("order", order))
	return positions, nil
}

func (p *PositionService) SearchPositionByUserIdAndSymbol(ctx context.Context, userId, symbol string) (*entity.Positions, error) {
	p.logger.Info("Search position by userId and symbol", zap.String("userId", userId), zap.String("symbol", symbol))
	position, err := p.PU.SearchPositionByUserIdAndSymbol(ctx, userId, symbol)

	if err != nil {
		p.logger.Info("error processing positions by orders")
		return nil, err
	}

	p.logger.Info("search position by userId and symbol successfully")
	return position, nil
}

func (p *PositionService) UpdatePositionByOrder(ctx context.Context, order *entity.Order, position *entity.Positions) error {
	p.logger.Info("iniciate update position")
	err := p.PU.UpdatePositionByUserIdAndSymbol(ctx, order.UserId, order.Symbol, position)

	if err != nil {
		p.logger.Info("error on update position")
		return err
	}

	p.logger.Info("updating position by order sucessfully")
	return nil
}
