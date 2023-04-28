package payload

import (
	"context"

	"github.com/uchupx/pintro-golang/data"
	"github.com/uchupx/pintro-golang/data/model"
)

type GameRequest struct {
	PerPage   uint64         `form:"per_page"`
	Page      uint64         `form:"page"`
	Relations *RelationsType `form:"relations"`
}

type GamePostRequest struct {
	Name      string `json:"name"`
	Platform  uint64 `json:"platform"`
	Genre     uint64 `json:"genre"`
	Publisher uint64 `json:"publisher"`
}

type GameResponseGenerator struct {
	GenreRepository                data.GenreRepository
	GamePublisherRepository        data.GamePublisherRepository
	GamePlatformRepository         data.GamePlatformRepository
	PublisherRepository            data.PublisherRepository
	PlatformRepository             data.PlatformRepository
	GamePublisherResponseGenerator GamePublisherResponseGenerator
}

func (g GameResponseGenerator) LoadRelations(ctx context.Context, games []ResponseData, mapRelations RelationsType) (GroupedRelations, error) {
	relations := make(GroupedRelations)

	for _, game := range games {
		relations[game.Id] = make(map[string]interface{})
	}

	if mapRelations == nil {
		return nil, nil
	}

	relations, err := g.loadGenreRelations(ctx, games, mapRelations, relations)
	if err != nil {
		return nil, err
	}

	relations, err = g.loadPublisherRelations(ctx, games, mapRelations, relations)
	if err != nil {
		return nil, err
	}

	return relations, nil
}

func (g GameResponseGenerator) loadGenreRelations(ctx context.Context, games []ResponseData, mapRelations RelationsType, relations GroupedRelations) (GroupedRelations, error) {
	var genreIds []uint64

	for _, game := range games {
		genreIds = append(genreIds, game.Data.(model.Game).GenreId)
	}

	_, shouldLoadRelations := mapRelations["genre"]

	if !shouldLoadRelations {
		for _, mapRelation := range mapRelations {
			if mapRelation == "genre" {
				shouldLoadRelations = true
			}
		}
	}

	if shouldLoadRelations {
		responses := make(map[uint64]ResponseData)

		genres, err := g.GenreRepository.FindByIds(ctx, genreIds)

		if err != nil {
			return nil, err
		}

		for _, genre := range genres {
			response := ResponseData{
				Id:     genre.Id,
				Entity: "genre",
				Data:   genre,
			}

			responses[genre.Id] = response
		}

		for _, item := range games {
			relations[item.Id]["genre"] = map[string][]model.Genre{"data": nil}
		}

		for _, item := range games {
			if response, ok := responses[item.Data.(model.Game).GenreId]; ok {
				relations[item.Id]["genre"] = response
			}
		}
	}
	return relations, nil
}

func (g GameResponseGenerator) loadPublisherRelations(ctx context.Context, games []ResponseData, mapRelations RelationsType, relations GroupedRelations) (GroupedRelations, error) {
	var gameIds []uint64

	for _, game := range games {
		gameIds = append(gameIds, game.Data.(model.Game).Id)
	}

	reqRelations, shouldLoadRelations := mapRelations["game_publisher"]

	if !shouldLoadRelations {
		for _, mapRelation := range mapRelations {
			if mapRelation == "game_publisher" {
				shouldLoadRelations = true
			}
		}
	}

	if shouldLoadRelations {
		var responses []ResponseData

		gamesPublishers, err := g.GamePublisherRepository.FindByGameIds(ctx, gameIds)
		if err != nil {
			return nil, err
		}

		for _, item := range gamesPublishers {
			response := ResponseData{
				Id:     item.Id,
				Entity: "game_publisher",
				Data:   item,
			}

			responses = append(responses, response)
		}

		if reqRelationsMap, ok := reqRelations.(map[string]interface{}); ok {
			relations, err := g.GamePublisherResponseGenerator.LoadRelations(ctx, responses, reqRelationsMap)

			if err != nil {
				return nil, err
			}

			for idx, response := range responses {
				relation, exists := relations[response.Id]

				if exists {
					response.Relations = &relation
				}

				responses[idx] = response
			}
		}

		for _, item := range games {
			relations[item.Id]["game_publisher"] = map[string][]model.Publisher{"data": nil}
		}

		for _, item := range games {
			for _, res := range responses {
				if item.Id == res.Data.(model.GamePublisher).GameId {
					relations[item.Id]["game_publisher"] = res
				}
			}
		}
	}
	return relations, nil
}
