package model

// Pagination represents all information regarding pagination of search results
type Pagination struct {
	CurrentPage    int    `json:"current_page,omitempty"`
	TotalPages     int    `json:"total_pages,omitempty"`
	URLWithoutPage string `json:"url_without_page,omitempty"`
	Limit          int    `json:"limit,omitempty"`
	LimitOptions   []int  `json:"limit_options,omitempty"`
}
