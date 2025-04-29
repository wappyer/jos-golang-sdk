package eclpIbdOrderDeclareOrder

// CustomsOrder 定义了跨境订单申报的主结构体
type CustomsOrder struct {
	apiParas map[string]interface{}

	IsvUUID                 string  `json:"isvUUID"`                 // 商家订单号
	IsvSource               string  `json:"isvSource"`               // ISV来源编号
	PlatformId              string  `json:"platformId"`              // 三方平台编号
	PlatformName            string  `json:"platformName"`            // 三方平台名称
	PlatformType            string  `json:"platformType"`            // 销售平台类型
	SpSoNo                  string  `json:"spSoNo"`                  // 销售平台订单号
	DeptNo                  string  `json:"deptNo"`                  // 事业部编号
	InJdwms                 string  `json:"inJdwms"`                 // 是否入京仓
	SalesPlatformCreateTime string  `json:"salesPlatformCreateTime"` // 订单创建时间
	VenderId                string  `json:"venderId"`                // 商家编号（可选）
	VenderName              string  `json:"venderName"`              // 商家名称（可选）
	ConsigneeName           string  `json:"consigneeName"`           // 收货人姓名
	ConsigneeMobile         string  `json:"consigneeMobile"`         // 收货人手机（可选）
	ConsigneePhone          string  `json:"consigneePhone"`          // 收货人电话（可选）
	ConsigneeEmail          string  `json:"consigneeEmail"`          // 收货人邮箱（可选）
	ConsigneeAddress        string  `json:"consigneeAddress"`        // 收货人地址
	ConsigneePostcode       string  `json:"consigneePostcode"`       // 收货人邮政编码（可选）
	ConsigneeCountry        string  `json:"consigneeCountry"`        // 收货人所在国家(地区)代码
	AddressProvince         string  `json:"addressProvince"`         // 收货地址-省名称
	AddressCity             string  `json:"addressCity"`             // 收货地址-市名称
	AddressCounty           string  `json:"addressCounty"`           // 收货地址-县名称
	AddressTown             string  `json:"addressTown"`             // 收货地址-镇名称（可选）
	SoType                  string  `json:"soType"`                  // 订单类型（可选）
	ExpectDate              string  `json:"expectDate"`              // 期望发货时间（可选）
	InvoiceTitle            string  `json:"invoiceTitle"`            // 发票抬头（可选）
	InvoiceContent          string  `json:"invoiceContent"`          // 发票内容（可选）
	DeclareOrder            string  `json:"declareOrder"`            // 是否申报订单
	CcProvider              string  `json:"ccProvider"`              // 清关服务商编码
	CcProviderName          string  `json:"ccProviderName"`          // 清关服务商名称
	PostType                string  `json:"postType"`                // 电商平台的订单类型
	Pattern                 string  `json:"pattern"`                 // 跨境业务模式
	Customs                 string  `json:"customs"`                 // 保税区编码
	WarehouseNo             string  `json:"warehouseNo"`             // 京东仓库编号
	EbpCode                 string  `json:"ebpCode"`                 // 电商平台代码（可选）
	EbpName                 string  `json:"ebpName"`                 // 电商平台名称（可选）
	EbcCode                 string  `json:"ebcCode"`                 // 电商企业代码（可选）
	EbcName                 string  `json:"ebcName"`                 // 电商企业名称（可选）
	Delivery                string  `json:"delivery"`                // 直邮发货地（可选）
	Discount                float64 `json:"discount"`                // 优惠券减免金额
	DiscountNote            string  `json:"discountNote"`            // 优惠说明（可选）
	Istax                   string  `json:"istax"`                   // 是否包税
	TaxTotal                float64 `json:"taxTotal"`                // 税费
	Freight                 float64 `json:"freight"`                 // 运费
	OtherPrice              float64 `json:"otherPrice"`              // 其他费用
	GoodsValue              float64 `json:"goodsValue"`              // 商品价格
	Weight                  float64 `json:"weight"`                  // 订单总毛重
	NetWeight               float64 `json:"netWeight"`               // 订单总净重
	BatchNumbers            string  `json:"batchNumbers"`            // 商品批次号（可选）
	BuyerRegNo              string  `json:"buyerRegNo"`              // 订购人注册号
	BuyerPhone              string  `json:"buyerPhone"`              // 订购人电话
	BuyerName               string  `json:"buyerName"`               // 订购人姓名
	BuyerIdType             string  `json:"buyerIdType"`             // 订购人证件类型
	BuyerIdNumber           string  `json:"buyerIdNumber"`           // 订购人证件号码
	SenderName              string  `json:"senderName"`              // 发件人名称
	SenderCompanyName       string  `json:"senderCompanyName"`       // 发件方公司名称
	SenderCountry           string  `json:"senderCountry"`           // 发件地国家
	SenderZip               string  `json:"senderZip"`               // 发件地邮编
	SenderCity              string  `json:"senderCity"`              // 发件地城市
	SenderProvince          string  `json:"senderProvince"`          // 发件地省/州名
	SenderTel               string  `json:"senderTel"`               // 发件人电话
	SenderAddr              string  `json:"senderAddr"`              // 发件人地址
	CustomsRemark           string  `json:"customsRemark"`           // 申报备注（可选）
	DeclarePaymentList      string  `json:"declarePaymentList"`      // 是否申报支付单
	PaymentType             string  `json:"paymentType"`             // 支付方式
	PayCode                 string  `json:"payCode"`                 // 支付企业代码
	PayName                 string  `json:"payName"`                 // 支付企业名称
	PayTransactionId        string  `json:"payTransactionId"`        // 支付流水号
	Currency                string  `json:"currency"`                // 支付币制
	PaymentConfirmTime      string  `json:"paymentConfirmTime"`      // 支付时间
	ShouldPay               float64 `json:"shouldPay"`               // 实际支付金额
	ReceiveNo               string  `json:"receiveNo"`               // 收件人证件号（可选）
	PayRemark               string  `json:"payRemark"`               // 支付备注（可选）
	DeclareWaybill          string  `json:"declareWaybill"`          // 是否申报运单
	LogisticsCode           string  `json:"logisticsCode"`           // 物流企业代码
	LogisticsName           string  `json:"logisticsName"`           // 物流企业名称
	BdOwnerNo               string  `json:"bdOwnerNo"`               // 商家京配编号（可选）
	LogisticsNo             string  `json:"logisticsNo"`             // 物流运单编号（可选）
	PackNo                  int     `json:"packNo"`                  // 包裹数
	LogisticsRemark         string  `json:"logisticsRemark"`         // 物流备注（可选）
	IsDelivery              int     `json:"isDelivery"`              // 是否货到付款
	Receivable              float64 `json:"receivable"`              // 订单应收金额（可选）
	ConsigneeRemark         string  `json:"consigneeRemark"`         // 客户留言（可选）
	InsuredPriceFlag        string  `json:"insuredPriceFlag"`        // 是否保价（可选）
	InsuredValue            float64 `json:"insuredValue"`            // 保价声明价值（可选）
	InsuredFee              float64 `json:"insuredFee"`              // 保价费用
	ShopNo                  string  `json:"shopNo"`                  // 店铺编号
	IsSupervise             int     `json:"isSupervise"`             // 是否代监管支付（可选）
	InitalRequest           string  `json:"initalRequest"`           // 原始请求报文（可选）
	InitalResponse          string  `json:"initalResponse"`          // 原始响应报文（可选）
	PayTransactionIdYh      string  `json:"payTransactionIdYh"`      // 支付交易流水号（可选）
	IsvParentId             string  `json:"isvParentId"`             // 商家母订单号（可选）
	IsvOrderIdList          string  `json:"isvOrderIdList"`          // 商家子订单号集合（可选）
	TotalAmount             float64 `json:"totalAmount"`             // 交易金额（可选）
	VerDept                 int     `json:"verDept"`                 // 核验机构（可选）
	PayType                 int     `json:"payType"`                 // 支付类型（可选）
	RecpAccount             string  `json:"recpAccount"`             // 收款账号（可选）
	RecpCode                string  `json:"recpCode"`                // 收款企业代码（可选）
	RecpName                string  `json:"recpName"`                // 收款企业名称（可选）
	ConsNameEN              string  `json:"consNameEN"`              // 收货人英文名称（可选）
	ConsAddressEN           string  `json:"consAddressEN"`           // 收货人英文地址（可选）
	SenderNameEN            string  `json:"senderNameEN"`            // 发件人（公司）英文名称（可选）
	SenderCityEN            string  `json:"senderCityEN"`            // 发件人英文城市（可选）
	SenderAddrEN            string  `json:"senderAddrEN"`            // 发件人英文地址（可选）
	ConsigneeIdType         string  `json:"consigneeIdType"`         // 证件类型（可选）
	WrapType                string  `json:"wrapType"`                // 包装种类（可选）
	Oaid                    string  `json:"oaid"`                    // 京东零售订单oaid（可选）
}

