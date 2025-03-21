package request

import (
	"encoding/json"
)

/*
 * jingdong.eclp.ibd.order.declareOrder
 * 跨境进口入仓商家订单 申报并下仓生产接口
 * 用于跨境进口入仓商家的订单申报 并下仓生产
 */

type EclpIbdOrderDeclareOrderRequest struct {
	apiParas map[string]interface{}

	Version      string       `json:"version,omitempty"`
	CustomsOrder CustomsOrder `json:"customsOrder,omitempty"`
	GoodsList    []Goods      `json:"goodsList,omitempty"`
}

type CustomsOrder struct {
	// 根据实际需求定义CustomsOrder的字段
	// 例如：
	OrderNo string `json:"orderNo,omitempty"`
	// 其他字段...
}

type Goods struct {
	// 根据实际需求定义Goods的字段
	// 例如：
	GoodsNo string `json:"goodsNo,omitempty"`
	// 其他字段...
}

func (r *EclpIbdOrderDeclareOrderRequest) GetApiMethodName() string {
	return "jingdong.eclp.ibd.order.declareOrder"
}

func (r *EclpIbdOrderDeclareOrderRequest) GetApiParas() string {
	if len(r.apiParas) == 0 {
		return "{}"
	}
	data, _ := json.Marshal(r.apiParas)
	return string(data)
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

func (r *EclpIbdOrderDeclareOrderRequest) SetCustomsOrder(customsOrder CustomsOrder) {
	r.CustomsOrder = customsOrder
	r.apiParas["customsOrder"] = customsOrder
}

func (r *EclpIbdOrderDeclareOrderRequest) GetCustomsOrder() CustomsOrder {
	return r.CustomsOrder
}

func (r *EclpIbdOrderDeclareOrderRequest) SetGoodsList(goodsList []Goods) {
	r.GoodsList = goodsList
	r.apiParas["goodsList"] = goodsList
}

func (r *EclpIbdOrderDeclareOrderRequest) GetGoodsList() []Goods {
	return r.GoodsList
}
