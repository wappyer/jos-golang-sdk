package jd

// Request Assuming Request is an interface
type Request interface {
	GetApiMethodName() string
	GetApiParas() map[string]interface{}
}
