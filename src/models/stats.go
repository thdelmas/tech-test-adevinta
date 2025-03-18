package models

// StatsResponse represents the response for statistics
type StatsResponse struct {
	MostFrequentRequest FizzBuzzRequest `json:"most_frequent_request"`
	Hits                int             `json:"hits"`
}
