package usecases

import (
	"context"
	"gopher-finance-engine/internal/domain/application/usecases"
	"gopher-finance-engine/internal/domain/entity"
	"gopher-finance-engine/internal/domain/infra/repository"

	"go.uber.org/zap"
)

type PositionUsecase struct {
	Logger       *zap.Logger
	PositionRepo repository.PositionsRepositoryI
	OrdersRepo   repository.OrdersRepositoryI
}

func NewPositionUsecase(
	logger *zap.Logger,
	repo repository.PositionsRepositoryI,
	orders repository.OrdersRepositoryI,
) usecases.PositionUsecasesI {
	return &PositionUsecase{
		Logger:       logger,
		PositionRepo: repo,
		OrdersRepo:   orders,
	}
}

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
