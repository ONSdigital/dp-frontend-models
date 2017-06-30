package cmd

import "github.com/ONSdigital/dp-frontend-models/model"

// Page ...
type Page struct {
	model.Page
	Data FilterOverview `json:"data"`
}

// FilterOverview ...
type FilterOverview struct {
	Dimensions []Dimension `json:"dimensions"`
}

// Dimension ...
type Dimension struct {
	Filter          string     `json:"filter"`
	AddedCategories string     `json:"added_categories"`
	Link            FilterLink `json:"link"`
}

// FilterLink ...
type FilterLink struct {
	URL   string `json:"url"`
	Label string `json:"label"`
}
