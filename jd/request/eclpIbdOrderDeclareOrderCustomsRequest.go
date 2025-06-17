package request

import (
	"encoding/json"
	"gitee.com/dimension-huimei/jos-golang-sdk/jd/request/domain/eclpIbdOrderDeclareOrder"
)

/*
 * jingdong.eclp.ibd.order.declareOrderCustoms
 * 跨境进口入仓商家订单 申报并下仓生产接口
 * 用于跨境进口入仓商家的订单申报 并下仓生产
 */

type EclpIbdOrderDeclareOrderCustomsRequest struct {
	apiParas map[string]interface{}

	Version      string                                `json:"version,omitempty"`
	CustomsOrder eclpIbdOrderDeclareOrder.CustomsOrder `json:"customsOrder,omitempty"`
	GoodsList    []eclpIbdOrderDeclareOrder.Goods      `json:"goodsList,omitempty"`

	responseError ErrorResp
	responseData  interface{}
}

func NewEclpIbdOrderDeclareOrderCustomsRequest() *EclpIbdOrderDeclareOrderCustomsRequest {
	r := &EclpIbdOrderDeclareOrderCustomsRequest{
		apiParas: make(map[string]interface{}),
	}
	return r
}

func (r *EclpIbdOrderDeclareOrderCustomsRequest) GetApiMethodName() string {
	return "jingdong.eclp.ibd.order.declareOrderCustoms"
}

func (r *EclpIbdOrderDeclareOrderCustomsRequest) GetApiParas() map[string]interface{} {
	if len(r.apiParas) == 0 {
		return map[string]interface{}{}
	}
	return r.apiParas
}

func (r *EclpIbdOrderDeclareOrderCustomsRequest) Check() {
	// 实现检查逻辑
}

func (r *EclpIbdOrderDeclareOrderCustomsRequest) PutOtherTextParam(key string, value interface{}) {
	r.apiParas[key] = value
}

func (r *EclpIbdOrderDeclareOrderCustomsRequest) SetVersion(version string) {
	r.Version = version
	r.apiParas["version"] = version
}

func (r *EclpIbdOrderDeclareOrderCustomsRequest) GetVersion() string {
	return r.Version
}

func (r *EclpIbdOrderDeclareOrderCustomsRequest) SetCustomsOrder(customsOrder eclpIbdOrderDeclareOrder.CustomsOrder) {
	r.CustomsOrder = customsOrder
	r.apiParas["customsOrder"] = customsOrder
}

func (r *EclpIbdOrderDeclareOrderCustomsRequest) GetCustomsOrder() eclpIbdOrderDeclareOrder.CustomsOrder {
	return r.CustomsOrder
}

func (r *EclpIbdOrderDeclareOrderCustomsRequest) SetGoodsList(goodsList []eclpIbdOrderDeclareOrder.Goods) {
	r.GoodsList = goodsList
	r.apiParas["goodsList"] = goodsList
}

func (r *EclpIbdOrderDeclareOrderCustomsRequest) GetGoodsList() []eclpIbdOrderDeclareOrder.Goods {
	return r.GoodsList
}

func (r *EclpIbdOrderDeclareOrderCustomsRequest) SetResponseError(err ErrorResp) {
	r.responseError = err
}

func (r *EclpIbdOrderDeclareOrderCustomsRequest) GetResponseError() ErrorResp {
	return r.responseError
}

func (r *EclpIbdOrderDeclareOrderCustomsRequest) SetResponseData(data string) error {
	if err := json.Unmarshal([]byte(data), &r.responseData); err != nil {
		return err
	}
	return nil
}

func (r *EclpIbdOrderDeclareOrderCustomsRequest) GetResponseData(responseData interface{}) error {
	tmp, err := json.Marshal(responseData)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, responseData)
	if err != nil {
		return err
	}
	return nil
}

func (r *EclpIbdOrderDeclareOrderCustomsRequest) GetResponse() (interface{}, ErrorResp) {
	return r.responseData, r.responseError
}
