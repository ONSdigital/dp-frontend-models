package hierarchy

import "github.com/ONSdigital/dp-frontend-models/model"

// Page ...
type Page struct {
	model.Page
	Data     Hierarchy `json:"data"`
	FilterID string    `json:"filter_id"`
}

// Hierarchy ...
type Hierarchy struct {
	Title         string   `json:"title"`
	SaveAndReturn Link     `json:"save_and_return"`
	Cancel        Link     `json:"cancel"`
	FiltersAmount string   `json:"filters_amount"`
	AddAllFilters AddAll   `json:"add_all"`
	FilterList    []List   `json:"filter_list"`
	FiltersAdded  []Filter `json:"filters_added"`
	RemoveAll     Link     `json:"remove_all"`
	GoBack        Link     `json:"go_back"`
	Parent        string   `json:"parent"`
}

// AddAll ...
type AddAll struct {
	Amount string `json:"amount"`
	URL    string `json:"url"`
}

// Filter ...
type Filter struct {
	Label     string `json:"label"`
	RemoveURL string `json:"remove_url"`
}

// List ...
type List struct {
	Label   string `json:"label"`
	SubNum  string `json:"sub_num"`
	SubType string `json:"sub_type"`
	SubURL  string `json:"sub_url"`
}

// Link ...
type Link struct {
	URL   string `json:"url"`
	Label string `json:"label"`
}
