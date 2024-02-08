package pagination

type PageInfo struct {
	PageNumber int    `json:"pageNumber"`
	PageSize   int    `json:"pageSize"`
	SortBy     string `json:"sortBy"`
	SortType   string `json:"sortType"`
}

type Page[T any] struct {
	Elements   *[]T `json:"elements"`
	PageNumber int  `json:"pageNumber"`
	PageSize   int  `json:"pageSize"`
}
