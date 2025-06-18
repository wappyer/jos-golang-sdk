package request

import "encoding/json"

type UnionOpenCategoryGoodsGetRequest struct {
	apiParas map[string]interface{}

	version string
	req     UnionOpenCategoryGoodsGetRequestReq

	responseError ErrorResp
	responseData  interface{}
}

type UnionOpenCategoryGoodsGetRequestReq struct {
	ParentId int `json:"parentId"`
	Grade    int `json:"grade"`
}

func NewUnionOpenCategoryGoodsGetRequest() *UnionOpenCategoryGoodsGetRequest {
	return &UnionOpenCategoryGoodsGetRequest{
		apiParas: make(map[string]interface{}),
	}
}

func (r *UnionOpenCategoryGoodsGetRequest) GetApiMethodName() string {
	return "jd.union.open.category.goods.get"
}

func (r *UnionOpenCategoryGoodsGetRequest) GetApiParas() map[string]interface{} {
	if len(r.apiParas) == 0 {
		return map[string]interface{}{}
	}
	return r.apiParas
}

func (r *UnionOpenCategoryGoodsGetRequest) Check() {
	// 实现具体的检查逻辑
}

func (r *UnionOpenCategoryGoodsGetRequest) PutOtherTextParam(key string, value interface{}) {
	r.apiParas[key] = value
}

func (r *UnionOpenCategoryGoodsGetRequest) SetVersion(version string) {
	r.version = version
}

func (r *UnionOpenCategoryGoodsGetRequest) GetVersion() string {
	return r.version
}

func (r *UnionOpenCategoryGoodsGetRequest) SetReq(req UnionOpenCategoryGoodsGetRequestReq) {
	r.req = req
}

func (r *UnionOpenCategoryGoodsGetRequest) GetReq() UnionOpenCategoryGoodsGetRequestReq {
	return r.req
}

func (r *UnionOpenCategoryGoodsGetRequest) SetResponseError(err ErrorResp) {
	r.responseError = err
}

func (r *UnionOpenCategoryGoodsGetRequest) GetResponseError() ErrorResp {
	return r.responseError
}

func (r *UnionOpenCategoryGoodsGetRequest) SetResponseData(data string) error {
	if err := json.Unmarshal([]byte(data), &r.responseData); err != nil {
		return err
	}
	return nil
}

func (r *UnionOpenCategoryGoodsGetRequest) GetResponseData(responseData interface{}) error {
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

func (r *UnionOpenCategoryGoodsGetRequest) GetResponse() (interface{}, ErrorResp) {
	return r.responseData, r.responseError
}
