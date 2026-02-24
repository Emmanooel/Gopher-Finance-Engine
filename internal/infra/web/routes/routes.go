package routes

import (
	"gopher-finance-engine/internal/infra/web/routes/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, h handlers.Handlers) *gin.Engine {
	healthGroup := router.Group("/health")
	{
		healthGroup.GET("/health", h.HealthCheck)
	}

	userGroup := router.Group("/user")
	{
		userGroup.POST("/create", h.CreateUser)
		userGroup.POST("/login", h.Login)

	}

	appsAuth := router.Group("/v1")
	{
		appsAuth.POST("/orders", h.CreateOrders)
		appsAuth.GET("/portfolio/:id", h.GetPositionByUserId)
	}
	return router
}
