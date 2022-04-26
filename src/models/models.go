package models

type Holiday struct {
	Date        string `json:"date"`
	LocalName   string `json:"localName"`
	Name        string `json:"name"`
	CountryCode string `json:"countryCode"`
	Fixed       bool   `json:"fixed"`
	Global      bool   `json:"global"`
	Counties    bool   `json:"counties"`
	LaunchYear  bool   `json:"launchYear"`
	Yype        string `json:"type"`
}
