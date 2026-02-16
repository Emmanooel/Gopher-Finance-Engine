package usecases

import (
	"context"
	"gopher-finance-engine/internal/domain/application/usecases"
	"gopher-finance-engine/internal/domain/entity"
	"gopher-finance-engine/internal/domain/infra/repository"

	"go.uber.org/zap"
)

type PositionUsecase struct {
	Logger *zap.Logger
	Repo   repository.PositionsRepositoryI
}

func NewPositionUsecase(logger *zap.Logger, repo repository.PositionsRepositoryI) usecases.PositionUsecasesI {
	return &PositionUsecase{
		Logger: logger,
		Repo:   repo,
	}
}

func (p *PositionUsecase) GetPositionByUserId(ctx context.Context, id string) (*entity.ResponsePositions, error) {
	positions, err := p.Repo.GetAllPositions(ctx, id)

	if err != nil {
		p.Logger.Error("GetAllPositions", zap.Error(err))
		return nil, err
	}

	resp := &entity.ResponsePositions{
		Positions: positions,
	}

	return resp, nil
}
