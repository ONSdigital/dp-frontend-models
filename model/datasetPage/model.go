package datasetPage

import "github.com/ONSdigital/dp-frontend-models/model"

//Page contains data re-used for each page type a Data struct for data specific to the page type
type Page struct {
	model.Page
	DatasetPage DatasetPage `json:"data"`
	model.ContactDetails
}

//DatasetLandingPage represents a frontend dataset landing page
/*
type DatasetPage struct {
	DatasetID           string    `json:"dataset_id"`
	FilterID            string    `json:"filter_id"`
	Related             Related   `json:"related"`
	Datasets            []Dataset `json:"datasets"`
	Notes               string    `json:"markdown"`
	MetaDescription     string    `json:"meta_description"`
	IsNationalStatistic bool      `json:"national_statistic"`
	ReleaseDate         string    `json:"release_date"`
	NextRelease         string    `json:"next_release"`
	IsTimeseries        bool      `json:"is_timeseries"`
	Corrections         []Message `json:"corrections"`
	Notices             []Message `json:"notices"`
	ParentPath          string    `json:"parent_path"`
}*/

//Related content (split by type) to this page
type Related struct {
	Publications       []model.Related `json:"related_publications"`
	FilterableDatasets []model.Related `json:"related_filterable_datasets"`
	Datasets           []model.Related `json:"related_datasets"`
	Methodology        []model.Related `json:"related_methodology"`
	Links              []model.Related `json:"related_links"`
}

//Dataset has the file and title information for an individual dataset
type DatasetPage struct {
	Downloads          []Download          `json:"downloads"`
	Versions           []Version           `json:"versions"`
	SupplementaryFiles []SupplementaryFile `json:"supplementary_files"`
}

//Download has the details for the an individual dataset's downloadable files
type Download struct {
	Extension string `json:"extension"`
	Size      string `json:"size"`
	URI       string `json:"uri"`
}

//SupplementaryFile is a downloadable file that is associated to an individual dataset
type SupplementaryFile struct {
	Title     string `json:"title"`
	Extension string `json:"extension"`
	Size      string `json:"size"`
	URI       string `json:"uri"`
}

type Version struct {
	URI              string `json:"url"`
	UpdateDate       string `json:"updateDate"`
	CorrectionNotice string `json:"correctionNotice"`
	Label            string `json:"label"`
}
