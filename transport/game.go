package transport

import (
	"net/http"
	"strconv"

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
		c.JSON(http.StatusBadRequest, err)
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

func (h GameHandler) Post(c *gin.Context) {
	// var responses []payload.ResponseData
	var request payload.GamePostRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	game := model.Game{
		GameName: request.Name,
		GenreId:  request.Genre,
	}

	_, err = h.GameRepository.Insert(c, game)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "success")
	return
}

func (h GameHandler) Put(c *gin.Context) {
	var request payload.GamePostRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	gameIdStr := c.Param("id")

	gameId, err := strconv.Atoi(gameIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	games, err := h.GameRepository.FindByIds(c, []uint64{uint64(gameId)})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if len(games) == 0 {
		c.JSON(http.StatusBadRequest, "game not found")
		return
	}

	game := games[0]

	game.GameName = request.Name
	game.GenreId = request.Genre

	_, err = h.GameRepository.Update(c, game)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "success")
	return
}

func (h GameHandler) Delete(c *gin.Context) {
	gameIdStr := c.Param("id")
	gameId, err := strconv.Atoi(gameIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	games, err := h.GameRepository.FindByIds(c, []uint64{uint64(gameId)})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if len(games) == 0 {
		c.JSON(http.StatusBadRequest, "game not found")
		return
	}

	game := games[0]

	_, err = h.GameRepository.Delete(c, game)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "success")
	return
}
