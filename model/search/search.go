package search

import "github.com/ONSdigital/dp-frontend-models/model"

// Page ...
type Page struct {
	model.Page
	Data Search `json:"data"`
}

// Search ..
type Search struct {
	Query    string   `json:"query"`
	Filter   []string `json:"filter,omitempty"`
	Sort     string   `json:"sort,omitempty"`
	Limit    int      `json:"limit,omitempty"`
	Offset   int      `json:"offset,omitempty"`
	Response Response `json:"response"`
}

// Response ...
type Response struct {
	Count        int           `json:"count"`
	ContentTypes []contentType `json:"content_types"`
	Items        []contentItem `json:"items"`
	Suggestions  []string      `json:"suggestions,omitempty"`
}

type contentType struct {
	Type  string `json:"type"`
	Count int    `json:"count"`
}

type contentItem struct {
	Description description `json:"description"`
	Type        string      `json:"type"`
	URI         string      `json:"uri"`
	Matches     *matches    `json:"matches,omitempty"`
}

type description struct {
	Contact           *contact  `json:"contact,omitempty"`
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

type contact struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone,omitempty"`
	Email     string `json:"email"`
}

type matches struct {
	Description struct {
		Summary         *[]matchDetails `json:"summary"`
		Title           *[]matchDetails `json:"title"`
		Edition         *[]matchDetails `json:"edition,omitempty"`
		MetaDescription *[]matchDetails `json:"meta_description,omitempty"`
		Keywords        *[]matchDetails `json:"keywords,omitempty"`
		DatasetID       *[]matchDetails `json:"dataset_id,omitempty"`
	} `json:"description"`
}

type matchDetails struct {
	Value string `json:"value,omitempty"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}
