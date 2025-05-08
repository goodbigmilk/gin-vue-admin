
package app

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
    appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
)

type ProductService struct {}
// CreateProduct 创建商品记录
// Author [yourname](https://github.com/yourname)
func (productService *ProductService) CreateProduct(ctx context.Context, product *app.Product) (err error) {
	err = global.GVA_DB.Create(product).Error
	return err
}

// DeleteProduct 删除商品记录
// Author [yourname](https://github.com/yourname)
func (productService *ProductService)DeleteProduct(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&app.Product{},"id = ?",ID).Error
	return err
}

// DeleteProductByIds 批量删除商品记录
// Author [yourname](https://github.com/yourname)
func (productService *ProductService)DeleteProductByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]app.Product{},"id in ?",IDs).Error
	return err
}

// UpdateProduct 更新商品记录
// Author [yourname](https://github.com/yourname)
func (productService *ProductService)UpdateProduct(ctx context.Context, product app.Product) (err error) {
	err = global.GVA_DB.Model(&app.Product{}).Where("id = ?",product.ID).Updates(&product).Error
	return err
}

// GetProduct 根据ID获取商品记录
// Author [yourname](https://github.com/yourname)
func (productService *ProductService)GetProduct(ctx context.Context, ID string) (product app.Product, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&product).Error
	return
}
// GetProductInfoList 分页获取商品记录
// Author [yourname](https://github.com/yourname)
func (productService *ProductService)GetProductInfoList(ctx context.Context, info appReq.ProductSearch) (list []app.Product, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&app.Product{})
    var products []app.Product
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.ProductID != nil && *info.ProductID != "" {
        db = db.Where("productID = ?",*info.ProductID)
    }
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name = ?",*info.Name)
    }
    if info.Price != nil {
        db = db.Where("price <> ?",*info.Price)
    }
    if info.SupplierID != nil && *info.SupplierID != "" {
        db = db.Where("supplierID = ?",*info.SupplierID)
    }
    if info.CategoryName != nil && *info.CategoryName != "" {
        db = db.Where("categoryName = ?",*info.CategoryName)
    }
    if info.CategoryID != nil && *info.CategoryID != "" {
        db = db.Where("categoryID = ?",*info.CategoryID)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&products).Error
	return  products, total, err
}
func (productService *ProductService)GetProductPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
