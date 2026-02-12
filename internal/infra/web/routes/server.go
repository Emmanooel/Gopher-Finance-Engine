package routes

import "github.com/gin-gonic/gin"

type Server struct {
	Router *gin.Engine
	//usecases
}

func NewServer() *Server {
	engine := gin.Default()

	server := &Server{
		Router: engine,
	}

	server.Router = Routes(engine, server)
	return server

}

func (s *Server) StartServer(addr string) error {
	return s.Router.Run(addr)
}
