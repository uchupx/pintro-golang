package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uchupx/pintro-golang/data"
)

type GenreHandler struct {
	GenreRepository data.GenreRepository
}

func (h GenreHandler) Get(c *gin.Context) {
	genres, err := h.GenreRepository.FindAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, genres)
	return
}
