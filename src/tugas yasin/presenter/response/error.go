package response

type StatusResponse struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

type ResponseError struct {
	Status StatusResponse `json:"status"`
}

type ResponseSuccess struct {
	Status StatusResponse `json:"status"`
	Data   interface{}    `json:"data"`
}
