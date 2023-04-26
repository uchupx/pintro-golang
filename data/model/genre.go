package model

type Genre struct {
	Id        uint64 `json:"id" db:"id"`
	GenreName string `json:"genre_name" db:"genre_name"`
}
