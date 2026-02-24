package handlers

import "github.com/gin-gonic/gin"

func (s *Handlers) HealthCheck(c *gin.Context) {

	c.JSON(200, gin.H{"status": "ok"})
}
