package models

// FizzBuzzRequest represents the parameters for a fizzbuzz request
type FizzBuzzRequest struct {
	Int1  int    `json:"int1"`
	Int2  int    `json:"int2"`
	Limit int    `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
}

// FizzBuzzResponse represents the response for a fizzbuzz request
type FizzBuzzResponse struct {
	Result []string `json:"result"`
}

// StatsResponse represents the response for statistics
type StatsResponse struct {
	MostFrequentRequest FizzBuzzRequest `json:"most_frequent_request"`
	Hits                int             `json:"hits"`
}
