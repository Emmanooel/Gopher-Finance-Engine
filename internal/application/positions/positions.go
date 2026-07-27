package positions

import (
	"context"
	"gopher-finance-engine/internal/domain/entity"

	"go.uber.org/zap"
)

func (p *PositionUsecase) SavePositionByNewOrder(ctx context.Context, order *entity.Order) (*entity.Positions, error) {
	p.Logger.Info("SavePositionByNewOrder", zap.String("userId", order.UserId), zap.String("symbol", order.Symbol))
	position := &entity.Positions{}

	position.BuildPositionByOrder(*order)

	go p.PositionRepo.SaveNewPosition(ctx, position)

	p.Logger.Info("save new position sucessfully", zap.Any("new position", position))
	return position, nil
}

func (p *PositionUsecase) GetPositionByUserId(ctx context.Context, id string) (*entity.ResponsePositions, error) {
	var output *entity.ResponsePositions
	position, err := p.PositionRepo.GetPositionByUserId(ctx, id)

	if err != nil {
		p.Logger.Info("error get positions")
		return nil, err
	}

	output.Positions = position

	return output, nil
}

func (p *PositionUsecase) SearchPositionByUserIdAndSymbol(ctx context.Context, userId, symbol string) (*entity.Positions, error) {
	p.Logger.Info("search position by userId: " + userId + " and " + "symbol: " + symbol)

	positions, err := p.PositionRepo.GetPositionByUserId(ctx, userId)

	if err != nil {
		p.Logger.Info("[positionUsecase] error get position by userId")
		return nil, err
	}

	for _, p := range positions {
		if p.Symbol == symbol {
			return p, nil
		}
	}

	return nil, nil
}

func (p *PositionUsecase) ListAllPositionByUserId(ctx context.Context, id string) (*entity.ResponsePositions, error) {
	var output *entity.ResponsePositions
	position, err := p.PositionRepo.GetPositionByUserId(ctx, id)

	if err != nil {
		p.Logger.Info("error list positions")
		return nil, err
	}

	output.Positions = position

	return output, nil
}

func (p *PositionUsecase) UpdatePositionByUserIdAndSymbol(ctx context.Context, userId string, symbol string, position *entity.Positions) error {
	return nil
}

func (p *PositionUsecase) DeletePositionByUserId(ctx context.Context, userId string) error {
	return nil
}
