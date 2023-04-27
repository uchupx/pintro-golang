package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uchupx/pintro-golang/data"
	"github.com/uchupx/pintro-golang/data/model"
	"github.com/uchupx/pintro-golang/helper"
	"github.com/uchupx/pintro-golang/transport/payload"
)

type GameHandler struct {
	GameRepository          data.GameRepository
	GamePublisherRepository data.GamePublisherRepository
	GameResponseGenerator   *payload.GameResponseGenerator
}

func (h GameHandler) Get(c *gin.Context) {
	var responses []payload.ResponseData
	var request payload.GameRequest

	err := shouldBind(c, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	query := data.GameQuery{
		PerPage: helper.UintOptional(request.PerPage, 30),
		Page:    helper.UintOptional(request.Page, 1),
	}

	games, err := h.GameRepository.FindByQuery(c, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	for _, item := range games.Data.([]model.Game) {
		responses = append(responses, payload.ResponseData{
			Id:        item.Id,
			Entity:    "game",
			Data:      item,
			Relations: nil,
		})
	}

	if request.Relations != nil {
		relations, err := h.GameResponseGenerator.LoadRelations(c, responses, *request.Relations)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		for idx, item := range responses {
			relation, exists := relations[item.Id]

			if exists {
				item.Relations = &relation
			}

			responses[idx] = item
		}
	}

	c.JSON(http.StatusOK, CollectionsResponse{
		Perpage: query.PerPage,
		Page:    query.Page,
		Data:    responses,
	})
	return
}