func NewCustomsOrder() *CustomsOrder {
	return &CustomsOrder{
		apiParas: map[string]interface{}{
			"@type": "com.jd.eclp.isv.domain.so.OrderCustomsParam",
		},
	}
}

// GetApiParas 获取完整参数映射
func (r CustomsOrder) GetApiParas() map[string]interface{} {
	return r.apiParas
}

func (r CustomsOrder) SetIsvUUID(isvUUID string) {
	r.IsvUUID = isvUUID
	r.apiParas["isvUUID"] = isvUUID
}

func (r CustomsOrder) SetIsvSource(isvSource string) {
	r.IsvSource = isvSource
	r.apiParas["isvSource"] = isvSource
}

func (r CustomsOrder) SetPlatformId(platformId string) {
	r.PlatformId = platformId
	r.apiParas["platformId"] = platformId
}

func (r CustomsOrder) SetPlatformName(platformName string) {
	r.PlatformName = platformName
	r.apiParas["platformName"] = platformName
}

func (r CustomsOrder) SetPlatformType(platformType string) {
	r.PlatformType = platformType
	r.apiParas["platformType"] = platformType
}

func (r CustomsOrder) SetSpSoNo(spSoNo string) {
	r.SpSoNo = spSoNo
	r.apiParas["spSoNo"] = spSoNo
}

