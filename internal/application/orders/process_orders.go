package orders

import (
	"context"
	"gopher-finance-engine/internal/domain/entity"

	"go.uber.org/zap"
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
		u.logger.Error("error get orders in pending by userId", zap.String("userId", userId), zap.Error(err))
		return err
	}

	stockQuantity := u.getQuantityBySymbol(pendingOrders)
	stockPrice := u.getPriceBySymbol(pendingOrders)

	u.logger.Info("Process orders", zap.Any("stockQuantity", stockQuantity), zap.Any("stockPrice", stockPrice))

	return nil
}

func (u *OrdersUsecase) getQuantityBySymbol(orders []*entity.Order) map[string]int64 {
	quantityBySymbol := make(map[string]int64)
	for _, order := range orders {
		quantityBySymbol[order.Symbol] += order.Amount
	}

	u.logger.Info("Quantity by symbol", zap.Any("quantityBySymbol", quantityBySymbol))
	return quantityBySymbol
}

func (u *OrdersUsecase) getPriceBySymbol(orders []*entity.Order) map[string]int64 {
	priceBySymbol := make(map[string]int64)
	for _, order := range orders {
		priceBySymbol[order.Symbol] += order.Price
	}
	u.logger.Info("Price by symbol", zap.Any("priceBySymbol", priceBySymbol))
	return priceBySymbol
}
