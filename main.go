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
	genreHandler := trans.GetGenreHandler(conf)
	publisherHandler := trans.GetPublisherHandler(conf)
	platformHandler := trans.GetPlatformHandler(conf)
	regionHandler := trans.GetRegionHandler(conf)
	userHandler := trans.GetUserHandler(conf)

	middlware := trans.GetMiddleware(conf)

	router := gin.Default()

	auth := router.Group("/", middlware.Authorization)
	auth.POST("/games", gameHandler.Post)
	auth.PUT("/games/:id", gameHandler.Put)
	auth.DELETE("/games/:id", gameHandler.Delete)

	router.POST("/users", userHandler.Post)
	router.POST("/users/sign-in", userHandler.Login)

	router.GET("/games", gameHandler.Get)
	router.GET("/genres", genreHandler.Get)
	router.GET("/publishers", publisherHandler.Get)
	router.GET("/platforms", platformHandler.Get)
	router.GET("/regions", regionHandler.Get)

	router.Run(":8081")
}