func (r CustomsOrder) SetDeptNo(deptNo string) {
	r.DeptNo = deptNo
	r.apiParas["deptNo"] = deptNo
}

func (r CustomsOrder) SetInJdwms(inJdwms string) {
	r.InJdwms = inJdwms
	r.apiParas["inJdwms"] = inJdwms
}

func (r CustomsOrder) SetSalesPlatformCreateTime(salesPlatformCreateTime string) {
	r.SalesPlatformCreateTime = salesPlatformCreateTime
	r.apiParas["salesPlatformCreateTime"] = salesPlatformCreateTime
}

func (r CustomsOrder) SetVenderId(venderId string) {
	r.VenderId = venderId
	r.apiParas["venderId"] = venderId
}

func (r CustomsOrder) SetVenderName(venderName string) {
	r.VenderName = venderName
	r.apiParas["venderName"] = venderName
}

func (r CustomsOrder) SetConsigneeName(consigneeName string) {
	r.ConsigneeName = consigneeName
	r.apiParas["consigneeName"] = consigneeName
}

func (r CustomsOrder) SetConsigneeMobile(consigneeMobile string) {
	r.ConsigneeMobile = consigneeMobile
	r.apiParas["consigneeMobile"] = consigneeMobile
}

func (r CustomsOrder) SetConsigneePhone(consigneePhone string) {
	r.ConsigneePhone = consigneePhone
	r.apiParas["consigneePhone"] = consigneePhone
}

