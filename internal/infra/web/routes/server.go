package routes

import (
	"gopher-finance-engine/internal/application/orders"
	"gopher-finance-engine/internal/application/positions"
	"gopher-finance-engine/internal/application/users"
	"gopher-finance-engine/internal/infra/auth"
	"gopher-finance-engine/internal/infra/web/routes/handlers"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	Router  *gin.Engine
	Handler *handlers.Handlers
}

func NewServer(
	logger *zap.Logger,
	userUsecase users.UserUsecasesI,
	positionUsecase positions.PositionUsecasesI,
	orderUsecase orders.OrdersUsecaseI,
	tokenService auth.TokenService,
) *Server {
	engine := gin.Default()

	handlers := handlers.NewHandlers(logger, userUsecase, positionUsecase, orderUsecase)

	server := &Server{
		Router:  engine,
		Handler: handlers,
	}

	server.Router = Routes(engine, *server.Handler, tokenService)
	return server

}

func (s *Server) StartServer(addr string) error {
	return s.Router.Run(addr)
}
