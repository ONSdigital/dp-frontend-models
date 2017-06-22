package datasetLandingPageStatic

import "github.com/ONSdigital/dp-frontend-models/model"

//Page contains data re-used for each page type a Data struct for data specific to the page type
type Page struct {
	model.Page
	DatasetLandingPage DatasetLandingPage `json:"data"`
	model.ContactDetails
}

//DatasetLandingPage ...
type DatasetLandingPage struct {
	DatasetID           string    `json:"datasetID"`
	FilterID            string    `json:"filterID"`
	RelatedPublications []Related `json:"relatedPublications"`
	RelatedDatasets     []Related `json:"relatedDatasets"`
	RelatedMethodology  []Related `json:"relatedMethodology"`
	RelatedLinks        []Related `json:"relatedLinks"`
	Datasets            []Dataset `json:"datasets"`
	Notes               string    `json:"markdown"`
	MetaDescription     string    `json:"metaDescription"`
	NationalStatistic   bool      `json:"nationalStatistic"`
	ReleaseDate         bool      `json:"releaseDate"`
	NextRelease         bool      `json:"nextRelease"`
	IsTimeseries        bool      `json:"timeseries"`
}

//Related page details
type Related struct {
	Title string `json:"title"`
	URI   string `json:"uri"`
}

//Dataset has the file and title information for an individual dataset
type Dataset struct {
	Title                 string              `json:"title"`
	Downloads             []Download          `json:"downloads"`
	URI                   string              `json:"uri"`
	HasVersions           bool                `json:"hasVersions"`
	HasSupplementaryFiles bool                `json:"hasSupplementaryFiles"`
	SupplementaryFiles    []SupplementaryFile `json:"supplementaryFiles"`
}

//Download has the details for the an individual dataset's downloadable files
type Download struct {
	Title     string `json:"title"`
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
