package response

type CommonResponse struct {
	code    uint16
	message string
	data    interface{}
}

func NewResponse(code uint16, message string, data interface{}) *CommonResponse {
	return &CommonResponse{
		code:    code,
		message: message,
		data:    data,
	}
}
