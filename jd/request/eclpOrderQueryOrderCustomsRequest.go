package request

import (
	"encoding/json"
)

/*
 * jingdong.eclp.order.queryOrderCustoms
 * 订单申报结果查询接口
 * 用于商家查询订单申报状态
 */

type EclpOrderQueryOrderCustomsRequest struct {
	apiParas  map[string]interface{}
	Version   string `json:"version,omitempty"`
	DeptNo    string `json:"deptNo,omitempty"`
	IsvUUID   string `json:"isvUUID,omitempty"`
	PageNo    int    `json:"pageNo,omitempty"`
	PageSize  int    `json:"pageSize,omitempty"`
	StartDate string `json:"startDate,omitempty"`
	EndDate   string `json:"endDate,omitempty"`
}

func (r *EclpOrderQueryOrderCustomsRequest) GetApiMethodName() string {
	return "jingdong.eclp.order.queryOrderCustoms"
}

func (r *EclpOrderQueryOrderCustomsRequest) GetApiParas() string {
	if len(r.apiParas) == 0 {
		return "{}"
	}
	data, _ := json.Marshal(r.apiParas)
	return string(data)
}

func (r *EclpOrderQueryOrderCustomsRequest) Check() {
	// 实现检查逻辑
}

func (r *EclpOrderQueryOrderCustomsRequest) PutOtherTextParam(key string, value interface{}) {
	r.apiParas[key] = value
}

func (r *EclpOrderQueryOrderCustomsRequest) SetVersion(version string) {
	r.Version = version
	r.apiParas["version"] = version
}

func (r *EclpOrderQueryOrderCustomsRequest) GetVersion() string {
	return r.Version
}

func (r *EclpOrderQueryOrderCustomsRequest) SetDeptNo(deptNo string) {
	r.DeptNo = deptNo
	r.apiParas["deptNo"] = deptNo
}

func (r *EclpOrderQueryOrderCustomsRequest) GetDeptNo() string {
	return r.DeptNo
}

func (r *EclpOrderQueryOrderCustomsRequest) SetIsvUUID(isvUUID string) {
	r.IsvUUID = isvUUID
	r.apiParas["isvUUID"] = isvUUID
}

func (r *EclpOrderQueryOrderCustomsRequest) GetIsvUUID() string {
	return r.IsvUUID
}

func (r *EclpOrderQueryOrderCustomsRequest) SetPageNo(pageNo int) {
	r.PageNo = pageNo
	r.apiParas["pageNo"] = pageNo
}

func (r *EclpOrderQueryOrderCustomsRequest) GetPageNo() int {
	return r.PageNo
}

func (r *EclpOrderQueryOrderCustomsRequest) SetPageSize(pageSize int) {
	r.PageSize = pageSize
	r.apiParas["pageSize"] = pageSize
}

func (r *EclpOrderQueryOrderCustomsRequest) GetPageSize() int {
	return r.PageSize
}

func (r *EclpOrderQueryOrderCustomsRequest) SetStartDate(startDate string) {
	r.StartDate = startDate
	r.apiParas["startDate"] = startDate
}

func (r *EclpOrderQueryOrderCustomsRequest) GetStartDate() string {
	return r.StartDate
}

func (r *EclpOrderQueryOrderCustomsRequest) SetEndDate(endDate string) {
	r.EndDate = endDate
	r.apiParas["endDate"] = endDate
}

func (r *EclpOrderQueryOrderCustomsRequest) GetEndDate() string {
	return r.EndDate
}
