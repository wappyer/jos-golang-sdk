package request

import "encoding/json"

type AreasCountyGetRequest struct {
	apiParas map[string]interface{}

	version  string
	parentId int

	responseError ErrorResponse
	responseData  interface{}
}

func NewAreasCountyGetRequest() *AreasCountyGetRequest {
	return &AreasCountyGetRequest{
		apiParas: make(map[string]interface{}),
	}
}

func (r *AreasCountyGetRequest) GetApiMethodName() string {
	return "jingdong.areas.county.get"
}

func (r *AreasCountyGetRequest) GetApiParas() map[string]interface{} {
	if len(r.apiParas) == 0 {
		return map[string]interface{}{}
	}
	return r.apiParas
}

func (r *AreasCountyGetRequest) Check() {
	// 实现具体的检查逻辑
}

func (r *AreasCountyGetRequest) PutOtherTextParam(key string, value interface{}) {
	r.apiParas[key] = value
}

func (r *AreasCountyGetRequest) SetVersion(version string) {
	r.version = version
}

func (r *AreasCountyGetRequest) GetVersion() string {
	return r.version
}

func (r *AreasCountyGetRequest) SetParentId(parentId int) {
	r.parentId = parentId
	r.apiParas["parent_id"] = parentId
}

func (r *AreasCountyGetRequest) GetParentId() int {
	return r.parentId
}

func (r *AreasCountyGetRequest) SetResponseError(err ErrorResponse) {
	r.responseError = err
}

func (r *AreasCountyGetRequest) SetResponseData(data string) error {
	if err := json.Unmarshal([]byte(data), &r.responseData); err != nil {
		return err
	}
	return nil
}

func (r *AreasCountyGetRequest) GetResponse() (interface{}, ErrorResponse) {
	return r.responseData, r.responseError
}
