package request

import "encoding/json"

type AreasProvinceGetRequest struct {
	apiParas map[string]interface{}

	version string

	responseError ErrorResp
	responseData  AreasProvinceGetResponse
}

type AreasProvinceGetResponse struct {
	JingdongAreasProvinceGetResponce JingdongAreasProvinceGetResponce `json:"jingdong_areas_province_get_responce"`
}

type JingdongAreasProvinceGetResponce struct {
	BaseAreaServiceResponse BaseAreaServiceResponse `json:"baseAreaServiceResponse"`
	Code                    string                  `json:"code"`
	RequestID               string                  `json:"request_id"`
}

type BaseAreaServiceResponse struct {
	Data       []Area `json:"data"`
	ResultCode int    `json:"resultCode"`
}

type Area struct {
	AreaID   int    `json:"areaId"`
	AreaName string `json:"areaName"`
	Level    int    `json:"level"`
	ParentID int    `json:"parentId"`
	Status   int    `json:"status"`
}

func NewAreasProvinceGetRequest() *AreasProvinceGetRequest {
	return &AreasProvinceGetRequest{
		apiParas: make(map[string]interface{}),
	}
}

func (r *AreasProvinceGetRequest) GetApiMethodName() string {
	return "jingdong.areas.province.get"
}

func (r *AreasProvinceGetRequest) GetApiParas() map[string]interface{} {
	if len(r.apiParas) == 0 {
		return map[string]interface{}{}
	}
	return r.apiParas
}

func (r *AreasProvinceGetRequest) Check() {
	// 实现具体的检查逻辑
}

func (r *AreasProvinceGetRequest) PutOtherTextParam(key string, value interface{}) {
	r.apiParas[key] = value
}

func (r *AreasProvinceGetRequest) SetVersion(version string) {
	r.version = version
}

func (r *AreasProvinceGetRequest) GetVersion() string {
	return r.version
}

func (r *AreasProvinceGetRequest) SetResponseError(err ErrorResp) {
	r.responseError = err
}

func (r *AreasProvinceGetRequest) GetResponseError() ErrorResp {
	return r.responseError
}

func (r *AreasProvinceGetRequest) SetResponseData(data string) error {
	if err := json.Unmarshal([]byte(data), &r.responseData); err != nil {
		return err
	}
	return nil
}

func (r *AreasProvinceGetRequest) GetResponseData(responseData interface{}) error {
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

func (r *AreasProvinceGetRequest) GetResponse() (AreasProvinceGetResponse, ErrorResp) {
	return r.responseData, r.responseError
}
