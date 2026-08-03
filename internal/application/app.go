package application

import (
	"context"
	"gopher-finance-engine/configs"

	"gopher-finance-engine/internal/application/orders"
	"gopher-finance-engine/internal/application/positions"
	positions_service "gopher-finance-engine/internal/application/positions/service"
	"gopher-finance-engine/internal/application/users"
	"gopher-finance-engine/internal/infra/auth"
	orders_repository "gopher-finance-engine/internal/infra/repository/orders"
	positions_repository "gopher-finance-engine/internal/infra/repository/positions"
	users_repository "gopher-finance-engine/internal/infra/repository/users"
	"gopher-finance-engine/internal/infra/web/routes"
	"gopher-finance-engine/pkg/postgres"

	"go.uber.org/zap"
)

type Application struct {
	Logger      *zap.Logger
	Routes      *routes.Server
	authService auth.TokenService
	usecases    Usecases
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

	app.authService = auth.NewAuthService()

	app.Routes = routes.NewServer(
		app.Logger,
		app.usecases.UserUsecase,
		app.usecases.PositionsUsecase,
		app.usecases.OrderUsecase,
		app.authService,
	)

	return &app
}

func newUsecases(app *Application) Usecases {
	userRepository := users_repository.NewUserRepository(app.Logger)
	positionRepository := positions_repository.NewPositionRepository(app.Logger)
	orderRepository := orders_repository.NewOrdersRepository(app.Logger)

	userUsecase := users.NewUsersUsecase(app.Logger, userRepository, app.authService)
	positionUsecase := positions.NewPositionUsecase(app.Logger, positionRepository, orderRepository)
	positionService := positions_service.NewPositionService(app.Logger, positionUsecase)

	orderUsecase := orders.NewOrdersUsecase(app.Logger, orderRepository, positionService)

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
