package request

import "encoding/json"

/*
 * jingdong.eclp.goods.queryGoodsRecord
 * 跨境进口商家商品备案信息查询接口
 * 用于查询商品备案后信息
 * https://open.jd.com/v2/#/doc/api?apiCateId=138&apiId=2226&apiName=jingdong.eclp.goods.queryGoodsRecord
 */

type EclpGoodsQueryGoodsRecordRequest struct {
	apiParas map[string]interface{}

	Version    string
	DeptNo     string
	IsvGoodsNo string
	GoodsNo    string
	PageNo     int
	PageSize   int
	StartDate  string
	EndDate    string

	responseError ErrorResp
	responseData  interface{}
}

func NewEclpGoodsQueryGoodsRecordRequest() *EclpGoodsQueryGoodsRecordRequest {
	return &EclpGoodsQueryGoodsRecordRequest{
		apiParas: make(map[string]interface{}),
	}
}

func (r *EclpGoodsQueryGoodsRecordRequest) GetApiMethodName() string {
	return "jingdong.eclp.goods.queryGoodsRecord"
}

func (r *EclpGoodsQueryGoodsRecordRequest) GetApiParas() map[string]interface{} {
	if len(r.apiParas) == 0 {
		return map[string]interface{}{}
	}
	return r.apiParas
}

func (r *EclpGoodsQueryGoodsRecordRequest) Check() error {
	// Add validation logic here if needed
	return nil
}

func (r *EclpGoodsQueryGoodsRecordRequest) PutOtherTextParam(key string, value interface{}) {
	r.apiParas[key] = value
}

func (r *EclpGoodsQueryGoodsRecordRequest) SetVersion(version string) {
	r.Version = version
	r.apiParas["version"] = version
}

func (r *EclpGoodsQueryGoodsRecordRequest) GetVersion() string {
	return r.Version
}

func (r *EclpGoodsQueryGoodsRecordRequest) SetDeptNo(deptNo string) {
	r.DeptNo = deptNo
	r.apiParas["deptNo"] = deptNo
}

func (r *EclpGoodsQueryGoodsRecordRequest) GetDeptNo() string {
	return r.DeptNo
}

func (r *EclpGoodsQueryGoodsRecordRequest) SetIsvGoodsNo(isvGoodsNo string) {
	r.IsvGoodsNo = isvGoodsNo
	r.apiParas["isvGoodsNo"] = isvGoodsNo
}

func (r *EclpGoodsQueryGoodsRecordRequest) GetIsvGoodsNo() string {
	return r.IsvGoodsNo
}

func (r *EclpGoodsQueryGoodsRecordRequest) SetGoodsNo(goodsNo string) {
	r.GoodsNo = goodsNo
	r.apiParas["goodsNo"] = goodsNo
}

func (r *EclpGoodsQueryGoodsRecordRequest) GetGoodsNo() string {
	return r.GoodsNo
}

func (r *EclpGoodsQueryGoodsRecordRequest) SetPageNo(pageNo int) {
	r.PageNo = pageNo
	r.apiParas["pageNo"] = pageNo
}

func (r *EclpGoodsQueryGoodsRecordRequest) GetPageNo() int {
	return r.PageNo
}

func (r *EclpGoodsQueryGoodsRecordRequest) SetPageSize(pageSize int) {
	r.PageSize = pageSize
	r.apiParas["pageSize"] = pageSize
}

func (r *EclpGoodsQueryGoodsRecordRequest) GetPageSize() int {
	return r.PageSize
}

func (r *EclpGoodsQueryGoodsRecordRequest) SetStartDate(startDate string) {
	r.StartDate = startDate
	r.apiParas["startDate"] = startDate
}

func (r *EclpGoodsQueryGoodsRecordRequest) GetStartDate() string {
	return r.StartDate
}

func (r *EclpGoodsQueryGoodsRecordRequest) SetEndDate(endDate string) {
	r.EndDate = endDate
	r.apiParas["endDate"] = endDate
}

func (r *EclpGoodsQueryGoodsRecordRequest) GetEndDate() string {
	return r.EndDate
}

func (r *EclpGoodsQueryGoodsRecordRequest) SetResponseError(err ErrorResp) {
	r.responseError = err
}

func (r *EclpGoodsQueryGoodsRecordRequest) GetResponseError() ErrorResp {
	return r.responseError
}

func (r *EclpGoodsQueryGoodsRecordRequest) SetResponseData(data string) error {
	if err := json.Unmarshal([]byte(data), &r.responseData); err != nil {
		return err
	}
	return nil
}

func (r *EclpGoodsQueryGoodsRecordRequest) GetResponseData(responseData interface{}) error {
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

func (r *EclpGoodsQueryGoodsRecordRequest) GetResponse() (interface{}, ErrorResp) {
	return r.responseData, r.responseError
}
