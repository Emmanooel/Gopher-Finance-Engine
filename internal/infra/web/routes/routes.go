package routes

import (
	"gopher-finance-engine/internal/infra/auth"
	"gopher-finance-engine/internal/infra/web/routes/handlers"
	middleware "gopher-finance-engine/internal/infra/web/routes/middlewares"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, h handlers.Handlers, tokenService auth.TokenService) *gin.Engine {
	healthGroup := router.Group("/health")
	{
		healthGroup.GET("/health", h.HealthCheck)
	}

	userGroup := router.Group("/user")
	{
		userGroup.POST("/create", h.CreateUser)
		userGroup.POST("/login", h.Login)
		userGroup.GET("/me", middleware.AuthMiddleware(tokenService), h.GetUserById)

	}

	appsAuth := router.Group("/v1")
	{
		appsAuth.POST("/orders", h.CreateOrders)
		appsAuth.GET("/portfolio/:id", h.GetPositionByUserId)
	}
	return router
}
