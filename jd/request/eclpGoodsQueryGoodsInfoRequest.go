package request

import (
	"encoding/json"
	"strings"
)

/*
 * jingdong.eclp.goods.queryGoodsInfo
 * ECLP 商品主数据查询接口
 * 用于商家以事业部维度查询商品信
 * https://open.jd.com/v2/#/doc/api?apiCateId=138&apiId=945&apiName=jingdong.eclp.goods.queryGoodsInfo
 */

type EclpGoodsQueryGoodsInfoRequest struct {
	apiParas map[string]interface{}

	Version     string   `json:"version,omitempty"`
	DeptNo      string   `json:"deptNo,omitempty"`
	IsvGoodsNos []string `json:"isvGoodsNos,omitempty"`
	GoodsNos    []string `json:"goodsNos,omitempty"`
	QueryType   string   `json:"queryType,omitempty"`
	Barcodes    []string `json:"barcodes,omitempty"`
	PageNo      int      `json:"pageNo,omitempty"`
	PageSize    int      `json:"pageSize,omitempty"`

	responseError ErrorResp
	responseData  EclpGoodsQueryGoodsInfoResponse
}

type EclpGoodsQueryGoodsInfoResponse struct {
	JingdongEclpGoodsQueryGoodsInfoResponce JingdongEclpGoodsQueryGoodsInfoResponce `json:"jingdong_eclp_goods_queryGoodsInfo_responce"`
}

type JingdongEclpGoodsQueryGoodsInfoResponce struct {
	Code          string      `json:"code"`
	GoodsInfoList []GoodsInfo `json:"goodsInfoList"`
	RequestID     string      `json:"request_id"`
}

type GoodsInfo struct {
	Barcodes string `json:"barcodes"`
	DeptNo   string `json:"deptNo"`
	GoodsNo  string `json:"goodsNo"`
}

func NewEclpGoodsQueryGoodsInfoRequest() *EclpGoodsQueryGoodsInfoRequest {
	return &EclpGoodsQueryGoodsInfoRequest{
		apiParas: make(map[string]interface{}),
	}
}

func (r *EclpGoodsQueryGoodsInfoRequest) GetApiMethodName() string {
	return "jingdong.eclp.goods.queryGoodsInfo"
}

func (r *EclpGoodsQueryGoodsInfoRequest) GetApiParas() map[string]interface{} {
	if len(r.apiParas) == 0 {
		return map[string]interface{}{}
	}
	return r.apiParas
}

func (r *EclpGoodsQueryGoodsInfoRequest) Check() {
	// 实现检查逻辑
}

func (r *EclpGoodsQueryGoodsInfoRequest) PutOtherTextParam(key string, value interface{}) {
	r.apiParas[key] = value
}

func (r *EclpGoodsQueryGoodsInfoRequest) SetVersion(version string) {
	r.Version = version
	r.apiParas["version"] = version
}

func (r *EclpGoodsQueryGoodsInfoRequest) GetVersion() string {
	return r.Version
}

func (r *EclpGoodsQueryGoodsInfoRequest) SetDeptNo(deptNo string) {
	r.DeptNo = deptNo
	r.apiParas["deptNo"] = deptNo
}

func (r *EclpGoodsQueryGoodsInfoRequest) GetDeptNo() string {
	return r.DeptNo
}

func (r *EclpGoodsQueryGoodsInfoRequest) SetIsvGoodsNos(isvGoodsNos []string) {
	r.IsvGoodsNos = isvGoodsNos
	r.apiParas["isvGoodsNos"] = strings.Join(isvGoodsNos, ",")
}

func (r *EclpGoodsQueryGoodsInfoRequest) GetIsvGoodsNos() []string {
	return r.IsvGoodsNos
}

func (r *EclpGoodsQueryGoodsInfoRequest) SetGoodsNos(goodsNos []string) {
	r.GoodsNos = goodsNos
	r.apiParas["goodsNos"] = strings.Join(goodsNos, ",")
}

func (r *EclpGoodsQueryGoodsInfoRequest) GetGoodsNos() []string {
	return r.GoodsNos
}

func (r *EclpGoodsQueryGoodsInfoRequest) SetQueryType(queryType string) {
	r.QueryType = queryType
	r.apiParas["queryType"] = queryType
}

func (r *EclpGoodsQueryGoodsInfoRequest) GetQueryType() string {
	return r.QueryType
}

func (r *EclpGoodsQueryGoodsInfoRequest) SetBarcodes(barcodes []string) {
	r.Barcodes = barcodes
	r.apiParas["barcodes"] = strings.Join(barcodes, ",")
}

func (r *EclpGoodsQueryGoodsInfoRequest) GetBarcodes() []string {
	return r.Barcodes
}

func (r *EclpGoodsQueryGoodsInfoRequest) SetPageNo(pageNo int) {
	r.PageNo = pageNo
	r.apiParas["pageNo"] = pageNo
}

func (r *EclpGoodsQueryGoodsInfoRequest) GetPageNo() int {
	return r.PageNo
}

func (r *EclpGoodsQueryGoodsInfoRequest) SetPageSize(pageSize int) {
	r.PageSize = pageSize
	r.apiParas["pageSize"] = pageSize
}

func (r *EclpGoodsQueryGoodsInfoRequest) GetPageSize() int {
	return r.PageSize
}

func (r *EclpGoodsQueryGoodsInfoRequest) SetResponseError(err ErrorResp) {
	r.responseError = err
}

func (r *EclpGoodsQueryGoodsInfoRequest) GetResponseError() ErrorResp {
	return r.responseError
}

func (r *EclpGoodsQueryGoodsInfoRequest) SetResponseData(data string) error {
	if err := json.Unmarshal([]byte(data), &r.responseData); err != nil {
		return err
	}
	return nil
}

func (r *EclpGoodsQueryGoodsInfoRequest) GetResponseData(responseData interface{}) error {
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

func (r *EclpGoodsQueryGoodsInfoRequest) GetResponse() (EclpGoodsQueryGoodsInfoResponse, ErrorResp) {
	return r.responseData, r.responseError
}
