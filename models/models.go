package models

type Task struct {
	ID      int
	Title   string
	Checked bool
}

// ErrorResponse is interface for sending error message with code.
type ErrorResponse struct {
	Code    int
	Message string
}
