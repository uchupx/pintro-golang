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
	GameRepository            data.GameRepository
	GamePlatformRepository    data.GamePlatformRepository
	PlatformRepository        data.PlatformRepository
	PublisherRepository       data.PublisherRepository
	PlatfromResponseGenerator PlatfromResponseGenerator
}

func (g GamePublisherResponseGenerator) LoadRelations(ctx context.Context, publisher []ResponseData, mapRelations RelationsType) (GroupedRelations, error) {
	relations := make(GroupedRelations)

	for _, game := range publisher {
		relations[game.Id] = make(map[string]interface{})
	}

	if mapRelations == nil {
		return nil, nil
	}

	relations, err := g.loadPublisherRelations(ctx, publisher, mapRelations, relations)
	if err != nil {
		return nil, err
	}

	relations, err = g.loadGamePlatformRelations(ctx, publisher, mapRelations, relations)
	if err != nil {
		return nil, err
	}

	return relations, nil
}

func (g GamePublisherResponseGenerator) loadPublisherRelations(ctx context.Context, publishers []ResponseData, mapRelations RelationsType, relations GroupedRelations) (GroupedRelations, error) {
	var publishersIds []uint64

	for _, item := range publishers {
		publishersIds = append(publishersIds, item.Data.(model.GamePublisher).PublisherId)
	}

	_, shouldLoadRelations := mapRelations["publisher"]

	if !shouldLoadRelations {
		for _, mapRelation := range mapRelations {
			if mapRelation == "publisher" {
				shouldLoadRelations = true
			}
		}
	}

	if shouldLoadRelations {
		var responses []ResponseData

		items, err := g.PublisherRepository.FindByIds(ctx, publishersIds)

		if err != nil {
			return nil, err
		}

		for _, item := range items {
			response := ResponseData{
				Id:     item.Id,
				Entity: "publisher",
				Data:   item,
			}

			responses = append(responses, response)
		}

		for _, item := range publishers {
			relations[item.Id]["publisher"] = map[string][]model.Publisher{}
		}

		for _, item := range publishers {
			for _, res := range responses {
				if item.Data.(model.GamePublisher).PublisherId == res.Id {
					relations[item.Id]["publisher"] = res
				}
			}
		}
	}
	return relations, nil
}

func (g GamePublisherResponseGenerator) loadGamePlatformRelations(ctx context.Context, publisher []ResponseData, mapRelations RelationsType, relations GroupedRelations) (GroupedRelations, error) {
	var publisherIds []uint64

	for _, game := range publisher {
		publisherIds = append(publisherIds, game.Data.(model.GamePublisher).Id)
	}

	reqRelations, shouldLoadRelations := mapRelations["game_platform"]

	if !shouldLoadRelations {
		for _, mapRelation := range mapRelations {
			if mapRelation == "game_platform" {
				shouldLoadRelations = true
			}
		}
	}

	if shouldLoadRelations {
		var responses []ResponseData

		gamePlatforms, err := g.GamePlatformRepository.FindByPublisherIds(ctx, publisherIds)
		if err != nil {
			return nil, err
		}

		for _, item := range gamePlatforms {
			response := ResponseData{
				Id:     item.Id,
				Entity: "game_platform",
				Data:   item,
			}

			responses = append(responses, response)
		}

		if reqRelationsMap, ok := reqRelations.(map[string]interface{}); ok {
			relations, err := g.PlatfromResponseGenerator.LoadRelations(ctx, responses, reqRelationsMap)

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

		for _, item := range publisher {
			relations[item.Id]["game_platform"] = map[string][]model.Platform{"data": nil}
		}

		for _, item := range publisher {
			for _, res := range responses {
				if item.Id == res.Data.(model.GamePlatform).GamePublisherId {
					relations[item.Id]["game_platform"] = res
				}
			}
		}
	}
	return relations, nil
}
