package main

import (
	"github.com/gin-gonic/gin"
	config "github.com/uchupx/pintro-golang/config"
	transport "github.com/uchupx/pintro-golang/transport"
)

func main() {
	conf := config.GetConfig()
	trans := transport.Transport{}
	gameHandler := trans.GetGameHandler(conf)
	router := gin.Default()

	// router.POST("/transactions", transactionHandler.Posts)
	router.GET("/games", gameHandler.Get)
	router.Run(":8081")
}
