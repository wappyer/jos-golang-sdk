package eclpIbdOrderDeclareOrder

// Goods 定义了商品信息的结构体
type Goods struct {
	apiParas map[string]interface{}

	Gnum        int     `json:"gnum"`        // 商品序号
	IsvGoodsNo  string  `json:"isvGoodsNo"`  // 商家商品编码
	SpGoodsNo   string  `json:"spGoodsNo"`   // 销售平台商品编码
	Quantity    int     `json:"quantity"`    // 数量
	Price       float64 `json:"price"`       // 单价
	GoodsRemark string  `json:"goodsRemark"` // 备注（可选）
	ItemLink    string  `json:"itemLink"`    // 商品展示链接地址（可选）
}

func NewGoods() *Goods {
	return &Goods{
		apiParas: map[string]interface{}{
			"@type": "java.util.List",
		},
	}
}

// GetApiParas 获取完整参数映射
func (r *Goods) GetApiParas() map[string]interface{} {
	return r.apiParas
}

func (r *Goods) SetGnum(gnum int) {
	r.Gnum = gnum
	r.apiParas["gnum"] = gnum
}

func (r *Goods) SetIsvGoodsNo(isvGoodsNo string) {
	r.IsvGoodsNo = isvGoodsNo
	r.apiParas["isvGoodsNo"] = isvGoodsNo
}

func (r *Goods) SetSpGoodsNo(spGoodsNo string) {
	r.SpGoodsNo = spGoodsNo
	r.apiParas["spGoodsNo"] = spGoodsNo
}

func (r *Goods) SetQuantity(quantity int) {
	r.Quantity = quantity
	r.apiParas["quantity"] = quantity
}

func (r *Goods) SetPrice(price float64) {
	r.Price = price
	r.apiParas["price"] = price
}

func (r *Goods) SetGoodsRemark(goodsRemark string) {
	r.GoodsRemark = goodsRemark
	r.apiParas["goodsRemark"] = goodsRemark
}

func (r *Goods) SetItemLink(itemLink string) {
	r.ItemLink = itemLink
	r.apiParas["itemLink"] = itemLink
}
