package model

import "github.com/ONSdigital/dp-frontend-models/model"

//Page contains data re-used for each page type a Data struct for data specific to the page type
type Page struct {
	model.Page
	Data Homepage `json:"data"`
}

//Homepage contains data specific to this page type
type Homepage struct {
	MainFigures []MainFigure `json:"key_figures"`
	Releases    []Release    `json:"releases"`
	Featured    []Feature    `json:"featured"`
	AroundONS   []Feature    `json:"arounds_ons"`
}

//Release is the data for an individual release
type Release struct {
	Title       string `json:"title"`
	URI         string `json:"uri"`
	ReleaseDate string `json:"releaseDate"`
}

//MainFigure is the data for an individual timeseries
type MainFigure struct {
	Title            string     `json:"title"`
	Summary          string     `json:"summary"`
	Date             string     `json:"date"`
	Figure           string     `json:"figure"`
	Trend            Trend      `json:"trend"`
	TrendDescription string     `json:"trend_description"`
	Unit             string     `json:"unit"`
	FigureURIs       FigureURIs `json:"figure_uris"`
}

// FigureURIs struct contains URI's to related analysis and data
type FigureURIs struct {
	Analysis string `json:"analysis"`
	Data     string `json:"data"`
}

// Trend contains bool values about the trend of a key figure (up from last month, down from last month, etc.)
type Trend struct {
	IsUp   bool `json:"is_up"`
	IsDown bool `json:"is_down"`
	IsFlat bool `json:"is_flat"`
}

//Feature is data for linked content
type Feature struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
	URI     string `json:"uri"`
}
