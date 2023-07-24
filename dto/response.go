package dto

import "net/http"

// Response represents the API response format
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var NotImplementedResponse = Response{
	Status:  http.StatusNotImplemented,
	Message: "not implemented",
	Data:    nil,
}

var NotFoundResponse = Response{
	Status:  http.StatusNotFound,
	Message: "not found",
	Data:    nil,
}

var InvalidIdResponse = Response{
	Status:  http.StatusBadRequest,
	Message: "invalid id",
	Data:    nil,
}