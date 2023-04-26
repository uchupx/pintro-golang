package model

type GamePlatform struct {
	Id              uint64 `json:"id" db:"id"`
	GamePublisherId uint64 `json:"game_publisher_id" db:"game_publisher_id"`
	PlatformId      uint64 `json:"platform_id" db:"platform_id"`
	ReleaseYear     uint64 `json:"release_year" db:"release_year"`
}
