package eclpIbdOrderDeclareOrder

// CustomsOrder 定义了跨境订单申报的主结构体
type CustomsOrder struct {
	params map[string]interface{}

	IsvUUID                 string   `json:"isvUUID"`                      // 商家订单号
	IsvSource               string   `json:"isvSource"`                    // ISV来源编号
	PlatformId              string   `json:"platformId"`                   // 三方平台编号
	PlatformName            string   `json:"platformName"`                 // 三方平台名称
	PlatformType            string   `json:"platformType"`                 // 销售平台类型
	SpSoNo                  string   `json:"spSoNo"`                       // 销售平台订单号
	DeptNo                  string   `json:"deptNo"`                       // 事业部编号
	InJdwms                 string   `json:"inJdwms"`                      // 是否入京仓
	SalesPlatformCreateTime string   `json:"salesPlatformCreateTime"`      // 订单创建时间
	VenderId                *string  `json:"venderId,omitempty"`           // 商家编号（可选）
	VenderName              *string  `json:"venderName,omitempty"`         // 商家名称（可选）
	ConsigneeName           string   `json:"consigneeName"`                // 收货人姓名
	ConsigneeMobile         string   `json:"consigneeMobile"`              // 收货人手机（可选）
	ConsigneePhone          *string  `json:"consigneePhone,omitempty"`     // 收货人电话（可选）
	ConsigneeEmail          *string  `json:"consigneeEmail,omitempty"`     // 收货人邮箱（可选）
	ConsigneeAddress        string   `json:"consigneeAddress"`             // 收货人地址
	ConsigneePostcode       *string  `json:"consigneePostcode,omitempty"`  // 收货人邮政编码（可选）
	ConsigneeCountry        string   `json:"consigneeCountry"`             // 收货人所在国家(地区)代码
	AddressProvince         string   `json:"addressProvince"`              // 收货地址-省名称
	AddressCity             string   `json:"addressCity"`                  // 收货地址-市名称
	AddressCounty           string   `json:"addressCounty"`                // 收货地址-县名称
	AddressTown             *string  `json:"addressTown,omitempty"`        // 收货地址-镇名称（可选）
	SoType                  *string  `json:"soType,omitempty"`             // 订单类型（可选）
	ExpectDate              *string  `json:"expectDate,omitempty"`         // 期望发货时间（可选）
	InvoiceTitle            *string  `json:"invoiceTitle,omitempty"`       // 发票抬头（可选）
	InvoiceContent          *string  `json:"invoiceContent,omitempty"`     // 发票内容（可选）
	DeclareOrder            string   `json:"declareOrder"`                 // 是否申报订单
	CcProvider              string   `json:"ccProvider"`                   // 清关服务商编码
	CcProviderName          string   `json:"ccProviderName"`               // 清关服务商名称
	PostType                string   `json:"postType"`                     // 电商平台的订单类型
	Pattern                 string   `json:"pattern"`                      // 跨境业务模式
	Customs                 string   `json:"customs"`                      // 保税区编码
	WarehouseNo             string   `json:"warehouseNo"`                  // 京东仓库编号
	EbpCode                 *string  `json:"ebpCode,omitempty"`            // 电商平台代码（可选）
	EbpName                 *string  `json:"ebpName,omitempty"`            // 电商平台名称（可选）
	EbcCode                 string   `json:"ebcCode,omitempty"`            // 电商企业代码（可选）
	EbcName                 string   `json:"ebcName,omitempty"`            // 电商企业名称（可选）
	Delivery                string   `json:"delivery,omitempty"`           // 直邮发货地（可选）
	Discount                float64  `json:"discount"`                     // 优惠券减免金额
	DiscountNote            *string  `json:"discountNote,omitempty"`       // 优惠说明（可选）
	Istax                   string   `json:"istax"`                        // 是否包税
	TaxTotal                float64  `json:"taxTotal"`                     // 税费
	Freight                 float64  `json:"freight"`                      // 运费
	OtherPrice              float64  `json:"otherPrice"`                   // 其他费用
	GoodsValue              float64  `json:"goodsValue"`                   // 商品价格
	Weight                  float64  `json:"weight"`                       // 订单总毛重
	NetWeight               float64  `json:"netWeight"`                    // 订单总净重
	BatchNumbers            string   `json:"batchNumbers,omitempty"`       // 商品批次号（可选）
	BuyerRegNo              string   `json:"buyerRegNo"`                   // 订购人注册号
	BuyerPhone              string   `json:"buyerPhone"`                   // 订购人电话
	BuyerName               string   `json:"buyerName"`                    // 订购人姓名
	BuyerIdType             string   `json:"buyerIdType"`                  // 订购人证件类型
	BuyerIdNumber           string   `json:"buyerIdNumber"`                // 订购人证件号码
	SenderName              string   `json:"senderName"`                   // 发件人名称
	SenderCompanyName       string   `json:"senderCompanyName"`            // 发件方公司名称
	SenderCountry           string   `json:"senderCountry"`                // 发件地国家
	SenderZip               string   `json:"senderZip"`                    // 发件地邮编
	SenderCity              string   `json:"senderCity"`                   // 发件地城市
	SenderProvince          string   `json:"senderProvince"`               // 发件地省/州名
	SenderTel               string   `json:"senderTel"`                    // 发件人电话
	SenderAddr              string   `json:"senderAddr"`                   // 发件人地址
	CustomsRemark           *string  `json:"customsRemark,omitempty"`      // 申报备注（可选）
	DeclarePaymentList      string   `json:"declarePaymentList"`           // 是否申报支付单
	PaymentType             string   `json:"paymentType"`                  // 支付方式
	PayCode                 string   `json:"payCode"`                      // 支付企业代码
	PayName                 string   `json:"payName"`                      // 支付企业名称
	PayTransactionId        string   `json:"payTransactionId"`             // 支付流水号
	Currency                string   `json:"currency"`                     // 支付币制
	PaymentConfirmTime      string   `json:"paymentConfirmTime"`           // 支付时间
	ShouldPay               float64  `json:"shouldPay"`                    // 实际支付金额
	ReceiveNo               *string  `json:"receiveNo,omitempty"`          // 收件人证件号（可选）
	PayRemark               *string  `json:"payRemark,omitempty"`          // 支付备注（可选）
	DeclareWaybill          string   `json:"declareWaybill"`               // 是否申报运单
	LogisticsCode           string   `json:"logisticsCode"`                // 物流企业代码
	LogisticsName           string   `json:"logisticsName"`                // 物流企业名称
	BdOwnerNo               *string  `json:"bdOwnerNo,omitempty"`          // 商家京配编号（可选）
	LogisticsNo             *string  `json:"logisticsNo,omitempty"`        // 物流运单编号（可选）
	PackNo                  int      `json:"packNo"`                       // 包裹数
	LogisticsRemark         *string  `json:"logisticsRemark,omitempty"`    // 物流备注（可选）
	IsDelivery              int      `json:"isDelivery"`                   // 是否货到付款
	Receivable              *float64 `json:"receivable,omitempty"`         // 订单应收金额（可选）
	ConsigneeRemark         *string  `json:"consigneeRemark,omitempty"`    // 客户留言（可选）
	InsuredPriceFlag        *string  `json:"insuredPriceFlag,omitempty"`   // 是否保价（可选）
	InsuredValue            *float64 `json:"insuredValue,omitempty"`       // 保价声明价值（可选）
	InsuredFee              float64  `json:"insuredFee"`                   // 保价费用
	ShopNo                  string   `json:"shopNo"`                       // 店铺编号
	IsSupervise             int      `json:"isSupervise,omitempty"`        // 是否代监管支付（可选）
	InitalRequest           *string  `json:"initalRequest,omitempty"`      // 原始请求报文（可选）
	InitalResponse          *string  `json:"initalResponse,omitempty"`     // 原始响应报文（可选）
	PayTransactionIdYh      *string  `json:"payTransactionIdYh,omitempty"` // 支付交易流水号（可选）
	IsvParentId             *string  `json:"isvParentId,omitempty"`        // 商家母订单号（可选）
	IsvOrderIdList          *string  `json:"isvOrderIdList,omitempty"`     // 商家子订单号集合（可选）
	TotalAmount             *float64 `json:"totalAmount,omitempty"`        // 交易金额（可选）
	VerDept                 *int     `json:"verDept,omitempty"`            // 核验机构（可选）
	PayType                 *int     `json:"payType,omitempty"`            // 支付类型（可选）
	RecpAccount             *string  `json:"recpAccount,omitempty"`        // 收款账号（可选）
	RecpCode                *string  `json:"recpCode,omitempty"`           // 收款企业代码（可选）
	RecpName                *string  `json:"recpName,omitempty"`           // 收款企业名称（可选）
	ConsNameEN              *string  `json:"consNameEN,omitempty"`         // 收货人英文名称（可选）
	ConsAddressEN           *string  `json:"consAddressEN,omitempty"`      // 收货人英文地址（可选）
	SenderNameEN            *string  `json:"senderNameEN,omitempty"`       // 发件人（公司）英文名称（可选）
	SenderCityEN            *string  `json:"senderCityEN,omitempty"`       // 发件人英文城市（可选）
	SenderAddrEN            *string  `json:"senderAddrEN,omitempty"`       // 发件人英文地址（可选）
	ConsigneeIdType         *string  `json:"consigneeIdType,omitempty"`    // 证件类型（可选）
	WrapType                *string  `json:"wrapType,omitempty"`           // 包装种类（可选）
	Oaid                    *string  `json:"oaid,omitempty"`               // 京东零售订单oaid（可选）
}

func NewCustomsOrder() *CustomsOrder {
	return &CustomsOrder{
		params: map[string]interface{}{
			"@type": "com.jd.eclp.isv.domain.so.OrderCustomsParam",
		},
	}
}

// GetParams 获取完整参数映射
func (c *CustomsOrder) GetParams() map[string]interface{} {
	return c.params
}
