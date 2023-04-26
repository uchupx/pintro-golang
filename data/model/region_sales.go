package model

type RegionSales struct {
	RegionId       uint64 `json:"region_id" db:"region_id"`
	GamePlatformId uint64 `json:"game_platform_id" db:"game_platform_id"`
	NumSales       uint64 `json:"num_sales" db:"num_sales"`
}
