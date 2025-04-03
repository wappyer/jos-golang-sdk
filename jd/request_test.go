package jd

import (
	"fmt"
	"gitee.com/dimension-huimei/jos-golang-sdk/jd/request"
	"testing"
)

var client *Client
var deptNo string

func init() {
	client = NewClient("https://api-dev.jd.com/routerjson", "BB5039D6E7D11C7A481E092BF6350CB3", "")
	client.AccessToken = "47fd6890bfe8483ba1377119cbc9d85cy4ot"
	deptNo = "EBU4418055058519"
}

func TestClientExec_EclpGoodsQueryGoodsRecordRequest(t *testing.T) {
	req := request.NewEclpGoodsQueryGoodsRecordRequest()
	req.SetDeptNo(deptNo)
	req.SetStartDate("2023-01-01 00:00:00")
	req.SetEndDate("2025-03-20 00:00:00")
	req.SetPageNo(1)
	req.SetPageSize(10)
	err := client.Execute(req, client.AccessToken)
	if err != nil {
		fmt.Printf("Error:%#v\n", err)
	} else {
		respData, respErr := req.GetResponse()
		fmt.Printf("ResponseData:%#v\n", respData)
		fmt.Printf("ResponseError:%#v\n", respErr)
	}
}

func TestClientExec_EclpGoodsQueryGoodsInfoRequest(t *testing.T) {
	req := request.NewEclpGoodsQueryGoodsInfoRequest()
	req.SetGoodsNos([]string{"EMG4418387067782", "EMG4418461043468", "EMG4418457048961", "EMG4418640984775", "EMG4418624396784"})
	req.SetDeptNo(deptNo)
	req.SetQueryType("2")
	req.SetPageNo(1)
	req.SetPageSize(100)
	err := client.Execute(req, client.AccessToken)
	if err != nil {
		fmt.Printf("Error:%#v\n", err)
	} else {
		respData, respErr := req.GetResponse()
		fmt.Printf("ResponseData:%#v\n", respData)
		fmt.Printf("ResponseError:%#v\n", respErr)
	}
}

func TestClientExec_AreasProvinceGetRequest(t *testing.T) {
	req := request.NewAreasProvinceGetRequest()
	err := client.Execute(req, client.AccessToken)
	if err != nil {
		fmt.Printf("Error:%#v\n", err)
	} else {
		respData, respErr := req.GetResponse()
		fmt.Printf("ResponseData:%#v\n", respData)
		fmt.Printf("ResponseError:%#v\n", respErr)
	}
}

func TestClientExec_AreasCityGetRequest(t *testing.T) {
	req := request.NewAreasCityGetRequest()
	req.SetParentId(21)
	err := client.Execute(req, client.AccessToken)
	if err != nil {
		fmt.Printf("Error:%#v\n", err)
	} else {
		respData, respErr := req.GetResponse()
		fmt.Printf("ResponseData:%#v\n", respData)
		fmt.Printf("ResponseError:%#v\n", respErr)
	}
}

func TestClientExec_AreasCountyGetRequest(t *testing.T) {
	req := request.NewAreasCountyGetRequest()
	req.SetParentId(1827)
	err := client.Execute(req, client.AccessToken)
	if err != nil {
		fmt.Printf("Error:%#v\n", err)
	} else {
		respData, respErr := req.GetResponse()
		fmt.Printf("ResponseData:%#v\n", respData)
		fmt.Printf("ResponseError:%#v\n", respErr)
	}
}
