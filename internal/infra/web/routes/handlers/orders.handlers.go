package handlers

import (
	"gopher-finance-engine/internal/domain/entity"

	"github.com/gin-gonic/gin"
)

func (s *Handlers) CreateOrders(c *gin.Context) {
	ctx := c.Request.Context()
	var body *entity.Order

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, err.Error())
	}

	err := s.OrderUsecase.CreateOrders(ctx, body)
	if err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(202, "success")
}
