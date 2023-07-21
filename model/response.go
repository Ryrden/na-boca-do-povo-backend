package model

import "net/http"

// Response represents the API response format
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var NotImplementedResponse = Response{
	Status:  http.StatusNotImplemented,
	Message: "Not Implemented",
	Data:    nil,
}
