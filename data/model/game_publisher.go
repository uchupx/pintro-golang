package model

type GamePublisher struct {
	Id          uint64 `json:"id" db:"id"`
	GameId      uint64 `json:"game_id" db:"game_id"`
	PublisherId uint64 `json:"publisher_id" db:"publisher_id"`
}
