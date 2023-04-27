package payload

type RelationsType map[string]interface{}
type GroupedRelations map[uint64]RelationsType

type ResponseData struct {
	Id        uint64         `json:"id"`
	Entity    string         `json:"entity"`
	Data      interface{}    `json:"data"`
	Relations *RelationsType `json:"relations"`
}
