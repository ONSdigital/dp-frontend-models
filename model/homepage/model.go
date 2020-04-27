package model

import "github.com/ONSdigital/dp-frontend-models/model"

//Page contains data re-used for each page type a Data struct for data specific to the page type
type Page struct {
	model.Page
	Data Homepage `json:"data"`
}

//Homepage contains data specific to this page type
type Homepage struct {
	KeyFigures []KeyFigure `json:"key_figures"`
	Releases   []Release   `json:"releases"`
	Featured   []Feature   `json:"featured"`
	Other      []Feature   `json:"other"`
}

//Release is the data for an individual release
type Release struct {
	Title       string `json:"title"`
	URI         string `json:"uri"`
	ReleaseDate string `json:"releaseDate"`
}

//KeyFigure is the data for an individual timeseries
type KeyFigure struct {
	Title         string         `json:"title"`
	Summary       string         `json:"summary"`
	ReleaseDate   string         `json:"release_date"`
	LatestFigures []LatestFigure `json:"latest_figure"`
}

//LatestFigure is the extra information displayed for the latest figure for a timeseries
type LatestFigure struct {
	Date             string           `json:"date"`
	Figure           string           `json:"figure"`
	Trend            Trend            `json:"trend"`
	TrendDescription string           `json:"trend_description"`
	Unit             string           `json:"unit"`
	LatestFigureURIs LatestFigureURIs `json:"latest_figure_uris"`
}

// LatestFigureURIs struct contains URI's to related analysis and data
type LatestFigureURIs struct {
	Analysis string `json:"analysis"`
	Data     string `json:"data"`
}

type Trend struct {
	isUp   bool `json:"is_up"`
	isDown bool `json:"is_down"`
	isFlat bool `json:"is_flat"`
}

//Feature is data for linked content
type Feature struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
	URI     string `json:"uri"`
}
