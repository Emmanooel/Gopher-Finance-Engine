package engine

import (
	"context"
	"gopher-finance-engine/internal/application/orders"
	"time"

	"go.uber.org/zap"
)

type WorkerI interface {
	Start(ctx context.Context)
}

type Worker struct {
	logger       *zap.Logger
	OrderUsecase orders.OrdersUsecaseI
}

func NewWorker(
	logger *zap.Logger,
	OrderUsecase orders.OrdersUsecaseI,

) WorkerI {
	return &Worker{
		logger:       logger,
		OrderUsecase: OrderUsecase,
	}
}

func (w *Worker) Start(ctx context.Context) {
	w.logger.Info("Worker started")

	w.process(ctx)
	ticker := time.NewTicker(15 * time.Second)

	for {
		select {
		case <-ctx.Done():
			w.logger.Info("Context done, stopping worker")
			return
		case <-ticker.C:
			w.logger.Info("Processing pending orders")
			w.process(ctx)
		}
	}
}

func (w *Worker) process(ctx context.Context) {
	w.logger.Info("Processing pending orders")

	if err := w.OrderUsecase.ProcessPendingOrders(ctx); err != nil {
		w.logger.Error("Error processing pending orders", zap.Error(err))
		return
	}

	w.logger.Info("Finished processing pending orders")
}
