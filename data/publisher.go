package data

type Publisher struct {
	Id            uint64 `json:"id" db:"id"`
	PublisherName string `json:"publisher_name" db:"publisher_name"`
}
