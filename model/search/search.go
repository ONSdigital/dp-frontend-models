package search

import "github.com/ONSdigital/dp-frontend-models/model"

// Page represents the search page
type Page struct {
	model.Page
	Data Search `json:"data"`
}

// Search represents all search parameters and response data of the search
type Search struct {
	Query    string   `json:"query"`
	Filter   []string `json:"filter,omitempty"`
	Sort     Sort     `json:"sort,omitempty"`
	Limit    Limit    `json:"limit,omitempty"`
	Offset   Offset   `json:"offset,omitempty"`
	Response Response `json:"response"`
}

// Sort represents all the information of sorting related to the search page
type Sort struct {
	Query              string        `json:"query,omitempty"`
	LocaliseFilterKeys []string      `json:"filter_text,omitempty"`
	LocaliseSortKey    string        `json:"sort_text,omitempty"`
	Options            []SortOptions `json:"options,omitempty"`
}

// SortOptions represents all the information of different sorts available
type SortOptions struct {
	Query           string `json:"query,omitempty"`
	LocaliseKeyName string `json:"localise_key"`
}

// Limit represents the number of results to be shown in one page and all the possible limit options
type Limit struct {
	Query   int   `json:"query,omitempty"`
	Options []int `json:"options,omitempty"`
}

// Offset represents
type Offset struct {
	Page           int    `json:"page,omitempty"`
	TotalPages     int    `json:"total_pages,omitempty"`
	URLWithoutPage string `json:"url_without_page,omitempty"`
}

// Response represents the search results
type Response struct {
	Count       int           `json:"count"`
	Categories  []Category    `json:"categories"`
	Items       []ContentItem `json:"items"`
	Suggestions []string      `json:"suggestions,omitempty"`
}

// Category represents all the search categories in search page
type Category struct {
	Count           int           `json:"count"`
	LocaliseKeyName string        `json:"localise_key"`
	ContentTypes    []ContentType `json:"content_types"`
}

// ContentType represents the type of the search results and the number of results for each type
type ContentType struct {
	Type            string `json:"type"`
	Count           int    `json:"count"`
	LocaliseKeyName string `json:"localise_key"`
}

// ContentItem represents each search result
type ContentItem struct {
	Description Description `json:"description"`
	Type        string      `json:"type"`
	URI         string      `json:"uri"`
	Matches     *Matches    `json:"matches,omitempty"`
}

// Description represents each search result description
type Description struct {
	Contact           *Contact  `json:"contact,omitempty"`
	DatasetID         string    `json:"dataset_id,omitempty"`
	Edition           string    `json:"edition,omitempty"`
	Headline1         string    `json:"headline1,omitempty"`
	Headline2         string    `json:"headline2,omitempty"`
	Headline3         string    `json:"headline3,omitempty"`
	Keywords          *[]string `json:"keywords,omitempty"`
	LatestRelease     *bool     `json:"latest_release,omitempty"`
	Language          string    `json:"language,omitempty"`
	MetaDescription   string    `json:"meta_description,omitempty"`
	NationalStatistic *bool     `json:"national_statistic,omitempty"`
	NextRelease       string    `json:"next_release,omitempty"`
	PreUnit           string    `json:"pre_unit,omitempty"`
	ReleaseDate       string    `json:"release_date,omitempty"`
	Source            string    `json:"source,omitempty"`
	Summary           string    `json:"summary"`
	Title             string    `json:"title"`
	Unit              string    `json:"unit,omitempty"`
}

// Contact represents each search result contact details
type Contact struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone,omitempty"`
	Email     string `json:"email"`
}

// Matches represents each search result matches
type Matches struct {
	Description MatchDescription `json:"description"`
}

// MatchDescription represents each search result matches' description
type MatchDescription struct {
	Summary         *[]MatchDetails `json:"summary"`
	Title           *[]MatchDetails `json:"title"`
	Edition         *[]MatchDetails `json:"edition,omitempty"`
	MetaDescription *[]MatchDetails `json:"meta_description,omitempty"`
	Keywords        *[]MatchDetails `json:"keywords,omitempty"`
	DatasetID       *[]MatchDetails `json:"dataset_id,omitempty"`
}

// MatchDetails represents each search result matches' details
type MatchDetails struct {
	Value string `json:"value,omitempty"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}
