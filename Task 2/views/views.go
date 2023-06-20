package views

// error message to be shown in case any error occurs
type ErrorResponse struct {
	Error string `json:"error"`
}

// success message to be shown in case of success of the operation desired
type SuccessResponse struct {
	Message string `json:"message"`
}
