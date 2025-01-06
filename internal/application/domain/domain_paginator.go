package domain

type Paginator[T any] struct {
	Items      []T  `json:"items"`
	HasNext    bool `json:"hasNext"`
	NextCursor int  `json:"nextCursor"`
}
