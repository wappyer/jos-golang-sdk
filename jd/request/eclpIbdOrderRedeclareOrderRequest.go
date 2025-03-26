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
	apiParas     map[string]interface{}
	Version      string                                `json:"version,omitempty"`
	CustomsOrder eclpIbdOrderDeclareOrder.CustomsOrder `json:"customsOrder,omitempty"`
	GoodsList    []eclpIbdOrderDeclareOrder.Goods      `json:"goodsList,omitempty"`
}

func (r *EclpIbdOrderRedeclareOrderRequest) GetApiMethodName() string {
	return "jingdong.eclp.ibd.order.redeclareOrder"
}

func (r *EclpIbdOrderRedeclareOrderRequest) GetApiParas() string {
	if len(r.apiParas) == 0 {
		return "{}"
	}
	data, _ := json.Marshal(r.apiParas)
	return string(data)
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
	r.apiParas["customsOrder"] = customsOrder
}

func (r *EclpIbdOrderRedeclareOrderRequest) GetCustomsOrder() interface{} {
	return r.CustomsOrder
}

func (r *EclpIbdOrderRedeclareOrderRequest) SetGoodsList(goodsList []eclpIbdOrderDeclareOrder.Goods) {
	r.GoodsList = goodsList
	r.apiParas["goodsList"] = goodsList
}

func (r *EclpIbdOrderRedeclareOrderRequest) GetGoodsList() []eclpIbdOrderDeclareOrder.Goods {
	return r.GoodsList
}
