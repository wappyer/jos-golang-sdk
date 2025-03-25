package eclpIbdOrderDeclareOrder

// Goods 定义了商品信息的结构体
type Goods struct {
	Gnum        int     `json:"gnum"`                  // 商品序号
	IsvGoodsNo  string  `json:"isvGoodsNo"`            // 商家商品编码
	SpGoodsNo   string  `json:"spGoodsNo"`             // 销售平台商品编码
	Quantity    int     `json:"quantity"`              // 数量
	Price       float64 `json:"price"`                 // 单价
	GoodsRemark *string `json:"goodsRemark,omitempty"` // 备注（可选）
	ItemLink    *string `json:"itemLink,omitempty"`    // 商品展示链接地址（可选）
}
