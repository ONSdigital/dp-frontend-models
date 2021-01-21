package search

import "github.com/ONSdigital/dp-frontend-models/model"

// Page represents the search page
type Page struct {
	model.Page
	Data Search `json:"data"`
}

// Search represents all search parameters and response data of the search
type Search struct {
	Query    string                  `json:"query"`
	Filter   Filter                  `json:"filter,omitempty"`
	Sort     Sort                    `json:"sort,omitempty"`
	Limit    int                     `json:"limit,omitempty"`
	Offset   int                     `json:"offset,omitempty"`
	Category map[string][]FilterType `json:"category"`
	Response Response                `json:"response"`
}

// Filter represents all the information of filter related to the search page
type Filter struct {
	Query   string   `json:"query,omitempty"`
	Options []string `json:"options,omitempty"`
}

// Sort represents all the information of sorting related to the search page
type Sort struct {
	Query      string `json:"query,omitempty"`
	Text       string `json:"text,omitempty"`
	FilterText string `json:"filter_text,omitempty"`
}

// FilterType informs the name of the search type displayed on the website, the query retrieved from renderer and all the subtypes to pass to the logic
type FilterType struct {
	LocaliseKeyName string `json:"localise_key"`
	QueryType       string `json:"query_type"`
}

// Response represents the search results
type Response struct {
	Count        int           `json:"count"`
	ContentTypes []ContentType `json:"content_types"`
	Items        []ContentItem `json:"items"`
	Suggestions  []string      `json:"suggestions,omitempty"`
}

// ContentType represents the type of the search results and the number of results for each type
type ContentType struct {
	Type  string `json:"type"`
	Count int    `json:"count"`
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
