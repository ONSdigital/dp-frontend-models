package filterOverview

import "github.com/ONSdigital/dp-frontend-models/model"

// Page ...
type Page struct {
	model.Page
	Data     FilterOverview `json:"data"`
	FilterID string         `json:"filter_id"`
}

// FilterOverview ...
type FilterOverview struct {
	Dimensions                 []Dimension   `json:"dimensions"`
	UnsetDimensions            []string      `json:"unset_dimensions"`
	PreviewAndDownload         Link          `json:"preview_and_download"`
	PreviewAndDownloadDisabled bool          `json:"preview_and_download_disabled"`
	ClearAll                   Link          `json:"clear_all"`
	Cancel                     Link          `json:"cancel"`
	IsLatestVersion            bool          `json:"is_latest_version"`
	LatestVersion              LatestVersion `json:"latest_version"`
	DatasetTitle               string        `json:"dataset_title"`
	HasError                   bool          `json:"has_error"`
}

// Dimension ...
type Dimension struct {
	Filter          string   `json:"filter"`
	AddedCategories []string `json:"added_categories"`
	Link            Link     `json:"link"`
	HasNoCategory   bool     `json:"has_no_category"`
}

// Link ...
type Link struct {
	URL   string `json:"url"`
	Label string `json:"label"`
}

// LatestVersion ...
type LatestVersion struct {
	DatasetLandingPageURL          string `json:"dataset_landing_page_url"`
	FilterJourneyWithLatestJourney string `json:"filter_journey_with_latest_version"`
}
