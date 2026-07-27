package orders

import (
	"context"
	"errors"
	"gopher-finance-engine/internal/domain/entity"

	"go.uber.org/zap"
)

const (
	BUY  = "BUY"
	SELL = "SELL"
)

//step 1: obtem todas as ordens do usuário com status PENDING ::check
//step 2: verifica se a ordem é de compra ou venda ::check
//step 3: se for de compra, obter quantidade comprada e valor total da compra ::check
//step 4: atualizar a posição do usuário
//step 5: atualizar preço medio do usuário
//step 6: atualizar status da ordem para COMPLETED
//step 7: se for de venda, verificar se o usuário possui a quantidade de ações para vender
//step 8: atualizar a posição do usuário
//step 9: atualizar PM da posição do usuário
//step 10: cria alerta de taxa
//step 11: atualizar status da ordem para COMPLETED

func (u *OrdersUsecase) ProcessOrders(ctx context.Context, userId string) error {
	u.logger.Info("Iniciate process orders", zap.String("userId", userId))

	pendingOrders, err := u.repo.GetOrdersInPendingByUserId(ctx, userId)
	if err != nil {
		u.logger.Info("error get orders in pending by userId", zap.String("userId", userId))
		return err
	}

	for _, order := range pendingOrders {
		positions, err := u.pS.SearchPositionByUserIdAndSymbol(ctx, userId, order.Symbol)

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
