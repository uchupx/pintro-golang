package data

type Platform struct {
	Id           uint64 `json:"id" db:"id"`
	PlatformName string `json:"platform_name" db:"platform_name"`
}
