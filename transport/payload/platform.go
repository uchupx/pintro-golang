package payload

import (
	"context"

	"github.com/uchupx/pintro-golang/data"
	"github.com/uchupx/pintro-golang/data/model"
)

type PlatfromResponseGenerator struct {
	PlatformRepository    data.PlatformRepository
	RegionSalesRepository data.RegionSalesRepository
	RegionRepository      data.RegionRepository
}

func (g PlatfromResponseGenerator) LoadRelations(ctx context.Context, platforms []ResponseData, mapRelations RelationsType) (GroupedRelations, error) {
	relations := make(GroupedRelations)

	for _, platform := range platforms {
		relations[platform.Id] = make(map[string]interface{})
	}

	if mapRelations == nil {
		return nil, nil
	}

	relations, err := g.loadPlatformRelations(ctx, platforms, mapRelations, relations)
	if err != nil {
		return nil, err
	}

	relations, err = g.loadRegionSalesRelations(ctx, platforms, mapRelations, relations)
	if err != nil {
		return nil, err
	}

	return relations, nil
}

func (g PlatfromResponseGenerator) loadPlatformRelations(ctx context.Context, gamePlatforms []ResponseData, mapRelations RelationsType, relations GroupedRelations) (GroupedRelations, error) {
	var platformIds []uint64

	for _, item := range gamePlatforms {
		platformIds = append(platformIds, item.Data.(model.GamePlatform).PlatformId)
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
		var responses []ResponseData

		items, err := g.PlatformRepository.FindByIds(ctx, platformIds)

		if err != nil {
			return nil, err
		}

		for _, item := range items {
			response := ResponseData{
				Id:     item.Id,
				Entity: "platform",
				Data:   item,
			}

			responses = append(responses, response)
		}

		for _, item := range gamePlatforms {
			relations[item.Id]["platform"] = map[string][]model.Platform{}
		}

		for _, item := range gamePlatforms {
			for _, res := range responses {
				if item.Data.(model.GamePlatform).PlatformId == res.Id {
					relations[item.Id]["platform"] = res
				}
			}
		}
	}
	return relations, nil
}

func (g PlatfromResponseGenerator) loadRegionSalesRelations(ctx context.Context, publisher []ResponseData, mapRelations RelationsType, relations GroupedRelations) (GroupedRelations, error) {
	var gamePlatformIDs []uint64

	for _, game := range publisher {
		gamePlatformIDs = append(gamePlatformIDs, game.Data.(model.GamePlatform).Id)
	}

	_, shouldLoadRelations := mapRelations["sales"]

	if !shouldLoadRelations {
		for _, mapRelation := range mapRelations {
			if mapRelation == "sales" {
				shouldLoadRelations = true
			}
		}
	}

	if shouldLoadRelations {
		var responses []ResponseData

		gamePlatforms, err := g.RegionRepository.FindByIds(ctx, gamePlatformIDs)
		if err != nil {
			return nil, err
		}

		for _, item := range gamePlatforms {
			response := ResponseData{
				Id:     item.Id,
				Entity: "sales",
				Data:   item,
			}

			responses = append(responses, response)
		}

		for _, item := range publisher {
			relations[item.Id]["sales"] = map[string][]model.Platform{"data": nil}
		}

		for _, item := range publisher {
			for _, res := range responses {
				if item.Id == res.Data.(model.RegionJoinRegionSales).GamePlatformId {
					relations[item.Id]["sales"] = res
				}
			}
		}
	}
	return relations, nil
}
