package orders

import (
	"context"
	"gopher-finance-engine/internal/domain/entity"
	"log"

	"github.com/google/uuid"
)

func (u *OrdersUsecase) CreateOrders(ctx context.Context, body *entity.Order) error {
	body.ID = uuid.NewString()
	body.Status = "PENDING"

	err := u.repo.CreateOrders(ctx, body)

	if err != nil {
		u.logger.Error("error create orders")
		return err
	}

	log.Println(*body)

	return nil
}
