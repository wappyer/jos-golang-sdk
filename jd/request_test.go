package jd

import (
	"fmt"
	"gitee.com/dimension-huimei/jos-golang-sdk/jd/request"
	"gitee.com/dimension-huimei/jos-golang-sdk/jd/request/domain/eclpIbdOrderDeclareOrder"
	"testing"
	"time"
)

var client *Client
var deptNo string

func init() {
	client = NewClient("https://api-dev.jd.com/routerjson", "", "")
	client.AccessToken = ""
	deptNo = ""
}

func TestClientExec_EclpGoodsQueryGoodsInfoRequest(t *testing.T) {
	req := request.NewEclpGoodsQueryGoodsInfoRequest()
	req.SetGoodsNos([]string{"", ""})
	req.SetDeptNo(deptNo)
	req.SetQueryType("2")
	req.SetPageNo(1)
	req.SetPageSize(100)
	respData := &request.EclpGoodsQueryGoodsInfoResponse{}
	executeAndLog(req, respData)
}

func TestClientExec_AreasProvinceGetRequest(t *testing.T) {
	req := request.NewAreasProvinceGetRequest()
	respData := &request.AreasProvinceGetResponse{}
	executeAndLog(req, respData)
}

func TestClientExec_AreasCityGetRequest(t *testing.T) {
	req := request.NewAreasCityGetRequest()
	req.SetParentId(21)
	respData := &request.AreasCityGetResponse{}
	executeAndLog(req, respData)
}

func TestClientExec_AreasCountyGetRequest(t *testing.T) {
	req := request.NewAreasCountyGetRequest()
	req.SetParentId(1827)
	respData := &request.AreasCountyGetResponse{}
	executeAndLog(req, respData)
}

func TestClientExec_EclpIbdOrderDeclareOrderRequest(t *testing.T) {
	customsOrder := eclpIbdOrderDeclareOrder.NewCustomsOrder()
	customsOrder.SetIsvUUID("11529515063552")
	customsOrder.SetIsvSource("ISV0020008048912")
	customsOrder.SetPlatformId("8041326")
	customsOrder.SetPlatformName("")
	customsOrder.SetPlatformType("1")
	customsOrder.SetSpSoNo("123")
	customsOrder.SetDeptNo("")
	customsOrder.SetInJdwms("1")
	customsOrder.SetSalesPlatformCreateTime(time.Now().Format("2006-01-02 15:04:05"))
	customsOrder.SetConsigneeName("")
	customsOrder.SetConsigneeMobile("")
	customsOrder.SetConsigneeAddress("生命科学产业园")
	customsOrder.SetConsigneeCountry("143")
	customsOrder.SetAddressProvince("江西")
	customsOrder.SetAddressCity("南昌")
	customsOrder.SetAddressCounty("南昌")
	customsOrder.SetDeclareOrder("2")
	customsOrder.SetCcProvider("010059")
	customsOrder.SetCcProviderName("骏多拉")
	customsOrder.SetPostType("I")
	customsOrder.SetPattern("grkuaijian")
	customsOrder.SetCustoms("guangzhou")
	customsOrder.SetWarehouseNo("118073145")
	customsOrder.SetEbcCode("3301960VWZ")
	customsOrder.SetEbcName("杭州维泰诺曼医药科技有限公司")
	customsOrder.SetDelivery("广东")
	customsOrder.SetDiscount(0)
	customsOrder.SetIstax("0")
	customsOrder.SetTaxTotal(0)
	customsOrder.SetFreight(0)
	customsOrder.SetOtherPrice(0)
	customsOrder.SetGoodsValue(100)
	customsOrder.SetWeight(0)
	customsOrder.SetNetWeight(0)
	customsOrder.SetBuyerRegNo("1")
	customsOrder.SetBuyerPhone("")
	customsOrder.SetBuyerName("")
	customsOrder.SetBuyerIdType("1")
	customsOrder.SetBuyerIdNumber("")
	customsOrder.SetSenderName("汪先生")
	customsOrder.SetSenderCompanyName("")
	customsOrder.SetSenderCountry("601")
	customsOrder.SetSenderZip("999077")
	customsOrder.SetSenderCity("广州市")
	customsOrder.SetSenderProvince("广东")
	customsOrder.SetSenderTel("")
	customsOrder.SetSenderAddr("广州市花都区九湖一队三街与九湖一队四街交叉口东北400米易客满驻场营业点")
	customsOrder.SetDeclarePaymentList("2")
	customsOrder.SetPaymentType("1,1")
	customsOrder.SetPayCode("4403169D3W")
	customsOrder.SetPayName("财付通支付科技有限公司")
	customsOrder.SetPayTransactionId("123")
	customsOrder.SetCurrency("142")
	customsOrder.SetPaymentConfirmTime("20250409105656")
	customsOrder.SetShouldPay(100)
	customsOrder.SetDeclareWaybill("2")
	customsOrder.SetLogisticsCode("CYS0000010")
	customsOrder.SetLogisticsName("京东物流")
	customsOrder.SetPackNo(1)
	customsOrder.SetIsDelivery(0)
	customsOrder.SetInsuredFee(0)
	customsOrder.SetShopNo("ESP0020008986131")
	customsOrder.SetIsSupervise(2)
	customsOrder.SetBatchNumbers("123")
	customsOrder.SetWrapType("22")
	customsOrder.SetConsigneeIdType("1")

	goods := eclpIbdOrderDeclareOrder.NewGoods()
	var goodsList []eclpIbdOrderDeclareOrder.Goods
	goods.SetGnum(1)
	goods.SetIsvGoodsNo("EMG4418502377853")
	goods.SetSpGoodsNo("EMG4418502377853")
	goods.SetQuantity(1)
	goods.SetPrice(100)
	goodsList = append(goodsList, *goods)

	req := request.NewEclpIbdOrderDeclareOrderRequest()
	req.SetCustomsOrder(*customsOrder)
	req.SetGoodsList(goodsList)

	respData := &request.EclpIbdOrderDeclareOrderResponse{}
	executeAndLog(req, respData)
}

func TestClientExec_EclpOrderQueryOrderCustomsRequest(t *testing.T) {
	req := request.NewEclpOrderQueryOrderCustomsRequest()
	req.SetDeptNo("EBU4418055058519")
	req.SetIsvUUID("11560334411520")
	respData := &request.EclpOrderQueryOrderCustomsResponse{}
	executeAndLog(req, respData)
}

func executeAndLog(req request.Request, respData interface{}) {
	fmt.Printf("param:%v\n", req.GetApiParas())
	err := client.Execute(req, client.AccessToken)
	if err != nil {
		fmt.Printf("Error:%#v\n", err)
	} else {
		_ = req.GetResponseData(respData)
		fmt.Printf("ResponseData:%#v\n", respData)
		respErr := req.GetResponseError()
		fmt.Printf("ResponseError:%#v\n", respErr)
	}
}
