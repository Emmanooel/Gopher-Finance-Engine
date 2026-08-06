package orders

import (
	"context"
	"gopher-finance-engine/internal/application/strategy"
	"gopher-finance-engine/internal/domain/entity"

	"go.uber.org/zap"
)

const (
	BUY  = "BUY"
	SELL = "SELL"
)

func (u *OrdersUsecase) ProcessPendingOrders(ctx context.Context) error {
	u.logger.Info("Iniciate process orders")

	pendingOrders, err := u.repo.GetAllOrdersInPending(ctx, 10)
	if err != nil {
		u.logger.Info("error get orders in pending")
		return err
	}

	sideStrategy := strategy.NewSideStrategy(u.logger, u.pS)

	for _, order := range pendingOrders {
		positions, err := u.getPositions(ctx, order.UserId, order.Symbol, order)

		if err != nil {
			u.logger.Info("error get position by userId and symbol")
			return err
		}

		strategy, err := sideStrategy.Get(order.Side)

		if err != nil {
			u.logger.Info("error on get strategy")
			return err
		}

		err = strategy.Execute(ctx, positions, order)

		if err != nil {
			u.logger.Info("error on execute strategy")
			return err
		}

		err = u.markOrderProcessed(ctx, order.ID)

		if err != nil {
			u.logger.Info("error on update order status")
			return err
		}

		u.logger.Info("order processed successfully", zap.Any("order_id: ", order.ID))
	}

	u.logger.Info("Process all orders completed")

	return nil
}

func (u *OrdersUsecase) getPositions(ctx context.Context, userId string, symbol string, order *entity.Order) (*entity.Positions, error) {
	positions, err := u.pS.SearchPositionByUserIdAndSymbol(ctx, userId, symbol)

	if err != nil {
		u.logger.Info("error get position by userId and symbol")
		return nil, err
	}

	if positions == nil {
		u.logger.Info("symbol don't save in user wallet, iniciate save position")
		positions, err = u.SaveNewPosition(ctx, order)

		if err != nil {
			u.logger.Info("error on update position by order")
			return nil, err
		}
	}
	return positions, nil
}

func (u *OrdersUsecase) SaveNewPosition(ctx context.Context, order *entity.Order) (*entity.Positions, error) {
	positions, err := u.pS.SaveNewPositionByOrder(ctx, order)

	if err != nil {
		u.logger.Info("error save new position")
		return nil, err
	}

	return positions, nil
}

func (u *OrdersUsecase) markOrderProcessed(ctx context.Context, order_id string) error {
	u.logger.Info("updating status order")
	status := PROCESSED_STATUS

	err := u.repo.UpdateStatusOrders(ctx, order_id, status)

	if err != nil {
		return err
	}

	u.logger.Info("update status order sucessfully", zap.Any("order_status: ", status))
	return nil
}
