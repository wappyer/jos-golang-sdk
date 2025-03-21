package jd

import "jos-golang-sdk/jd/request"

// Request Assuming Request is an interface
type Request interface {
	GetApiMethodName() string
	GetApiParas() map[string]interface{}
}

func NewRequest(requestName string) Request {
	switch requestName {
	case "EclpGoodsQueryGoodsRecordRequest":
		return request.NewEclpGoodsQueryGoodsRecordRequest()
	case "EclpGoodsQueryGoodsInfoRequest":
		return request.NewEclpGoodsQueryGoodsInfoRequest()
	default:
		return nil
	}
}
