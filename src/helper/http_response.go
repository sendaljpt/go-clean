package helper

type HttpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}

func Response(code int, message string, data, errors interface{}) HttpResponse {
	res := HttpResponse{}
	res.Code = code
	res.Message = message
	res.Data = data
	res.Errors = errors

	return res
}
