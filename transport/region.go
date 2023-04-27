package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uchupx/pintro-golang/data"
)

type RegionHandler struct {
	RegionRepository data.RegionRepository
}

func (h RegionHandler) Get(c *gin.Context) {
	genres, err := h.RegionRepository.FindAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, genres)
	return
}