func (r CustomsOrder) SetConsigneeEmail(consigneeEmail string) {
	r.ConsigneeEmail = consigneeEmail
	r.apiParas["consigneeEmail"] = consigneeEmail
}

func (r CustomsOrder) SetConsigneeAddress(consigneeAddress string) {
	r.ConsigneeAddress = consigneeAddress
	r.apiParas["consigneeAddress"] = consigneeAddress
}

func (r CustomsOrder) SetConsigneePostcode(consigneePostcode string) {
	r.ConsigneePostcode = consigneePostcode
	r.apiParas["consigneePostcode"] = consigneePostcode
}

func (r CustomsOrder) SetConsigneeCountry(consigneeCountry string) {
	r.ConsigneeCountry = consigneeCountry
	r.apiParas["consigneeCountry"] = consigneeCountry
}

func (r CustomsOrder) SetAddressProvince(addressProvince string) {
	r.AddressProvince = addressProvince
	r.apiParas["addressProvince"] = addressProvince
}

func (r CustomsOrder) SetAddressCity(addressCity string) {
	r.AddressCity = addressCity
	r.apiParas["addressCity"] = addressCity
}

func (r CustomsOrder) SetAddressCounty(addressCounty string) {
	r.AddressCounty = addressCounty
	r.apiParas["addressCounty"] = addressCounty
}

func (r CustomsOrder) SetAddressTown(addressTown string) {
	r.AddressTown = addressTown
	r.apiParas["addressTown"] = addressTown
}

func (r CustomsOrder) SetSoType(soType string) {
	r.SoType = soType
	r.apiParas["soType"] = soType
}

func (r CustomsOrder) SetExpectDate(expectDate string) {
	r.ExpectDate = expectDate
	r.apiParas["expectDate"] = expectDate
}

func (r CustomsOrder) SetInvoiceTitle(invoiceTitle string) {
	r.InvoiceTitle = invoiceTitle
	r.apiParas["invoiceTitle"] = invoiceTitle
}

func (r CustomsOrder) SetInvoiceContent(invoiceContent string) {
	r.InvoiceContent = invoiceContent
	r.apiParas["invoiceContent"] = invoiceContent
}

func (r CustomsOrder) SetDeclareOrder(declareOrder string) {
	r.DeclareOrder = declareOrder
	r.apiParas["declareOrder"] = declareOrder
}

func (r CustomsOrder) SetCcProvider(ccProvider string) {
	r.CcProvider = ccProvider
	r.apiParas["ccProvider"] = ccProvider
}

func (r CustomsOrder) SetCcProviderName(ccProviderName string) {
	r.CcProviderName = ccProviderName
	r.apiParas["ccProviderName"] = ccProviderName
}

func (r CustomsOrder) SetPostType(postType string) {
	r.PostType = postType
	r.apiParas["postType"] = postType
}

func (r CustomsOrder) SetPattern(pattern string) {
	r.Pattern = pattern
	r.apiParas["pattern"] = pattern
}

func (r CustomsOrder) SetCustoms(customs string) {
	r.Customs = customs
	r.apiParas["customs"] = customs
}

func (r CustomsOrder) SetWarehouseNo(warehouseNo string) {
	r.WarehouseNo = warehouseNo
	r.apiParas["warehouseNo"] = warehouseNo
}

func (r CustomsOrder) SetEbpCode(ebpCode string) {
	r.EbpCode = ebpCode
	r.apiParas["ebpCode"] = ebpCode
}

func (r CustomsOrder) SetEbpName(ebpName string) {
	r.EbpName = ebpName
	r.apiParas["ebpName"] = ebpName
}

func (r CustomsOrder) SetEbcCode(ebcCode string) {
	r.EbcCode = ebcCode
	r.apiParas["ebcCode"] = ebcCode
}

func (r CustomsOrder) SetEbcName(ebcName string) {
	r.EbcName = ebcName
	r.apiParas["ebcName"] = ebcName
}

func (r CustomsOrder) SetDelivery(delivery string) {
	r.Delivery = delivery
	r.apiParas["delivery"] = delivery
}

