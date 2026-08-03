package orders

import (
	"context"
	"gopher-finance-engine/internal/domain/entity"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	PENDING_STATUS   = "PENDING"
	PROCESSED_STATUS = "PROCESSED"
	FAILED_STATUS    = "FAILED"
)

func (u *OrdersUsecase) CreateOrders(ctx context.Context, body *entity.Order) error {
	body.ID = uuid.NewString()
	body.Status = PENDING_STATUS

	err := u.repo.CreateOrders(ctx, body)

	if err != nil {
		u.logger.Info("error create orders")
		return err
	}

	u.logger.Info("order created: ", zap.Any("body", *body))

	return nil
}

func (u *OrdersUsecase) UpdateStatusOrders(ctx context.Context, order_id string) error {
	u.logger.Info("updating status order")
	status := PROCESSED_STATUS

	err := u.repo.UpdateStatusOrders(ctx, order_id, status)

	if err != nil {
		return err
	}

	u.logger.Info("update status order sucessfully", zap.Any("order_status: ", status))
	return nil
}

func (u *OrdersUsecase) GetAllOrdersByUserId(ctx context.Context, userId string) (*entity.OrderResponse, error) {
	u.logger.Info("getting all orders by userId")
	orders, err := u.repo.GetAllOrdersByUserId(ctx, userId)

	if err != nil {
		u.logger.Info("error get all orders by userId")
		return nil, err
	}

	return &entity.OrderResponse{
		Data: orders,
	}, nil
}
