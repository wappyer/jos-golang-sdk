package service

import (
	"encoding/json"
	"errors"
	"gitee.com/dimension-huimei/jos-golang-sdk/jd"
	"gitee.com/dimension-huimei/jos-golang-sdk/jd/request"
)

func MatchJdProvince(client *jd.Client, search string) (province string, errResp jd.ErrorResp, err error) {
	req := request.NewAreasProvinceGetRequest()
	dataResp, errResp, err := client.Execute(req, client.AccessToken)
	if err != nil || errResp.Code != "" {
		return
	}

	resp := request.AreasProvinceGetResponse{}
	if e := json.Unmarshal([]byte(dataResp), &resp); e != nil {
		err = errors.New("格式化消息返回失败")
		return
	}
	data := resp.JingdongAreasProvinceGetResponce.BaseAreaServiceResponse.Data
	for _, v := range data {
		if v.AreaName == search {
			province = v.AreaName
			return
		}
	}

	return
}
