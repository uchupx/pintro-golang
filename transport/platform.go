package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uchupx/pintro-golang/data"
)

type PlatformHandler struct {
	PlatformRepository data.PlatformRepository
}

func (h PlatformHandler) Get(c *gin.Context) {
	platforms, err := h.PlatformRepository.FindAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, platforms)
	return
}
