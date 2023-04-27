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
	router := gin.Default()

	router.GET("/games", gameHandler.Get)
	// router.POST("/games", gameHandler.Get)
	// router.PUT("/games", gameHandler.Get)
	// router.DELETE("/games", gameHandler.Get)
	router.GET("/games/publisher", gameHandler.Get)
	// router.GET("/games/platform", gameHandler.Get)

	router.GET("/genres", genreHandler.Get)
	// router.POST("/genres", genreHandler.Get)
	// router.PUT("/genres", genreHandler.Get)
	// router.DELETE("/genres", genreHandler.Get)

	router.GET("/publishers", publisherHandler.Get)
	// router.POST("/publishers", publisherHandler.Get)
	// router.PUT("/publishers", publisherHandler.Get)
	// router.DELETE("/publishers", publisherHandler.Get)

	router.GET("/platforms", platformHandler.Get)
	// router.POST("/platforms", publisherHandler.Get)
	// router.PUT("/platforms", publisherHandler.Get)
	// router.DELETE("/platforms", publisherHandler.Get)

	router.GET("/regions", regionHandler.Get)
	router.GET("/regions/sales", publisherHandler.Get)

	router.Run(":8081")
}
