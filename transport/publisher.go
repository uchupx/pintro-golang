package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uchupx/pintro-golang/data"
)

type PublisherHandler struct {
	PublisherRepository data.PublisherRepository
}

func (h PublisherHandler) Get(c *gin.Context) {
	genres, err := h.PublisherRepository.FindAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, genres)
	return
}
