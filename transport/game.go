package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uchupx/pintro-golang/data"
)

type GameHandler struct {
	GameRepository data.GameRepository
}

func (h GameHandler) Get(c *gin.Context) {

	games, err := h.GameRepository.FindByQuery(c, data.GameQuery{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, games)
	return
}
