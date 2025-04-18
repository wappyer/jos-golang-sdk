package request

import (
	"encoding/json"
	"gitee.com/dimension-huimei/jos-golang-sdk/jd/request/domain/eclpIbdOrderDeclareOrder"
)

/*
 * jingdong.eclp.ibd.order.declareOrder
 * 跨境进口入仓商家订单 申报并下仓生产接口
 * 用于跨境进口入仓商家的订单申报 并下仓生产
 */

type EclpIbdOrderDeclareOrderRequest struct {
	apiParas map[string]interface{}

	Version      string                                `json:"version,omitempty"`
	CustomsOrder eclpIbdOrderDeclareOrder.CustomsOrder `json:"customsOrder,omitempty"`
	GoodsList    []eclpIbdOrderDeclareOrder.Goods      `json:"goodsList,omitempty"`

	responseError ErrorResp
	responseData  interface{}
}

func NewEclpIbdOrderDeclareOrderRequest() *EclpIbdOrderDeclareOrderRequest {
	r := &EclpIbdOrderDeclareOrderRequest{
		apiParas: make(map[string]interface{}),
	}
	return r
}

func (r *EclpIbdOrderDeclareOrderRequest) GetApiMethodName() string {
	return "jingdong.eclp.ibd.order.declareOrder"
}

func (r *EclpIbdOrderDeclareOrderRequest) GetApiParas() map[string]interface{} {
	if len(r.apiParas) == 0 {
		return map[string]interface{}{}
	}
	return r.apiParas
}

func (r *EclpIbdOrderDeclareOrderRequest) Check() {
	// 实现检查逻辑
}

func (r *EclpIbdOrderDeclareOrderRequest) PutOtherTextParam(key string, value interface{}) {
	r.apiParas[key] = value
}

func (r *EclpIbdOrderDeclareOrderRequest) SetVersion(version string) {
	r.Version = version
	r.apiParas["version"] = version
}

func (r *EclpIbdOrderDeclareOrderRequest) GetVersion() string {
	return r.Version
}

func (r *EclpIbdOrderDeclareOrderRequest) SetCustomsOrder(customsOrder eclpIbdOrderDeclareOrder.CustomsOrder) {
	r.CustomsOrder = customsOrder
	r.apiParas["customsOrder"] = customsOrder
}

func (r *EclpIbdOrderDeclareOrderRequest) GetCustomsOrder() eclpIbdOrderDeclareOrder.CustomsOrder {
	return r.CustomsOrder
}

func (r *EclpIbdOrderDeclareOrderRequest) SetGoodsList(goodsList []eclpIbdOrderDeclareOrder.Goods) {
	r.GoodsList = goodsList
	r.apiParas["goodsList"] = goodsList
}

func (r *EclpIbdOrderDeclareOrderRequest) GetGoodsList() []eclpIbdOrderDeclareOrder.Goods {
	return r.GoodsList
}

func (r *EclpIbdOrderDeclareOrderRequest) SetResponseError(err ErrorResponse) {
	r.responseError = err.ErrorResp
}

func (r *EclpIbdOrderDeclareOrderRequest) SetResponseData(data string) error {
	if err := json.Unmarshal([]byte(data), &r.responseData); err != nil {
		return err
	}
	return nil
}

func (r *EclpIbdOrderDeclareOrderRequest) GetResponse() (interface{}, ErrorResp) {
	return r.responseData, r.responseError
}
