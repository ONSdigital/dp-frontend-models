package model

// Pagination represents all information regarding pagination of search results
type Pagination struct {
	CurrentPage    int             `json:"current_page"`
	PagesToDisplay []PageToDisplay `json:"pages_to_display"`
	TotalPages     int             `json:"total_pages"`
	URLWithoutPage string          `json:"url_without_page"`
	Limit          int             `json:"limit"`
	LimitOptions   []int           `json:"limit_options,omitempty"`
}

// PageToDisplay represents a page to display in pagination with their corresponding URL
type PageToDisplay struct {
	Page int    `json:"page"`
	URL  string `json:"url"`
}
