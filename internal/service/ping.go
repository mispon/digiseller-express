package service

import "github.com/gin-gonic/gin"

func (s *Service) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
