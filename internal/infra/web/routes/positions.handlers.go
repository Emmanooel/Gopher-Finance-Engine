package routes

import "github.com/gin-gonic/gin"

func (s *Server) GetPositionByUserId(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	response, err := s.PositionUsecase.GetPositionByUserId(ctx, id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": response})
}
