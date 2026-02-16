package routes

import (
	"gopher-finance-engine/internal/domain/application/usecases"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router          *gin.Engine
	UserUsecase     usecases.UserUsecasesI
	PositionUsecase usecases.PositionUsecasesI
	OrderUsecase    usecases.OrdersUsecaseI
}

func NewServer(
	userUsecase usecases.UserUsecasesI,
	positionUsecase usecases.PositionUsecasesI,
	orderUsecase usecases.OrdersUsecaseI,
) *Server {
	engine := gin.Default()

	server := &Server{
		Router:          engine,
		UserUsecase:     userUsecase,
		PositionUsecase: positionUsecase,
		OrderUsecase:    orderUsecase,
	}

	server.Router = Routes(engine, server)
	return server

}

func (s *Server) StartServer(addr string) error {
	return s.Router.Run(addr)
}
