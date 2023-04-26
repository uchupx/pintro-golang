package model

type Game struct {
	Id       uint64 `json:"id" db:"id"`
	GenreId  uint64 `json:"genre_id" db:"genre_id"`
	GameName string `json:"game_name" db:"game_name"`
}
