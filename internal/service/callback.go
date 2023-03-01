package service

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Service) Callback(c *gin.Context) {
	code := time.Now().Unix()
	c.HTML(200, "index.tmpl", gin.H{
		"title": "Your order is confirmed!",
		"code":  fmt.Sprintf("%d", code),
	})
}
