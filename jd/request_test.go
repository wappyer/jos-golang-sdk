package jd

import (
	"fmt"
	"testing"
)

var client *Client
var deptNo string

func init() {
	client = NewClient("BB5039D6E7D11C7A481E092BF6350CB3", "")
	client.AccessToken = ""
	deptNo = ""
}

func TestClientExec_EclpGoodsQueryGoodsRecordRequest(t *testing.T) {
	PrintRes(func() (interface{}, error) {
		return client.Exec(map[string]interface{}{
			"method":    "jingdong.eclp.goods.queryGoodsRecord",
			"deptNo":    deptNo,
			"goodsNo":   "",
			"pageNo":    1,
			"pageSize":  10,
			"startDate": "2023-01-01 00:00:00",
			"endDate":   "2025-03-20 00:00:00",
			"session":   client.AccessToken,
		})
	})
}

func TestClientExec_EclpGoodsQueryGoodsInfoRequest(t *testing.T) {
	PrintRes(func() (interface{}, error) {
		return client.Exec(map[string]interface{}{
			"method": "jingdong.eclp.goods.queryGoodsInfo",
			"deptNo": deptNo,
			//"isvGoodsNos": []string{"ESG4418851920152", "ESG4418851919664"},
			"goodsNos":  []string{"EMG4418387067782", "EMG4418461043468", "EMG4418457048961", "EMG4418640984775", "EMG4418624396784"},
			"queryType": "2",
			"pageNo":    1,
			"pageSize":  100,
			"session":   client.AccessToken,
		})
	})
}

func PrintRes(f func() (interface{}, error)) {
	resp, err := f()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Response:", resp)
	}
}
