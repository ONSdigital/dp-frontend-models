package model

// EmergencyBanner data
type EmergencyBanner struct {
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URI         string `json:"uri"`
	LinkText    string `json:"link_text"`
}
