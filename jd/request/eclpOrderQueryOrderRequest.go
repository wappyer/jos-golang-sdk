package request

/*
 * jingdong.eclp.order.queryOrder
 * 订单详情查询接口
 * 用于商家查询订单详情(订单当前 状态，订单信息，订单状态流水)
 */

type EclpOrderQueryOrderRequest struct {
	apiParas map[string]interface{}
	Version  string `json:"version,omitempty"`
	EclpSoNo string `json:"eclpSoNo,omitempty"`
}

func NewEclpOrderQueryOrderRequest() *EclpOrderQueryOrderRequest {
	return &EclpOrderQueryOrderRequest{
		apiParas: make(map[string]interface{}),
	}
}

func (r *EclpOrderQueryOrderRequest) GetApiMethodName() string {
	return "jingdong.eclp.order.queryOrder"
}

func (r *EclpOrderQueryOrderRequest) GetApiParas() map[string]interface{} {
	if len(r.apiParas) == 0 {
		return map[string]interface{}{}
	}
	return r.apiParas
}

func (r *EclpOrderQueryOrderRequest) Check() {
	// 实现检查逻辑
}

func (r *EclpOrderQueryOrderRequest) PutOtherTextParam(key string, value interface{}) {
	r.apiParas[key] = value
}

func (r *EclpOrderQueryOrderRequest) SetVersion(version string) {
	r.Version = version
	r.apiParas["version"] = version
}

func (r *EclpOrderQueryOrderRequest) GetVersion() string {
	return r.Version
}

func (r *EclpOrderQueryOrderRequest) SetEclpSoNo(eclpSoNo string) {
	r.EclpSoNo = eclpSoNo
	r.apiParas["eclpSoNo"] = eclpSoNo
}

func (r *EclpOrderQueryOrderRequest) GetEclpSoNo() string {
	return r.EclpSoNo
}
