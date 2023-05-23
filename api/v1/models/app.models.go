package models

type Response struct {
	StatusCode int         `json:"statusCode"`
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Response   interface{} `json:"response"`
}

func (r Response) CreateResponse(statusCode int, status string, message string, response interface{}) Response {
	return Response{StatusCode: statusCode, Status: status, Message: message, Response: response}
}

func CreateResponse(respCode int, respStatus, respMessage string, response interface{}) Response {
	return Response{StatusCode: respCode, Status: respStatus, Message: respMessage, Response: response}
}
