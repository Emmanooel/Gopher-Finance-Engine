package worker

import (
	"context"
	"gopher-finance-engine/internal/domain/application/usecases"
	"gopher-finance-engine/internal/domain/entity"
	"log"
	"sync"

	"go.uber.org/zap"
)

type WorkerSaveNewPositionI interface {
	WorkerSavePositionByNewOrder(ctx context.Context, o *entity.Order)
}

type WorkerSaveNewPosition struct {
	l        *zap.Logger
	pUsecase usecases.PositionUsecasesI
}

func NewWorkerSaveNewPosition(
	logger *zap.Logger,
	pUsecase usecases.PositionUsecasesI,
) WorkerSaveNewPositionI {
	return &WorkerSaveNewPosition{
		l:        logger,
		pUsecase: pUsecase,
	}
}

func (w *WorkerSaveNewPosition) WorkerSavePositionByNewOrder(ctx context.Context, o *entity.Order) {
	w.l.Info("initalize worker")
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		log.Println(*o)
		defer wg.Done()
		err := w.pUsecase.SavePositionByNewOrder(context.TODO(), o)

		if err != nil {
			w.l.Error("error worker save new position")
			return
		}
	}()
	wg.Wait()

	w.l.Info("worker done")
}
