package datasetPage

import "github.com/ONSdigital/dp-frontend-models/model"

//Page contains data re-used for each page type a Data struct for data specific to the page type
type Page struct {
	model.Page
	DatasetPage DatasetPage `json:"data"`
	model.ContactDetails
}

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
	Versions            []Version           `json:"versions"`
	SupplementaryFiles  []SupplementaryFile `json:"supplementary_files"`
	Downloads           []Download          `json:"downloads"`
	IsNationalStatistic bool                `json:"national_statistic"`
	ReleaseDate         string              `json:"releaseDate"`
	NextRelease         string              `json:"nextRelease"`
	DatasetID           string              `json:"datasetId"`
	URI                 string              `json:"uri"`
	Edition             string              `json:"edition"`
	Markdown            string              `json:"markdown"`
}

//Download has the details for the an individual dataset's downloadable files
type Download struct {
	Extension string `json:"extension"`
	Size      string `json:"size"`
	URI       string `json:"uri"`
	File      string `json:"file"`
}

//SupplementaryFile is a downloadable file that is associated to an individual dataset
type SupplementaryFile struct {
	Title     string `json:"title"`
	Extension string `json:"extension"`
	Size      string `json:"size"`
	URI       string `json:"uri"`
}

type Version struct {
	URI              string     `json:"url"`
	UpdateDate       string     `json:"updateDate"`
	CorrectionNotice string     `json:"correctionNotice"`
	Label            string     `json:"label"`
	Downloads        []Download `json:"downloads"`
}
