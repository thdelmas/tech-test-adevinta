package models

// FizzBuzzRequest represents the parameters for a fizzbuzz request
type FizzBuzzRequest struct {
	Int1  int    `form:"int1"`
	Int2  int    `form:"int2"`
	Limit int    `form:"limit"`
	Str1  string `form:"str1"`
	Str2  string `form:"str2"`
}

// FizzBuzzResponse represents the response for a fizzbuzz request
type FizzBuzzResponse struct {
	Result []string `json:"result"`
}