func (r CustomsOrder) SetDiscount(discount float64) {
	r.Discount = discount
	r.apiParas["discount"] = discount
}

func (r CustomsOrder) SetDiscountNote(discountNote string) {
	r.DiscountNote = discountNote
	r.apiParas["discountNote"] = discountNote
}

func (r CustomsOrder) SetIstax(istax string) {
	r.Istax = istax
	r.apiParas["istax"] = istax
}

func (r CustomsOrder) SetTaxTotal(taxTotal float64) {
	r.TaxTotal = taxTotal
	r.apiParas["taxTotal"] = taxTotal
}

func (r CustomsOrder) SetFreight(freight float64) {
	r.Freight = freight
	r.apiParas["freight"] = freight
}

func (r CustomsOrder) SetOtherPrice(otherPrice float64) {
	r.OtherPrice = otherPrice
	r.apiParas["otherPrice"] = otherPrice
}

func (r CustomsOrder) SetGoodsValue(goodsValue float64) {
	r.GoodsValue = goodsValue
	r.apiParas["goodsValue"] = goodsValue
}

func (r CustomsOrder) SetWeight(weight float64) {
	r.Weight = weight
	r.apiParas["weight"] = weight
}

func (r CustomsOrder) SetNetWeight(netWeight float64) {
	r.NetWeight = netWeight
	r.apiParas["netWeight"] = netWeight
}

func (r CustomsOrder) SetBatchNumbers(batchNumbers string) {
	r.BatchNumbers = batchNumbers
	r.apiParas["batchNumbers"] = batchNumbers
}

func (r CustomsOrder) SetBuyerRegNo(buyerRegNo string) {
	r.BuyerRegNo = buyerRegNo
	r.apiParas["buyerRegNo"] = buyerRegNo
}

func (r CustomsOrder) SetBuyerPhone(buyerPhone string) {
	r.BuyerPhone = buyerPhone
	r.apiParas["buyerPhone"] = buyerPhone
}

func (r CustomsOrder) SetBuyerName(buyerName string) {
	r.BuyerName = buyerName
	r.apiParas["buyerName"] = buyerName
}

func (r CustomsOrder) SetBuyerIdType(buyerIdType string) {
	r.BuyerIdType = buyerIdType
	r.apiParas["buyerIdType"] = buyerIdType
}

func (r CustomsOrder) SetBuyerIdNumber(buyerIdNumber string) {
	r.BuyerIdNumber = buyerIdNumber
	r.apiParas["buyerIdNumber"] = buyerIdNumber
}

func (r CustomsOrder) SetSenderName(senderName string) {
	r.SenderName = senderName
	r.apiParas["senderName"] = senderName
}

func (r CustomsOrder) SetSenderCompanyName(senderCompanyName string) {
	r.SenderCompanyName = senderCompanyName
	r.apiParas["senderCompanyName"] = senderCompanyName
}

func (r CustomsOrder) SetSenderCountry(senderCountry string) {
	r.SenderCountry = senderCountry
	r.apiParas["senderCountry"] = senderCountry
}

func (r CustomsOrder) SetSenderZip(senderZip string) {
	r.SenderZip = senderZip
	r.apiParas["senderZip"] = senderZip
}

func (r CustomsOrder) SetSenderCity(senderCity string) {
	r.SenderCity = senderCity
	r.apiParas["senderCity"] = senderCity
}

func (r CustomsOrder) SetSenderProvince(senderProvince string) {
	r.SenderProvince = senderProvince
	r.apiParas["senderProvince"] = senderProvince
}

func (r CustomsOrder) SetSenderTel(senderTel string) {
	r.SenderTel = senderTel
	r.apiParas["senderTel"] = senderTel
}

func (r CustomsOrder) SetSenderAddr(senderAddr string) {
	r.SenderAddr = senderAddr
	r.apiParas["senderAddr"] = senderAddr
}

