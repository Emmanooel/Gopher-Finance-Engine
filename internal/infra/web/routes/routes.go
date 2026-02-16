package routes

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, server *Server) *gin.Engine {
	healthGroup := router.Group("/health")
	{
		healthGroup.GET("/health", server.HealthCheck)
	}

	userGroup := router.Group("/user")
	{
		userGroup.POST("/create", server.CreateUser)
		userGroup.POST("/login", server.Login)

	}

	appsAuth := router.Group("/v1")
	{
		appsAuth.POST("/orders", server.CreateOrders)
		appsAuth.GET("/portfolio/:id", server.GetPositionByUserId)
	}
	return router
}
