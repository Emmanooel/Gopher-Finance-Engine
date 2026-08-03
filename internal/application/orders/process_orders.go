package orders

import (
	"context"
	"errors"
	"gopher-finance-engine/internal/domain/entity"
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

	for _, order := range pendingOrders {
		positions, err := u.pS.SearchPositionByUserIdAndSymbol(ctx, order.UserId, order.Symbol)

		if err != nil {
			u.logger.Info("error get position by userId and symbol")
			return err
		}

		if positions == nil {
			u.logger.Info("symbol don't save in user wallet, iniciate save position")
			positions, err = u.SaveNewPosition(ctx, order)

			if err != nil {
				u.logger.Info("error on update position by order")
				return err
			}
		}

		switch order.Side {
		case BUY:
			p := positions.UpdatePositionByPurchaceOrder(*order)
			err = u.pS.UpdatePositionByOrder(ctx, order, p)
			if err != nil {
				u.logger.Info("error on update position by order")
				return err
			}

			u.logger.Info("update order status for process sucessully")
			err = u.UpdateStatusOrders(ctx, order.ID)

			if err != nil {
				u.logger.Info("error on update order status")
				return err
			}
		case SELL:
			if positions.TotalAmount < order.Amount {
				u.logger.Error("sales quantity less than total quantity")
				return errors.New("sales quantity less than total quantity")
			}
			p := positions.UpdatePositionBySellOrder(*order)
			err = u.pS.UpdatePositionByOrder(ctx, order, p)

			if err != nil {
				u.logger.Info("error on update position by order")
				return err
			}

			u.logger.Info("update order status for process sucessully")

			err = u.UpdateStatusOrders(ctx, order.ID)

			if err != nil {
				u.logger.Info("error on update order status")
				return err
			}

		default:
			return errors.New("unsupported action")
		}
	}

	u.logger.Info("Process all orders completed")

	return nil
}

func (u *OrdersUsecase) SaveNewPosition(ctx context.Context, order *entity.Order) (*entity.Positions, error) {
	positions, err := u.pS.SaveNewPositionByOrder(ctx, order)

	if err != nil {
		u.logger.Info("error save new position")
		return nil, err
	}

	return positions, nil
}