func (r CustomsOrder) SetCustomsRemark(customsRemark string) {
	r.CustomsRemark = customsRemark
	r.apiParas["customsRemark"] = customsRemark
}

func (r CustomsOrder) SetDeclarePaymentList(declarePaymentList string) {
	r.DeclarePaymentList = declarePaymentList
	r.apiParas["declarePaymentList"] = declarePaymentList
}

func (r CustomsOrder) SetPaymentType(paymentType string) {
	r.PaymentType = paymentType
	r.apiParas["paymentType"] = paymentType
}

func (r CustomsOrder) SetPayCode(payCode string) {
	r.PayCode = payCode
	r.apiParas["payCode"] = payCode
}

func (r CustomsOrder) SetPayName(payName string) {
	r.PayName = payName
	r.apiParas["payName"] = payName
}

func (r CustomsOrder) SetPayTransactionId(payTransactionId string) {
	r.PayTransactionId = payTransactionId
	r.apiParas["payTransactionId"] = payTransactionId
}

func (r CustomsOrder) SetCurrency(currency string) {
	r.Currency = currency
	r.apiParas["currency"] = currency
}

func (r CustomsOrder) SetPaymentConfirmTime(paymentConfirmTime string) {
	r.PaymentConfirmTime = paymentConfirmTime
	r.apiParas["paymentConfirmTime"] = paymentConfirmTime
}

func (r CustomsOrder) SetShouldPay(shouldPay float64) {
	r.ShouldPay = shouldPay
	r.apiParas["shouldPay"] = shouldPay
}

func (r CustomsOrder) SetReceiveNo(receiveNo string) {
	r.ReceiveNo = receiveNo
	r.apiParas["receiveNo"] = receiveNo
}

func (r CustomsOrder) SetPayRemark(payRemark string) {
	r.PayRemark = payRemark
	r.apiParas["payRemark"] = payRemark
}

func (r CustomsOrder) SetDeclareWaybill(declareWaybill string) {
	r.DeclareWaybill = declareWaybill
	r.apiParas["declareWaybill"] = declareWaybill
}

func (r CustomsOrder) SetLogisticsCode(logisticsCode string) {
	r.LogisticsCode = logisticsCode
	r.apiParas["logisticsCode"] = logisticsCode
}

func (r CustomsOrder) SetLogisticsName(logisticsName string) {
	r.LogisticsName = logisticsName
	r.apiParas["logisticsName"] = logisticsName
}

func (r CustomsOrder) SetBdOwnerNo(bdOwnerNo string) {
	r.BdOwnerNo = bdOwnerNo
	r.apiParas["bdOwnerNo"] = bdOwnerNo
}

func (r CustomsOrder) SetLogisticsNo(logisticsNo string) {
	r.LogisticsNo = logisticsNo
	r.apiParas["logisticsNo"] = logisticsNo
}

func (r CustomsOrder) SetPackNo(packNo int) {
	r.PackNo = packNo
	r.apiParas["packNo"] = packNo
}

func (r CustomsOrder) SetLogisticsRemark(logisticsRemark string) {
	r.LogisticsRemark = logisticsRemark
	r.apiParas["logisticsRemark"] = logisticsRemark
}

func (r CustomsOrder) SetIsDelivery(isDelivery int) {
	r.IsDelivery = isDelivery
	r.apiParas["isDelivery"] = isDelivery
}

func (r CustomsOrder) SetReceivable(receivable float64) {
	r.Receivable = receivable
	r.apiParas["receivable"] = receivable
}

func (r CustomsOrder) SetConsigneeRemark(consigneeRemark string) {
	r.ConsigneeRemark = consigneeRemark
	r.apiParas["consigneeRemark"] = consigneeRemark
}

func (r CustomsOrder) SetInsuredPriceFlag(insuredPriceFlag string) {
	r.InsuredPriceFlag = insuredPriceFlag
	r.apiParas["insuredPriceFlag"] = insuredPriceFlag
}

