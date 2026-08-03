package main

import (
	"context"
	"gopher-finance-engine/configs"
	"gopher-finance-engine/internal/application"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	configs.LoadConfigs()

	ctx := context.TODO()
	app := application.NewWorker()

	defer app.Logger.Sync()

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)

	defer stop()

	app.Worker.Start(ctx)
}
