package usecases

import (
	"context"
	"gopher-finance-engine/internal/domain/application/usecases"
	"gopher-finance-engine/internal/domain/entity"
	"gopher-finance-engine/internal/domain/infra/repository"
	"gopher-finance-engine/worker"
	"log"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type OrdersUsecase struct {
	logger          *zap.Logger
	repo            repository.OrdersRepositoryI
	positionUsecase usecases.PositionUsecasesI
	worker          worker.WorkerSaveNewPositionI
}

func NewOrdersUsecase(
	logger *zap.Logger,
	repo repository.OrdersRepositoryI,
	position usecases.PositionUsecasesI,
	w worker.WorkerSaveNewPositionI,
) usecases.OrdersUsecaseI {
	return &OrdersUsecase{
		logger:          logger,
		repo:            repo,
		positionUsecase: position,
		worker:          w,
	}
}

func (u *OrdersUsecase) CreateOrders(ctx context.Context, body *entity.Order) error {
	body.ID = uuid.NewString()
	body.Status = "PENDING"

	err := u.repo.CreateOrders(ctx, body)

	if err != nil {
		u.logger.Error("error create orders")
		return err
	}

	log.Println(*body)
	u.worker.WorkerSavePositionByNewOrder(ctx, body)

	return nil
}
