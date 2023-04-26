package model

type Region struct {
	Id         uint64 `json:"id" db:"id"`
	RegionName string `json:"region_name" db:"region_name"`
}
