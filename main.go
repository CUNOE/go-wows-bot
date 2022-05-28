package main

import (
	"github.com/gin-gonic/gin"
	"go-wows-bot/wows"
)

func main() {
	r := gin.Default()
	r.GET("/ws", wows.ReadMessage)
	r.Run(":9000")
}
