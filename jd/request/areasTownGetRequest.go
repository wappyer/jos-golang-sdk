package request

import "encoding/json"

type AreasTownGetRequest struct {
	apiParas map[string]interface{}

	version  string
	parentId int

	responseError ErrorResponse
	responseData  interface{}
}

func NewAreasTownGetRequest() *AreasTownGetRequest {
	return &AreasTownGetRequest{
		apiParas: make(map[string]interface{}),
	}
}

func (r *AreasTownGetRequest) GetApiMethodName() string {
	return "jingdong.areas.town.get"
}

func (r *AreasTownGetRequest) GetApiParas() map[string]interface{} {
	if len(r.apiParas) == 0 {
		return map[string]interface{}{}
	}
	return r.apiParas
}

func (r *AreasTownGetRequest) Check() {
	// 实现具体的检查逻辑
}

func (r *AreasTownGetRequest) PutOtherTextParam(key string, value interface{}) {
	r.apiParas[key] = value
}

func (r *AreasTownGetRequest) SetVersion(version string) {
	r.version = version
}

func (r *AreasTownGetRequest) GetVersion() string {
	return r.version
}

func (r *AreasTownGetRequest) SetParentId(parentId int) {
	r.parentId = parentId
	r.apiParas["parent_id"] = parentId
}

func (r *AreasTownGetRequest) GetParentId() int {
	return r.parentId
}

func (r *AreasTownGetRequest) SetResponseError(err ErrorResponse) {
	r.responseError = err
}

func (r *AreasTownGetRequest) SetResponseData(data string) error {
	if err := json.Unmarshal([]byte(data), &r.responseData); err != nil {
		return err
	}
	return nil
}

func (r *AreasTownGetRequest) GetResponse() (interface{}, ErrorResponse) {
	return r.responseData, r.responseError
}
