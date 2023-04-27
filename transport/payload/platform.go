package payload

import (
	"context"

	"github.com/uchupx/pintro-golang/data"
)

type PlatfromResponseGenerator struct {
	RegionSales data.RegionSalesRepository
	Region      data.RegionRepository
}

func (g PlatfromResponseGenerator) LoadRelations(ctx context.Context, platforms []ResponseData, mapRelations RelationsType) (GroupedRelations, error) {
	return GroupedRelations{}, nil
}

// func (g PlatfromResponseGenerator) loadSalesRelations(ctx context.Context, platforms []ResponseData, mapRelations RelationsType, relations GroupedRelations) (GroupedRelations, error) {
// 	var platformIds []uint64
// 	var salesIds []uint64

// 	for _, game := range platforms {
// 		platformIds = append(platformIds, game.Data.(model.Game).Id)
// 	}

// 	_, shouldLoadRelations := mapRelations["publisher"]

// 	if !shouldLoadRelations {
// 		for _, mapRelation := range mapRelations {
// 			if mapRelation == "publisher" {
// 				shouldLoadRelations = true
// 			}
// 		}
// 	}

// 	if shouldLoadRelations {
// 		responses := make(map[uint64]ResponseData)
// 		gamePublisherMap := make(map[uint64]uint64)

// 		gamesPublishers, err := g.GamePublisherRepository.FindByGameIds(ctx, platformIds)
// 		if err != nil {
// 			return nil, err
// 		}

// 		for _, item := range gamesPublishers {
// 			salesIds = append(salesIds, item.PublisherId)
// 			gamePublisherMap[item.GameId] = item.PublisherId
// 		}

// 		publishers, err := g.PublisherRepository.FindByIds(ctx, salesIds)
// 		if err != nil {
// 			return nil, err
// 		}

// 		for _, item := range publishers {
// 			response := ResponseData{
// 				Id:     item.Id,
// 				Entity: "publisher",
// 				Data:   item,
// 			}

// 			responses[item.Id] = response
// 		}

// 		for _, item := range platforms {
// 			relations[item.Id]["publisher"] = map[string][]model.Publisher{"data": nil}
// 		}

// 		for _, item := range platforms {
// 			if publisherId, ok := gamePublisherMap[item.Data.(model.Game).Id]; ok {
// 				relations[item.Id]["publisher"] = responses[publisherId]
// 			}
// 		}
// 	}
// 	return relations, nil
// }
