package request

// Request Assuming Request is an interface
type Request interface {
	GetApiMethodName() string
	GetApiParas() map[string]interface{}
	SetResponseError(ErrorResp)
	GetResponseError() ErrorResp
	SetResponseData(string) error
	GetResponseData(interface{}) error
}

type ErrorResponse struct {
	ErrorResp ErrorResp `json:"error_response"`
}

type ErrorResp struct {
	Code      string `json:"code"`
	EnDesc    string `json:"en_desc"`
	RequestID string `json:"request_id"`
	ZhDesc    string `json:"zh_desc"`
}

type ApiParams interface {
	GetApiParas() map[string]interface{}
}
