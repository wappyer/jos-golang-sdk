package request

import "encoding/json"

type EclpIbdGoodsTransportGoodsInfoRequest struct {
	apiParas map[string]interface{}

	DeptNo          string  `json:"deptNo"`
	IsvGoodsNo      string  `json:"isvGoodsNo"`
	SpGodosNo       string  `json:"spGodosNo"`
	Barcodes        string  `json:"barcodes,omitempty"`
	ThirdCategoryNo string  `json:"thirdCategoryNo"`
	GoodsName       string  `json:"goodsName"`
	BrandNo         string  `json:"brandNo,omitempty"`
	BrandName       string  `json:"brandName,omitempty"`
	GrossWeight     string  `json:"grossWeight,omitempty"`
	NetWeight       float64 `json:"netWeight,omitempty"`
	Length          string  `json:"length,omitempty"`
	Width           string  `json:"width,omitempty"`
	Height          string  `json:"height,omitempty"`
	SafeDays        int     `json:"safeDays,omitempty"`
	Batch           string  `json:"batch,omitempty"`
	OverseaPurchase string  `json:"overseaPurchase,omitempty"`
	QiRecord        string  `json:"qiRecord,omitempty"`
	CustomRecord    string  `json:"customRecord,omitempty"`
	Pattern         string  `json:"pattern,omitempty"`
	CcProvider      string  `json:"ccProvider,omitempty"`
	BondedArea      string  `json:"bondedArea,omitempty"`
	SellerRecord    string  `json:"sellerRecord,omitempty"`
	BatAttrIds      string  `json:"batAttrIds,omitempty"`
	NeedJDRecord    string  `json:"needJDRecord,omitempty"`
	ModelNumber     string  `json:"modelNumber,omitempty"`
	Spe             string  `json:"spe,omitempty"`
	VatRate         string  `json:"vatRate,omitempty"`
	TaxRate         string  `json:"taxRate,omitempty"`
	TaxNumberPost   string  `json:"taxNumberPost,omitempty"`
	PostRate        string  `json:"postRate,omitempty"`
	HsCode          string  `json:"hsCode,omitempty"`
	Country         string  `json:"country,omitempty"`
	QiCountry       string  `json:"qiCountry,omitempty"`
	Flag            string  `json:"flag,omitempty"`
	LegalUnit1      string  `json:"legalUnit1,omitempty"`
	LegalAmount1    string  `json:"legalAmount1,omitempty"`
	LegalUnit2      string  `json:"legalUnit2,omitempty"`
	LegalAmount2    string  `json:"legalAmount2,omitempty"`
	Measurement     string  `json:"measurement,omitempty"`
	QiMeasurement   string  `json:"qiMeasurement,omitempty"`
	Delivery        string  `json:"delivery,omitempty"`
	Hgsbys          string  `json:"hgsbys,omitempty"`
	GoodsNameEn     string  `json:"goodsNameEn,omitempty"`

	responseError ErrorResp
	responseData  EclpIbdGoodsTransportGoodsInfoResponse
}

type EclpIbdGoodsTransportGoodsInfoResponse struct {
	JingdongEclpIbdGoodsTransportGoodsInfoResponce JingdongEclpIbdGoodsTransportGoodsInfoResponce `json:"jingdong_eclp_ibd_goods_transportGoodsInfo_responce"`
}

type JingdongEclpIbdGoodsTransportGoodsInfoResponce struct {
	Code      string `json:"code"`
	GoodsNo   string `json:"goodsNo"`
	RequestID string `json:"request_id"`
}

func NewEclpIbdGoodsTransportGoodsInfoRequest() *EclpIbdGoodsTransportGoodsInfoRequest {
	r := &EclpIbdGoodsTransportGoodsInfoRequest{
		apiParas: make(map[string]interface{}),
	}
	return r
}

func (r *EclpIbdGoodsTransportGoodsInfoRequest) GetApiMethodName() string {
	return "jingdong.eclp.ibd.goods.transportGoodsInfo"
}

func (r *EclpIbdGoodsTransportGoodsInfoRequest) GetApiParas() map[string]interface{} {
	b, _ := json.Marshal(r)
	_ = json.Unmarshal(b, &r.apiParas)
	return r.apiParas
}

func (r *EclpIbdGoodsTransportGoodsInfoRequest) SetResponseError(err ErrorResp) {
	r.responseError = err
}

func (r *EclpIbdGoodsTransportGoodsInfoRequest) GetResponseError() ErrorResp {
	return r.responseError
}

func (r *EclpIbdGoodsTransportGoodsInfoRequest) SetResponseData(data string) error {
	if err := json.Unmarshal([]byte(data), &r.responseData); err != nil {
		return err
	}
	return nil
}

func (r *EclpIbdGoodsTransportGoodsInfoRequest) GetResponseData(responseData interface{}) error {
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

func (r *EclpIbdGoodsTransportGoodsInfoRequest) GetResponse() (EclpIbdGoodsTransportGoodsInfoResponse, ErrorResp) {
	return r.responseData, r.responseError
}
