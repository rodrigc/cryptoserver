package main

// ErrorMsg is used to send error messages back to HTTP clients as JSON
type ErrorMsg struct {
	Msg string `json:"error"`
}
