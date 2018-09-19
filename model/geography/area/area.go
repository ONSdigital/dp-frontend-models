package area

import "github.com/ONSdigital/dp-frontend-models/model"

// Page represents the template data structure used for the geography list page
type Page struct {
	model.Page
	Data GeographyAreaPage `json:"data"`
}

// GeographyAreaPage represents the data specific to an area page
type GeographyAreaPage struct {
	Datasets   []Dataset `json:"items"`
	Attributes Attribute `json:"attributes"`
}

// Dataset represents each dataset related to specific code
type Dataset struct {
	Label       string `json:"label"`
	Description string `json:"description"`
	ID          string `json:"id"`
	URI         string `json:"uri"`
}

// Attribute represents each attribute related to specific code
type Attribute struct {
	Code        string `json:"code"`
	ReleaseDate string `json:"release_date"`
}
