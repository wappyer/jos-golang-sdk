package service

import (
	"gitee.com/dimension-huimei/jos-golang-sdk/jd"
	"gitee.com/dimension-huimei/jos-golang-sdk/jd/request"
	"strings"
)

func MatchJdAddress(client *jd.Client, province, city, county string) (address []string, errResp request.ErrorResp, err error) {
	provinceId, cityId := 0, 0

	// 匹配省
	req := request.NewAreasProvinceGetRequest()
	err = client.Execute(req, client.AccessToken)
	if err != nil {
		return
	}
	provinceResp, errResp := req.GetResponse()
	if errResp.Code != "" {
		return
	}

	data := provinceResp.JingdongAreasProvinceGetResponce.BaseAreaServiceResponse.Data
	for _, v := range data {
		if strings.Index(province, v.AreaName) >= 0 {
			address = append(address, v.AreaName)
			provinceId = v.AreaID
		}
	}

	// 匹配市
	cityReq := request.NewAreasCityGetRequest()
	cityReq.SetParentId(provinceId)
	err = client.Execute(cityReq, client.AccessToken)
	if err != nil {
		return
	}
	cityResp, errResp := cityReq.GetResponse()
	if errResp.Code != "" {
		return
	}
	data = cityResp.JingdongAreasCityGetResponce.BaseAreaServiceResponse.Data
	for _, v := range data {
		if strings.Index(city, v.AreaName) >= 0 {
			address = append(address, v.AreaName)
			cityId = v.AreaID
		}
	}

	// 匹配区县
	countyReq := request.NewAreasCountyGetRequest()
	countyReq.SetParentId(cityId)
	err = client.Execute(countyReq, client.AccessToken)
	if err != nil {
		return
	}
	countyResp, errResp := countyReq.GetResponse()
	if errResp.Code != "" {
		return
	}
	data = countyResp.JingdongAreasCountyGetResponce.BaseAreaServiceResponse.Data
	for _, v := range data {
		if strings.Index(county, v.AreaName) >= 0 {
			address = append(address, v.AreaName)
		}
	}

	return
}
