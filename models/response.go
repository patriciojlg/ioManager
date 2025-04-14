package models

type Response struct {
	StatusCode  int    `json:"status_code"`
	Body        any    `json:"body"`
	Error       bool   `json:"error"`
	DetailError string `json:"detail_error,omitempty"`
}

func Error400Response(err error) Response {
	return Response{
		StatusCode:  400,
		Body:        nil,
		Error:       true,
		DetailError: err.Error(),
	}
}
