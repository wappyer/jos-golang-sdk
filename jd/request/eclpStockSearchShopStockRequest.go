package request

import (
	"encoding/json"
)

/*
 * jingdong.eclp.stock.searchShopStock
 * 店铺库存查询接口
 * 用于商家按照店铺维度查库存，不入库商家不能调用
 */

type EclpStockSearchShopStockRequest struct {
	apiParas    map[string]interface{}
	Version     string `json:"version,omitempty"`
	RequestId   string `json:"requestId,omitempty"`
	DeptNo      string `json:"deptNo,omitempty"`
	ShopNo      string `json:"shopNo,omitempty"`
	WarehouseNo string `json:"warehouseNo,omitempty"`
	GoodsNo     string `json:"goodsNo,omitempty"`
	PageSize    int    `json:"pageSize,omitempty"`
	PageNumber  int    `json:"pageNumber,omitempty"`
}

func (r *EclpStockSearchShopStockRequest) GetApiMethodName() string {
	return "jingdong.eclp.stock.searchShopStock"
}

func (r *EclpStockSearchShopStockRequest) GetApiParas() string {
	if len(r.apiParas) == 0 {
		return "{}"
	}
	data, _ := json.Marshal(r.apiParas)
	return string(data)
}

func (r *EclpStockSearchShopStockRequest) Check() {
	// 实现检查逻辑
}

func (r *EclpStockSearchShopStockRequest) PutOtherTextParam(key string, value interface{}) {
	r.apiParas[key] = value
}

func (r *EclpStockSearchShopStockRequest) SetVersion(version string) {
	r.Version = version
	r.apiParas["version"] = version
}

func (r *EclpStockSearchShopStockRequest) GetVersion() string {
	return r.Version
}

func (r *EclpStockSearchShopStockRequest) SetRequestId(requestId string) {
	r.RequestId = requestId
	r.apiParas["requestId"] = requestId
}

func (r *EclpStockSearchShopStockRequest) GetRequestId() string {
	return r.RequestId
}

func (r *EclpStockSearchShopStockRequest) SetDeptNo(deptNo string) {
	r.DeptNo = deptNo
	r.apiParas["deptNo"] = deptNo
}

func (r *EclpStockSearchShopStockRequest) GetDeptNo() string {
	return r.DeptNo
}

func (r *EclpStockSearchShopStockRequest) SetShopNo(shopNo string) {
	r.ShopNo = shopNo
	r.apiParas["shopNo"] = shopNo
}

func (r *EclpStockSearchShopStockRequest) GetShopNo() string {
	return r.ShopNo
}

func (r *EclpStockSearchShopStockRequest) SetWarehouseNo(warehouseNo string) {
	r.WarehouseNo = warehouseNo
	r.apiParas["warehouseNo"] = warehouseNo
}

func (r *EclpStockSearchShopStockRequest) GetWarehouseNo() string {
	return r.WarehouseNo
}

func (r *EclpStockSearchShopStockRequest) SetGoodsNo(goodsNo string) {
	r.GoodsNo = goodsNo
	r.apiParas["goodsNo"] = goodsNo
}

func (r *EclpStockSearchShopStockRequest) GetGoodsNo() string {
	return r.GoodsNo
}

func (r *EclpStockSearchShopStockRequest) SetPageSize(pageSize int) {
	r.PageSize = pageSize
	r.apiParas["pageSize"] = pageSize
}

func (r *EclpStockSearchShopStockRequest) GetPageSize() int {
	return r.PageSize
}

func (r *EclpStockSearchShopStockRequest) SetPageNumber(pageNumber int) {
	r.PageNumber = pageNumber
	r.apiParas["pageNumber"] = pageNumber
}

func (r *EclpStockSearchShopStockRequest) GetPageNumber() int {
	return r.PageNumber
}
