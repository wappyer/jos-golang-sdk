package request

import (
	"encoding/json"
)

/*
 * jingdong.eclp.trace.service.jos.OrderTr aceByOrderService
 * 商家按订单号查询物流轨迹接口
 * 用于商家通过订单号(沧海系统订 单号)查询全部的物流轨迹(包括: 清关，仓库生产，配送运输)(需要 额外申请权限)
 */

type EclpTraceServiceJosOrderTraceByOrderServiceRequest struct {
	apiParas map[string]interface{}
	Version  string `json:"version,omitempty"`
	OrderId  string `json:"orderId,omitempty"`
	Role     string `json:"role,omitempty"`
	UserId   string `json:"userId,omitempty"`
}

func NewEclpTraceServiceJosOrderTraceByOrderServiceRequest() *EclpTraceServiceJosOrderTraceByOrderServiceRequest {
	r := &EclpTraceServiceJosOrderTraceByOrderServiceRequest{
		apiParas: make(map[string]interface{}),
	}
	r.Version = "1.0"
	return r
}

func (r *EclpTraceServiceJosOrderTraceByOrderServiceRequest) GetApiMethodName() string {
	return "jingdong.eclp.trace.service.jos.OrderTraceByOrderService"
}

func (r *EclpTraceServiceJosOrderTraceByOrderServiceRequest) GetApiParas() string {
	if len(r.apiParas) == 0 {
		return "{}"
	}
	data, _ := json.Marshal(r.apiParas)
	return string(data)
}

func (r *EclpTraceServiceJosOrderTraceByOrderServiceRequest) Check() {
	// 实现检查逻辑
}

func (r *EclpTraceServiceJosOrderTraceByOrderServiceRequest) PutOtherTextParam(key string, value interface{}) {
	r.apiParas[key] = value
}

func (r *EclpTraceServiceJosOrderTraceByOrderServiceRequest) SetVersion(version string) {
	r.Version = version
	r.apiParas["version"] = version
}

func (r *EclpTraceServiceJosOrderTraceByOrderServiceRequest) GetVersion() string {
	return r.Version
}

func (r *EclpTraceServiceJosOrderTraceByOrderServiceRequest) SetOrderId(orderId string) {
	r.OrderId = orderId
	r.apiParas["orderId"] = orderId
}

func (r *EclpTraceServiceJosOrderTraceByOrderServiceRequest) GetOrderId() string {
	return r.OrderId
}

func (r *EclpTraceServiceJosOrderTraceByOrderServiceRequest) SetRole(role string) {
	r.Role = role
	r.apiParas["role"] = role
}

func (r *EclpTraceServiceJosOrderTraceByOrderServiceRequest) GetRole() string {
	return r.Role
}

func (r *EclpTraceServiceJosOrderTraceByOrderServiceRequest) SetUserId(userId string) {
	r.UserId = userId
	r.apiParas["userId"] = userId
}

func (r *EclpTraceServiceJosOrderTraceByOrderServiceRequest) GetUserId() string {
	return r.UserId
}
