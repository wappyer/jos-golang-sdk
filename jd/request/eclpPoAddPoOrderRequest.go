package request

import "encoding/json"

type EclpPoAddPoOrderRequest struct {
	apiParas map[string]interface{}

	SpPoOrderNo          string                  `json:"spPoOrderNo"`
	DeptNo               string                  `json:"deptNo"`
	ReferenceOrder       string                  `json:"referenceOrder"`
	InboundRemark        string                  `json:"inboundRemark"`
	Buyer                string                  `json:"buyer"`
	LogicParam           string                  `json:"logicParam"`
	SupplierNo           string                  `json:"supplierNo"`
	SellerSaleOrder      string                  `json:"sellerSaleOrder"`
	SaleOrder            string                  `json:"saleOrder"`
	OrderMark            string                  `json:"orderMark"`
	BillType             string                  `json:"billType"`
	AcceptUnQcFlag       string                  `json:"acceptUnQcFlag"`
	BoxFlag              string                  `json:"boxFlag"`
	EntirePrice          string                  `json:"entirePrice"`
	Boxes                []EclpPoAddPoOrderBox   `json:"boxes,omitempty"`
	PoReturnMode         string                  `json:"poReturnMode"`
	CustomsInfo          string                  `json:"customsInfo"`
	PoType               string                  `json:"poType"`
	BillOfLading         string                  `json:"billOfLading"`
	ReceiveLevel         string                  `json:"receiveLevel"`
	MultiReceivingFlag   string                  `json:"multiReceivingFlag"`
	WaybillNo            string                  `json:"waybillNo"`
	IsvOutWarehouse      string                  `json:"isvOutWarehouse"`
	BizType              string                  `json:"bizType"`
	WaitBoxDetailFlag    string                  `json:"waitBoxDetailFlag"`
	UnitFlag             string                  `json:"unitFlag"`
	SerialDetailMapJson  string                  `json:"serialDetailMapJson"`
	SerialNoScopeMapJson string                  `json:"serialNoScopeMapJson"`
	AllowLackFlag        string                  `json:"allowLackFlag"`
	IsUpdate             string                  `json:"isUpdate"`
	SellerOrderType      string                  `json:"sellerOrderType"`
	CustomField          string                  `json:"customField"`
	SellerWarehouseNo    string                  `json:"sellerWarehouseNo"`
	WhNo                 string                  `json:"whNo"`
	SoNo                 string                  `json:"soNo"`
	GoodsList            []EclpPoAddPoOrderGoods `json:"goodsList,omitempty"`

	responseError ErrorResp
	responseData  interface{}
}

type EclpPoAddPoOrderBox struct {
	BoxNo         string `json:"boxNo"`
	BoxGoodsNo    string `json:"boxGoodsNo"`
	BoxGoodsQty   string `json:"boxGoodsQty"`
	BoxSerialNo   string `json:"boxSerialNo"`
	BoxIsvGoodsNo string `json:"boxIsvGoodsNo"`
}

type EclpPoAddPoOrderGoods struct {
	DeptGoodsNo      string `json:"deptGoodsNo"`
	IsvGoodsNo       string `json:"isvGoodsNo"`
	NumApplication   string `json:"numApplication"`
	GoodsStatus      string `json:"goodsStatus"`
	BarCodeType      string `json:"barCodeType"`
	SidCheckout      string `json:"sidCheckout"`
	UnitPrice        string `json:"unitPrice"`
	TotalPrice       string `json:"totalPrice"`
	QualityCheckRate string `json:"qualityCheckRate"`
	BatAttrListJson  string `json:"batAttrListJson"`
	OrderLine        string `json:"orderLine"`
	IsvLotattrs      string `json:"isvLotattrs"`
	CheckLotattrs    string `json:"checkLotattrs"`
	GoodsPrice       string `json:"goodsPrice"`
	WarehousingFlag  string `json:"warehousingFlag"`
	IsvGoodsUnit     string `json:"isvGoodsUnit"`
}

func NewEclpPoAddPoOrderRequest() *EclpPoAddPoOrderRequest {
	r := &EclpPoAddPoOrderRequest{
		apiParas: make(map[string]interface{}),
	}
	return r
}

func (r *EclpPoAddPoOrderRequest) GetApiMethodName() string {
	return "jingdong.eclp.po.addPoOrder"
}

func (r *EclpPoAddPoOrderRequest) GetApiParas() map[string]interface{} {
	b, _ := json.Marshal(r)
	_ = json.Unmarshal(b, &r.apiParas)
	return r.apiParas
}

func (r *EclpPoAddPoOrderRequest) SetResponseError(err ErrorResp) {
	r.responseError = err
}

func (r *EclpPoAddPoOrderRequest) GetResponseError() ErrorResp {
	return r.responseError
}

func (r *EclpPoAddPoOrderRequest) SetResponseData(data string) error {
	if err := json.Unmarshal([]byte(data), &r.responseData); err != nil {
		return err
	}
	return nil
}

func (r *EclpPoAddPoOrderRequest) GetResponseData(responseData interface{}) error {
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

func (r *EclpPoAddPoOrderRequest) GetResponse() (interface{}, ErrorResp) {
	return r.responseData, r.responseError
}
