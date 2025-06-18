package request

import "encoding/json"

/*
 * jingdong.eclp.order.queryOrderCustoms
 * 订单申报结果查询接口
 * 用于商家查询订单申报状态
 */

type EclpOrderQueryOrderCustomsRequest struct {
	apiParas map[string]interface{}

	DeptNo    string `json:"deptNo"`
	IsvUUID   string `json:"isvUUID"`
	PageNo    int    `json:"pageNo"`
	PageSize  int    `json:"pageSize"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`

	responseError ErrorResp
	responseData  EclpOrderQueryOrderCustomsResponse
}

type EclpOrderQueryOrderCustomsResponse struct {
	JingdongEclpOrderQueryOrderCustomsResponce JingdongEclpOrderQueryOrderCustomsResponce `json:"jingdong_eclp_order_queryOrderCustoms_responce"`
}

type JingdongEclpOrderQueryOrderCustomsResponce struct {
	Code                    string                    `json:"code"`
	QueryordercustomsResult []QueryordercustomsResult `json:"queryordercustoms_result"`
	RequestID               string                    `json:"request_id"`
}

type QueryordercustomsResult struct {
	CustomID      string `json:"customId"`
	DeptNo        string `json:"deptNo"`
	ErrorCode     string `json:"errorCode"`
	GoodsCheck    int    `json:"goodsCheck"`
	IsvUUID       string `json:"isvUUID"`
	LogisticsCode string `json:"logisticsCode"`
	Message       string `json:"message"`
	Pattern       string `json:"pattern"`
	ProcessStatus string `json:"processStatus"`
	Result        int    `json:"result"`
	SpSoNo        string `json:"spSoNo"`
	StoreID       string `json:"storeId"`
	Time          int64  `json:"time"`
}

func NewEclpOrderQueryOrderCustomsRequest() *EclpOrderQueryOrderCustomsRequest {
	r := EclpOrderQueryOrderCustomsRequest{
		apiParas: make(map[string]interface{}),
	}
	return &r
}

func (r *EclpOrderQueryOrderCustomsRequest) GetApiMethodName() string {
	return "jingdong.eclp.order.queryOrderCustoms"
}

func (r *EclpOrderQueryOrderCustomsRequest) GetApiParas() map[string]interface{} {
	if len(r.apiParas) == 0 {
		return map[string]interface{}{}
	}
	return r.apiParas
}

func (r *EclpOrderQueryOrderCustomsRequest) Check() {
	// 实现检查逻辑
}

func (r *EclpOrderQueryOrderCustomsRequest) PutOtherTextParam(key string, value interface{}) {
	r.apiParas[key] = value
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

func (r *EclpOrderQueryOrderCustomsRequest) SetResponseError(err ErrorResp) {
	r.responseError = err
}

func (r *EclpOrderQueryOrderCustomsRequest) GetResponseError() ErrorResp {
	return r.responseError
}

func (r *EclpOrderQueryOrderCustomsRequest) SetResponseData(data string) error {
	if err := json.Unmarshal([]byte(data), &r.responseData); err != nil {
		return err
	}
	return nil
}

func (r *EclpOrderQueryOrderCustomsRequest) GetResponseData(responseData interface{}) error {
	tmp, err := json.Marshal(r.responseData)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, responseData)
	if err != nil {
		return err
	}
	return nil
}

func (r *EclpOrderQueryOrderCustomsRequest) GetResponse() (EclpOrderQueryOrderCustomsResponse, ErrorResp) {
	return r.responseData, r.responseError
}
