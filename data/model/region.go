package model

type Region struct {
	Id         uint64 `json:"id" db:"id"`
	RegionName string `json:"region_name" db:"region_name"`
}

type RegionJoinRegionSales struct {
	Id             uint64  `json:"id" db:"id"`
	RegionName     string  `json:"region_name" db:"region_name"`
	GamePlatformId uint64  `json:"game_platform_id" db:"game_platform_id"`
	NumSales       float64 `json:"num_sales" db:"num_sales"`
}
