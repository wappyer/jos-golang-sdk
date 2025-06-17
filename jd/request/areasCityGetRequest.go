package request

import "encoding/json"

type AreasCityGetRequest struct {
	apiParas map[string]interface{}

	version  string
	parentId int

	responseError ErrorResp
	responseData  AreasCityGetResponse
}

type AreasCityGetResponse struct {
	JingdongAreasCityGetResponce JingdongAreasCityGetResponce `json:"jingdong_areas_city_get_responce"`
}

type JingdongAreasCityGetResponce struct {
	BaseAreaServiceResponse BaseAreaServiceResponse `json:"baseAreaServiceResponse"`
	Code                    string                  `json:"code"`
	RequestID               string                  `json:"request_id"`
}

func NewAreasCityGetRequest() *AreasCityGetRequest {
	return &AreasCityGetRequest{
		apiParas: make(map[string]interface{}),
	}
}

func (r *AreasCityGetRequest) GetApiMethodName() string {
	return "jingdong.areas.city.get"
}

func (r *AreasCityGetRequest) GetApiParas() map[string]interface{} {
	if len(r.apiParas) == 0 {
		return map[string]interface{}{}
	}
	return r.apiParas
}

func (r *AreasCityGetRequest) Check() {
	// 实现具体的检查逻辑
}

func (r *AreasCityGetRequest) PutOtherTextParam(key string, value interface{}) {
	r.apiParas[key] = value
}

func (r *AreasCityGetRequest) SetVersion(version string) {
	r.version = version
}

func (r *AreasCityGetRequest) GetVersion() string {
	return r.version
}

func (r *AreasCityGetRequest) SetParentId(parentId int) {
	r.parentId = parentId
	r.apiParas["parent_id"] = parentId
}

func (r *AreasCityGetRequest) GetParentId() int {
	return r.parentId
}

func (r *AreasCityGetRequest) SetResponseError(err ErrorResp) {
	r.responseError = err
}

func (r *AreasCityGetRequest) GetResponseError() ErrorResp {
	return r.responseError
}

func (r *AreasCityGetRequest) SetResponseData(data string) error {
	if err := json.Unmarshal([]byte(data), &r.responseData); err != nil {
		return err
	}
	return nil
}

func (r *AreasCityGetRequest) GetResponseData(responseData interface{}) error {
	tmp, err := json.Marshal(responseData)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, responseData)
	if err != nil {
		return err
	}
	return nil
}

func (r *AreasCityGetRequest) GetResponse() (AreasCityGetResponse, ErrorResp) {
	return r.responseData, r.responseError
}
