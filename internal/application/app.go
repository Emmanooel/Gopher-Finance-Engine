package application

import (
	"gopher-finance-engine/configs"
	"gopher-finance-engine/internal/infra/web/routes"

	"go.uber.org/zap"
)

var (
	App *Application
)

type Application struct {
	Logger *zap.Logger
	// Add application fields here
}

func NewApplication() *routes.Server {
	App.Logger = initializeLogger()
	return routes.NewServer()
}

func initializeLogger() *zap.Logger {
	var logger *zap.Logger
	switch configs.App.Env {
	case "PROD":
		logger = zap.Must(zap.NewProduction())
	default:
		logger = zap.Must(zap.NewDevelopment())
	}

	return logger
}
