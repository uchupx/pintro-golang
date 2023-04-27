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
	var publisherIds []uint64

	for _, game := range games {
		gameIds = append(gameIds, game.Data.(model.Game).Id)
	}

	reqRelations, shouldLoadRelations := mapRelations["publisher"]

	if !shouldLoadRelations {
		for _, mapRelation := range mapRelations {
			if mapRelation == "publisher" {
				shouldLoadRelations = true
			}
		}
	}

	if shouldLoadRelations {
		var responses []ResponseData
		responsesMap := make(map[uint64]ResponseData)
		gamePublisherMap := make(map[uint64]uint64)

		gamesPublishers, err := g.GamePublisherRepository.FindByGameIds(ctx, gameIds)
		if err != nil {
			return nil, err
		}

		for _, item := range gamesPublishers {
			publisherIds = append(publisherIds, item.PublisherId)
			gamePublisherMap[item.GameId] = item.PublisherId
		}

		publishers, err := g.PublisherRepository.FindByIds(ctx, publisherIds)
		if err != nil {
			return nil, err
		}

		for _, item := range publishers {
			response := ResponseData{
				Id:     item.Id,
				Entity: "publisher",
				Data:   item,
			}

			responsesMap[item.Id] = response
			responses = append(responses, response)
		}

		if reqRelationsMap, ok := reqRelations.(map[string]interface{}); ok {
			relations, err := g.GamePublisherResponseGenerator.LoadRelations(ctx, responses, reqRelationsMap)

			if err != nil {
				return nil, err
			}

			for idx, response := range responsesMap {
				relation, exists := relations[response.Id]

				if exists {
					response.Relations = &relation
				}

				responsesMap[idx] = response
			}
		}

		for _, item := range games {
			relations[item.Id]["publisher"] = map[string][]model.Publisher{"data": nil}
		}

		for _, item := range games {
			if publisherId, ok := gamePublisherMap[item.Data.(model.Game).Id]; ok {
				relations[item.Id]["publisher"] = responsesMap[publisherId]
			}
		}
	}
	return relations, nil
}

// func (g GameResponseGenerator) loadPlatformRelations(ctx context.Context, games []ResponseData, mapRelations RelationsType, relations GroupedRelations) (GroupedRelations, error) {
// 	var gameIds []uint64
// 	var platformIds []uint64
// 	var publisherIds []uint64

// 	for _, game := range games {
// 		gameIds = append(gameIds, game.Data.(model.Game).Id)
// 	}

// 	_, shouldLoadRelations := mapRelations["platform"]

// 	if !shouldLoadRelations {
// 		for _, mapRelation := range mapRelations {
// 			if mapRelation == "platform" {
// 				shouldLoadRelations = true
// 			}
// 		}
// 	}

// 	if shouldLoadRelations {
// 		responses := make(map[uint64]ResponseData)
// 		gamePublisherMap := make(map[uint64]uint64)
// 		gamePlatformMap := make(map[uint64]uint64)

// 		gamesPublishers, err := g.GamePublisherRepository.FindByGameIds(ctx, gameIds)
// 		if err != nil {
// 			return nil, err
// 		}

// 		for _, item := range gamesPublishers {
// 			publisherIds = append(publisherIds, item.Id)
// 			gamePublisherMap[item.GameId] = item.Id
// 		}

// 		gamePlatforms, err := g.GamePlatformRepository.FindByPublisherIds(ctx, publisherIds)
// 		if err != nil {
// 			return nil, err
// 		}

// 		for _, item := range gamePlatforms {
// 			platformIds = append(platformIds, item.PlatformId)
// 			gamePlatformMap[item.GamePublisherId] = item.PlatformId
// 		}

// 		platforms, err := g.PlatformRepository.FindByIds(ctx, platformIds)
// 		if err != nil {
// 			return nil, err
// 		}

// 		for _, item := range platforms {
// 			response := ResponseData{
// 				Id:     item.Id,
// 				Entity: "platform",
// 				Data:   item,
// 			}

// 			responses[item.Id] = response
// 		}

// 		for _, item := range games {
// 			relations[item.Id]["platform"] = map[string][]model.Platform{"data": nil}
// 		}

// 		for _, item := range games {
// 			if publisherId, ok := gamePublisherMap[item.Data.(model.Game).Id]; ok {
// 				if platformId, ok := gamePlatformMap[publisherId]; ok {
// 					relations[item.Id]["platform"] = responses[platformId]
// 				}
// 			}
// 		}
// 	}
// 	return relations, nil
// }
