package routes

import "github.com/gin-gonic/gin"

func (s *Server) HealthCheck(c *gin.Context) {

	c.JSON(200, gin.H{"status": "ok"})
}
