package dto

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var NotImplementedResponse = Response{
	Success: false,
	Message: "not implemented",
	Data:    nil,
}

var NotFoundResponse = Response{
	Success: false,
	Message: "not found",
	Data:    nil,
}

var InvalidIdResponse = Response{
	Success: false,
	Message: "invalid id",
	Data:    nil,
}
