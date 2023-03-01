package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mispon/digiseller-express/internal/service"
)

func main() {
	app := gin.Default()
	app.LoadHTMLGlob("templates/*")

	svc := service.New()

	app.GET("/ping", svc.Ping)
	app.GET("/callback", svc.Callback)

	log.Fatal(app.Run())
}
