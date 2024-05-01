package got

type Direction string

const (
	DirectionAsc  Direction = "asc"
	DirectionDesc Direction = "desc"
)

type PaginationCursor struct {
	Column    string       `json:"column"`
	Direction Direction    `json:"direction"`
	Value     *interface{} `json:"value"`
}
