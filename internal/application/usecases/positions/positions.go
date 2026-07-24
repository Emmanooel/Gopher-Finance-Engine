package positions

import (
	"context"
	"gopher-finance-engine/internal/domain/entity"
)

func (p *PositionUsecase) GetPositionByUserId(ctx context.Context, id string) (*entity.ResponsePositions, error) {
	var output *entity.ResponsePositions
	position, err := p.PositionRepo.GetPositionByUserId(ctx, id)

	if err != nil {
		p.Logger.Error("error get positions, err:" + err.Error())
		return nil, err
	}

	output.Positions = position

	return output, nil
}

func (p *PositionUsecase) SavePositionByNewOrder(ctx context.Context, order *entity.Order) chan bool {
	position := &entity.Positions{}

	position.BuildPositionByOrder(*order)

	err := p.PositionRepo.SaveNewPosition(ctx, position)

	if err != nil {
		p.Logger.Error("error save new position")
		return nil
	}

	return nil
}
