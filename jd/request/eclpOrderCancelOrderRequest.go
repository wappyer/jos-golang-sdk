package request

/*
 * jingdong.eclp.order.cancelOrder
 * 销售出库单取消
 * 用于商家取消出库单，在 eclp 的订单取消，不入库商家不能调用
 */

type EclpOrderCancelOrderRequest struct {
	apiParas map[string]interface{}
	Version  string `json:"version,omitempty"`
	EclpSoNo string `json:"eclpSoNo,omitempty"`
}

func NewEclpOrderCancelOrderRequest() *EclpOrderCancelOrderRequest {
	return &EclpOrderCancelOrderRequest{
		apiParas: make(map[string]interface{}),
	}
}

func (r *EclpOrderCancelOrderRequest) GetApiMethodName() string {
	return "jingdong.eclp.order.cancelOrder"
}

func (r *EclpOrderCancelOrderRequest) GetApiParas() map[string]interface{} {
	if len(r.apiParas) == 0 {
		return map[string]interface{}{}
	}
	return r.apiParas
}

func (r *EclpOrderCancelOrderRequest) Check() {
	// 实现检查逻辑
}

func (r *EclpOrderCancelOrderRequest) PutOtherTextParam(key string, value interface{}) {
	r.apiParas[key] = value
}

func (r *EclpOrderCancelOrderRequest) SetVersion(version string) {
	r.Version = version
	r.apiParas["version"] = version
}

func (r *EclpOrderCancelOrderRequest) GetVersion() string {
	return r.Version
}

func (r *EclpOrderCancelOrderRequest) SetEclpSoNo(eclpSoNo string) {
	r.EclpSoNo = eclpSoNo
	r.apiParas["eclpSoNo"] = eclpSoNo
}

func (r *EclpOrderCancelOrderRequest) GetEclpSoNo() string {
	return r.EclpSoNo
}
