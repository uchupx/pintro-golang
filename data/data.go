package data

type Collection struct {
	PerPage uint64
	Total   uint64
	// Limit   uint64
	Data interface{}
}
