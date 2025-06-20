package request

import (
	"encoding/json"
	"gitee.com/dimension-huimei/jos-golang-sdk/jd/request/domain/eclpIbdOrderDeclareOrder"
)

/*
 * jingdong.eclp.ibd.order.redeclareOrder
 * 跨境进口商家修改申报单据并重新向海关申报单据信息接口
 * 用于跨境进口商家修改申报单据并重新向海关申报单据信息(适用于入仓和不入仓商家)
 */

type EclpIbdOrderRedeclareOrderRequest struct {
	apiParas map[string]interface{}

	Version      string                                `json:"version,omitempty"`
	CustomsOrder eclpIbdOrderDeclareOrder.CustomsOrder `json:"customsOrder,omitempty"`
	GoodsList    []eclpIbdOrderDeclareOrder.Goods      `json:"goodsList,omitempty"`

	responseError ErrorResp
	responseData  EclpIbdOrderRedeclareOrderResponse
}

type EclpIbdOrderRedeclareOrderResponse struct {
	JingdongEclpIbdOrderRedeclareOrderResponce JingdongEclpIbdOrderRedeclareOrderResponce `json:"jingdong_eclp_ibd_order_redeclareOrder_responce"`
}

type JingdongEclpIbdOrderRedeclareOrderResponce struct {
	Code      string `json:"code"`
	RequestID string `json:"request_id"`
	Result    Result `json:"result"`
}

type Result struct {
	ResultCode string `json:"resultCode"`
	Message    string `json:"message"`
}

func NewEclpIbdOrderRedeclareOrderRequest() *EclpIbdOrderRedeclareOrderRequest {
	r := &EclpIbdOrderRedeclareOrderRequest{
		apiParas: make(map[string]interface{}),
	}
	return r
}

func (r *EclpIbdOrderRedeclareOrderRequest) GetApiMethodName() string {
	return "jingdong.eclp.ibd.order.redeclareOrder"
}

func (r *EclpIbdOrderRedeclareOrderRequest) GetApiParas() map[string]interface{} {
	if len(r.apiParas) == 0 {
		return map[string]interface{}{}
	}
	return r.apiParas
}

func (r *EclpIbdOrderRedeclareOrderRequest) Check() {
	// 实现检查逻辑
}

func (r *EclpIbdOrderRedeclareOrderRequest) PutOtherTextParam(key string, value interface{}) {
	r.apiParas[key] = value
}

func (r *EclpIbdOrderRedeclareOrderRequest) SetVersion(version string) {
	r.Version = version
	r.apiParas["version"] = version
}

func (r *EclpIbdOrderRedeclareOrderRequest) GetVersion() string {
	return r.Version
}

func (r *EclpIbdOrderRedeclareOrderRequest) SetCustomsOrder(customsOrder eclpIbdOrderDeclareOrder.CustomsOrder) {
	r.CustomsOrder = customsOrder
	r.apiParas["customsOrder"] = customsOrder.GetApiParas()
}

func (r *EclpIbdOrderRedeclareOrderRequest) GetCustomsOrder() interface{} {
	return r.CustomsOrder
}

func (r *EclpIbdOrderRedeclareOrderRequest) SetGoodsList(goodsList []eclpIbdOrderDeclareOrder.Goods) {
	r.GoodsList = goodsList

	apiParamsGoodsList := make([]map[string]interface{}, len(goodsList))
	for i, goods := range goodsList {
		apiParamsGoodsList[i] = goods.GetApiParas()
	}
	r.apiParas["goodsList"] = apiParamsGoodsList
}

func (r *EclpIbdOrderRedeclareOrderRequest) GetGoodsList() []eclpIbdOrderDeclareOrder.Goods {
	return r.GoodsList
}

func (r *EclpIbdOrderRedeclareOrderRequest) SetResponseError(err ErrorResp) {
	r.responseError = err
}

func (r *EclpIbdOrderRedeclareOrderRequest) GetResponseError() ErrorResp {
	return r.responseError
}

func (r *EclpIbdOrderRedeclareOrderRequest) SetResponseData(data string) error {
	if err := json.Unmarshal([]byte(data), &r.responseData); err != nil {
		return err
	}
	return nil
}

func (r *EclpIbdOrderRedeclareOrderRequest) GetResponseData(responseData interface{}) error {
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

func (r *EclpIbdOrderRedeclareOrderRequest) GetResponse() (EclpIbdOrderRedeclareOrderResponse, ErrorResp) {
	return r.responseData, r.responseError
}
