package handlers

import (
	"context"
	"gopher-finance-engine/internal/domain/entity"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s *Handlers) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var body *entity.Users

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := s.UserUsecase.CreateUser(ctx, body)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(202, gin.H{"message": "omg deu certo!"})
}

func (s *Handlers) Login(c *gin.Context) {
	ctx := c.Request.Context()
	var body *entity.UserLogin

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response, err := s.UserUsecase.Login(ctx, *body)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"access_token": response,
	})
}

func (s *Handlers) GetUsers(c *gin.Context) {
	ctx := c.Request.Context()
	p := c.Query("page")

	page, _ := strconv.Atoi(p)

	response, err := s.UserUsecase.GetAllUsers(ctx, page)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, response)

}

func (s *Handlers) GetUserById(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.GetString("user_id")

	s.logger.Info("GetUserById", zap.String("user_id", id))

	if id == "" {
		c.JSON(400, gin.H{"error": "user_id is required"})
		return
	}

	response, err := s.UserUsecase.GetUserById(ctx, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, response)
}
