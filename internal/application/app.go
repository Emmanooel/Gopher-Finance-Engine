package application

import (
	"context"
	"gopher-finance-engine/configs"
	"gopher-finance-engine/internal/application/usecases/orders"
	"gopher-finance-engine/internal/application/usecases/positions"
	"gopher-finance-engine/internal/application/usecases/users"
	"gopher-finance-engine/internal/infra/auth"
	"gopher-finance-engine/internal/infra/repository"
	"gopher-finance-engine/internal/infra/web/routes"
	"gopher-finance-engine/pkg/postgres"

	"go.uber.org/zap"
)

type Application struct {
	Logger   *zap.Logger
	Routes   *routes.Server
	usecases Usecases
}

type Usecases struct {
	UserUsecase      users.UserUsecasesI
	PositionsUsecase positions.PositionUsecasesI
	OrderUsecase     orders.OrdersUsecaseI
}

func NewApplication() *Application {
	var app Application

	postgres.NewPostgresConn(context.Background(), configs.DbConn)

	app.Logger = initializeLogger()
	app.usecases = newUsecases(&app)
	app.Routes = routes.NewServer(
		app.usecases.UserUsecase,
		app.usecases.PositionsUsecase,
		app.usecases.OrderUsecase,
	)

	return &app
}

func newUsecases(app *Application) Usecases {
	userRepository := repository.NewUserRepository(app.Logger)
	positionRepository := repository.NewPositionRepository(app.Logger)
	orderRepository := repository.NewOrdersRepository(app.Logger)

	authService := auth.NewAuthService()

	userUsecase := users.NewUsersUsecase(app.Logger, userRepository, authService)
	positionUsecase := positions.NewPositionUsecase(app.Logger, positionRepository, orderRepository)

	orderUsecase := orders.NewOrdersUsecase(app.Logger, orderRepository, positionUsecase)

	return Usecases{
		UserUsecase:      userUsecase,
		PositionsUsecase: positionUsecase,
		OrderUsecase:     orderUsecase,
	}
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
