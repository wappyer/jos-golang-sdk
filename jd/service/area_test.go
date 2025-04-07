package service

import (
	"fmt"
	"gitee.com/dimension-huimei/jos-golang-sdk/jd"
	"testing"
)

func TestMatchJdProvince(t *testing.T) {
	client := jd.NewClient("https://api-dev.jd.com/routerjson", "BB5039D6E7D11C7A481E092BF6350CB3", "6aac140e6b1f4269a9839776d4611318")
	client.AccessToken = "47fd6890bfe8483ba1377119cbc9d85cy4ot"
	province, errResp, err := MatchJdAddress(client, "是会计法克里江西省斯朵夫", "南昌市房间看电视", "时间的卡夫卡南昌县")
	fmt.Printf("err:%v\n", err)
	fmt.Printf("errResp:%v\n", errResp)
	fmt.Printf("province:%v\n", province)
}
