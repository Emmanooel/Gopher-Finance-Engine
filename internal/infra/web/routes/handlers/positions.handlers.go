package handlers

import "github.com/gin-gonic/gin"

func (s *Handlers) GetPositionByUserId(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.GetString("user_id")

	if id == "" {
		c.JSON(400, gin.H{"error": "user_id is required"})
		return
	}

	response, err := s.PositionUsecase.GetPositionByUserId(ctx, id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": response})
}
