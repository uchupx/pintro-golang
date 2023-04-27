package payload

import (
	"context"

	"github.com/uchupx/pintro-golang/data"
	"github.com/uchupx/pintro-golang/data/model"
)

type GamePublisherRequest struct {
	PerPage   uint64         `form:"per_page"`
	Page      uint64         `form:"page"`
	Relations *RelationsType `form:"relations"`
}

type GamePublisherResponseGenerator struct {
	GameRepository data.GameRepository
}

func (g GamePublisherResponseGenerator) LoadRelations(ctx context.Context, games []ResponseData, mapRelations RelationsType) (GroupedRelations, error) {
	relations := make(GroupedRelations)

	for _, game := range games {
		relations[game.Id] = make(map[string]interface{})
	}

	if mapRelations == nil {
		return nil, nil
	}

	relations, err := g.loadGameRelations(ctx, games, mapRelations, relations)
	if err != nil {
		return nil, err
	}

	return relations, nil
}

func (g GamePublisherResponseGenerator) loadGameRelations(ctx context.Context, games []ResponseData, mapRelations RelationsType, relations GroupedRelations) (GroupedRelations, error) {
	var genreIds []uint64

	for _, game := range games {
		genreIds = append(genreIds, game.Data.(model.GamePublisher).GameId)
	}

	_, shouldLoadRelations := mapRelations["game"]

	if !shouldLoadRelations {
		for _, mapRelation := range mapRelations {
			if mapRelation == "game" {
				shouldLoadRelations = true
			}
		}
	}

	if shouldLoadRelations {
		responses := make(map[uint64]ResponseData)

		genres, err := g.GameRepository.FindByIds(ctx, genreIds)

		if err != nil {
			return nil, err
		}

		for _, genre := range genres {
			response := ResponseData{
				Id:     genre.Id,
				Entity: "game",
				Data:   genre,
			}

			responses[genre.Id] = response
		}

		for _, item := range games {
			relations[item.Id]["game"] = map[string][]model.Game{}
		}

		for _, item := range games {
			if response, ok := responses[item.Data.(model.GamePublisher).GameId]; ok {
				relations[item.Id]["game"] = response
			}
		}
	}
	return relations, nil
}
