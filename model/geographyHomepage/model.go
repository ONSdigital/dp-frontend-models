package geographyHomepage

import "github.com/ONSdigital/dp-frontend-models/model"

// Page contains the data re-used on each page as well as the data for the current page
type Page struct {
	model.Page
	Data GeographyLandingPage `json:"data"`
}

// GeographyLandingPage represents the data on the Geography Landing page
type GeographyLandingPage struct {
	AreaTypes []AreaType `json:"items"`
}

// AreaType represents the Type of data of the Geography page
type AreaType struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}
