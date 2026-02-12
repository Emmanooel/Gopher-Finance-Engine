package routes

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, server *Server) *gin.Engine {
	router.GET("/health", server.HealthCheck)
	return router
}
