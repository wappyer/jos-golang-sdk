package jd

import (
	"fmt"
	"gitee.com/dimension-huimei/jos-golang-sdk/jd/request"
	"testing"
)

var client *Client
var deptNo string

func init() {
	client = NewClient("", "")
	client.AccessToken = ""
	deptNo = ""
}

func TestClientExec_EclpGoodsQueryGoodsRecordRequest(t *testing.T) {
	PrintRes(func() (interface{}, error) {
		req := request.NewEclpGoodsQueryGoodsRecordRequest()
		req.SetDeptNo(deptNo)
		req.SetStartDate("2023-01-01 00:00:00")
		req.SetEndDate("2025-03-20 00:00:00")
		req.SetPageNo(1)
		req.SetPageSize(10)
		return client.Execute(req, client.AccessToken)
	})
}

func TestClientExec_EclpGoodsQueryGoodsInfoRequest(t *testing.T) {
	PrintRes(func() (interface{}, error) {
		req := request.NewEclpGoodsQueryGoodsInfoRequest()
		req.SetGoodsNos([]string{"EMG4418387067782", "EMG4418461043468", "EMG4418457048961", "EMG4418640984775", "EMG4418624396784"})
		req.SetDeptNo(deptNo)
		req.SetQueryType("2")
		req.SetPageNo(1)
		req.SetPageSize(100)
		return client.Execute(req, client.AccessToken)
	})
}

func PrintRes(f func() (interface{}, error)) {
	resp, err := f()
	if err != nil {
		fmt.Printf("Error:%#v\n", err)
	} else {
		fmt.Printf("Response:%#v\n", resp)
	}
}
