package area

import "github.com/ONSdigital/dp-frontend-models/model"

// Page represents the template data structure used for the geography list page
type Page struct {
	model.Page
	Data GeographyAreaPage `json:"data"`
}

// GeographyAreaPage represents the data specific to an area page
type GeographyAreaPage struct {
}
