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
	GameRepository         data.GameRepository
	GamePlatformRepository data.GamePlatformRepository
	PlatformRepository     data.PlatformRepository
}

func (g GamePublisherResponseGenerator) LoadRelations(ctx context.Context, publisher []ResponseData, mapRelations RelationsType) (GroupedRelations, error) {
	relations := make(GroupedRelations)

	for _, game := range publisher {
		relations[game.Id] = make(map[string]interface{})
	}

	if mapRelations == nil {
		return nil, nil
	}

	// relations, err := g.loadGameRelations(ctx, publisher, mapRelations, relations)
	// if err != nil {
	// 	return nil, err
	// }

	relations, err := g.loadPlatformRelations(ctx, publisher, mapRelations, relations)
	if err != nil {
		return nil, err
	}

	return relations, nil
}

// func (g GamePublisherResponseGenerator) loadGameRelations(ctx context.Context, publisher []ResponseData, mapRelations RelationsType, relations GroupedRelations) (GroupedRelations, error) {
// 	var genreIds []uint64

// 	for _, game := range publisher {
// 		genreIds = append(genreIds, game.Data.(model.Publisher).GameId)
// 	}

// 	_, shouldLoadRelations := mapRelations["game"]

// 	if !shouldLoadRelations {
// 		for _, mapRelation := range mapRelations {
// 			if mapRelation == "game" {
// 				shouldLoadRelations = true
// 			}
// 		}
// 	}

// 	if shouldLoadRelations {
// 		responses := make(map[uint64]ResponseData)

// 		genres, err := g.GameRepository.FindByIds(ctx, genreIds)

// 		if err != nil {
// 			return nil, err
// 		}

// 		for _, genre := range genres {
// 			response := ResponseData{
// 				Id:     genre.Id,
// 				Entity: "game",
// 				Data:   genre,
// 			}

// 			responses[genre.Id] = response
// 		}

// 		for _, item := range publisher {
// 			relations[item.Id]["game"] = map[string][]model.Game{}
// 		}

// 		for _, item := range publisher {
// 			if response, ok := responses[item.Data.(model.GamePublisher).GameId]; ok {
// 				relations[item.Id]["game"] = response
// 			}
// 		}
// 	}
// 	return relations, nil
// }

func (g GamePublisherResponseGenerator) loadPlatformRelations(ctx context.Context, publisher []ResponseData, mapRelations RelationsType, relations GroupedRelations) (GroupedRelations, error) {
	var platformIds []uint64
	var publisherIds []uint64

	for _, game := range publisher {
		publisherIds = append(publisherIds, game.Data.(model.Publisher).Id)
	}

	_, shouldLoadRelations := mapRelations["platform"]

	if !shouldLoadRelations {
		for _, mapRelation := range mapRelations {
			if mapRelation == "platform" {
				shouldLoadRelations = true
			}
		}
	}

	if shouldLoadRelations {
		responses := make(map[uint64]ResponseData)
		gamePlatformMap := make(map[uint64]uint64)

		gamePlatforms, err := g.GamePlatformRepository.FindByPublisherIds(ctx, publisherIds)
		if err != nil {
			return nil, err
		}

		for _, item := range gamePlatforms {
			platformIds = append(platformIds, item.PlatformId)
			gamePlatformMap[item.GamePublisherId] = item.PlatformId
		}

		platforms, err := g.PlatformRepository.FindByIds(ctx, platformIds)
		if err != nil {
			return nil, err
		}

		for _, item := range platforms {
			response := ResponseData{
				Id:     item.Id,
				Entity: "platform",
				Data:   item,
			}

			responses[item.Id] = response
		}

		for _, item := range publisher {
			relations[item.Id]["platform"] = map[string][]model.Platform{"data": nil}
		}

		for _, item := range publisher {
			if platformId, ok := gamePlatformMap[item.Data.(model.Publisher).Id]; ok {
				relations[item.Id]["platform"] = responses[platformId]
			}
		}
	}
	return relations, nil
}
