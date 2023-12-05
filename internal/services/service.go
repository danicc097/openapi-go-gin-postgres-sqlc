package services

type PaginationParamsDirection string

const (
	PaginationParamsDirectionAsc  = "asc"
	PaginationParamsDirectionDesc = "desc"
)

type PaginationParams struct {
	cursor    string
	limit     int
	direction PaginationParamsDirection
}
