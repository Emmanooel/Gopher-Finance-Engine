package routes

import (
	"gopher-finance-engine/internal/domain/application/usecases"
	"gopher-finance-engine/internal/infra/web/routes/handlers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router  *gin.Engine
	Handler *handlers.Handlers
}

func NewServer(
	userUsecase usecases.UserUsecasesI,
	positionUsecase usecases.PositionUsecasesI,
	orderUsecase usecases.OrdersUsecaseI,
) *Server {
	engine := gin.Default()

	handlers := handlers.NewHandlers(userUsecase, positionUsecase, orderUsecase)

	server := &Server{
		Router:  engine,
		Handler: handlers,
	}

	server.Router = Routes(engine, *server.Handler)
	return server

}

func (s *Server) StartServer(addr string) error {
	return s.Router.Run(addr)
}
