package model

import (
	"github.com/ONSdigital/dp-frontend-models/model"
)

//Page contains data re-used for each page type a Data struct for data specific to the page type
type Page struct {
	model.Page
	Description Description `json:"description"`
}

// Item represents the Type of data of the Geography page
type Description struct {
	Title             string  `json:"title"`
	Summary           string  `json:"summary"`
	NationalStatistic bool    `json:"nationalStatistic"`
	Contact           Contact `json:"contact"`
	ReleaseDate       string  `json:"releaseDate"`
}

type Contact struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	// Telephone string `json:"telephone"` not using at the mo
}
