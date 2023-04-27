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

type GameResponseGenerator struct {
	GenreRepository data.GenreRepository
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