func (r CustomsOrder) SetInsuredValue(insuredValue float64) {
	r.InsuredValue = insuredValue
	r.apiParas["insuredValue"] = insuredValue
}

func (r CustomsOrder) SetInsuredFee(insuredFee float64) {
	r.InsuredFee = insuredFee
	r.apiParas["insuredFee"] = insuredFee
}

func (r CustomsOrder) SetShopNo(shopNo string) {
	r.ShopNo = shopNo
	r.apiParas["shopNo"] = shopNo
}

func (r CustomsOrder) SetIsSupervise(isSupervise int) {
	r.IsSupervise = isSupervise
	r.apiParas["isSupervise"] = isSupervise
}

func (r CustomsOrder) SetInitalRequest(initalRequest string) {
	r.InitalRequest = initalRequest
	r.apiParas["initalRequest"] = initalRequest
}

func (r CustomsOrder) SetInitalResponse(initalResponse string) {
	r.InitalResponse = initalResponse
	r.apiParas["initalResponse"] = initalResponse
}

func (r CustomsOrder) SetPayTransactionIdYh(payTransactionIdYh string) {
	r.PayTransactionIdYh = payTransactionIdYh
	r.apiParas["payTransactionIdYh"] = payTransactionIdYh
}

func (r CustomsOrder) SetIsvParentId(isvParentId string) {
	r.IsvParentId = isvParentId
	r.apiParas["isvParentId"] = isvParentId
}

func (r CustomsOrder) SetIsvOrderIdList(isvOrderIdList string) {
	r.IsvOrderIdList = isvOrderIdList
	r.apiParas["isvOrderIdList"] = isvOrderIdList
}

func (r CustomsOrder) SetTotalAmount(totalAmount float64) {
	r.TotalAmount = totalAmount
	r.apiParas["totalAmount"] = totalAmount
}

func (r CustomsOrder) SetVerDept(verDept int) {
	r.VerDept = verDept
	r.apiParas["verDept"] = verDept
}

func (r CustomsOrder) SetPayType(payType int) {
	r.PayType = payType
	r.apiParas["payType"] = payType
}

func (r CustomsOrder) SetRecpAccount(recpAccount string) {
	r.RecpAccount = recpAccount
	r.apiParas["recpAccount"] = recpAccount
}

func (r CustomsOrder) SetRecpCode(recpCode string) {
	r.RecpCode = recpCode
	r.apiParas["recpCode"] = recpCode
}

func (r CustomsOrder) SetRecpName(recpName string) {
	r.RecpName = recpName
	r.apiParas["recpName"] = recpName
}

func (r CustomsOrder) SetConsNameEN(consNameEN string) {
	r.ConsNameEN = consNameEN
	r.apiParas["consNameEN"] = consNameEN
}

func (r CustomsOrder) SetConsAddressEN(consAddressEN string) {
	r.ConsAddressEN = consAddressEN
	r.apiParas["consAddressEN"] = consAddressEN
}

func (r CustomsOrder) SetSenderNameEN(senderNameEN string) {
	r.SenderNameEN = senderNameEN
	r.apiParas["senderNameEN"] = senderNameEN
}

func (r CustomsOrder) SetSenderCityEN(senderCityEN string) {
	r.SenderCityEN = senderCityEN
	r.apiParas["senderCityEN"] = senderCityEN
}

func (r CustomsOrder) SetSenderAddrEN(senderAddrEN string) {
	r.SenderAddrEN = senderAddrEN
	r.apiParas["senderAddrEN"] = senderAddrEN
}

func (r CustomsOrder) SetConsigneeIdType(consigneeIdType string) {
	r.ConsigneeIdType = consigneeIdType
	r.apiParas["consigneeIdType"] = consigneeIdType
}

func (r CustomsOrder) SetWrapType(wrapType string) {
	r.WrapType = wrapType
	r.apiParas["wrapType"] = wrapType
}

func (r CustomsOrder) SetOaid(oaid string) {
	r.Oaid = oaid
	r.apiParas["oaid"] = oaid
}
