
package app

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
    appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
)

type PurchasesService struct {}
// CreatePurchases 创建purchases记录
// Author [yourname](https://github.com/yourname)
func (purchasesService *PurchasesService) CreatePurchases(ctx context.Context, purchases *app.Purchases) (err error) {
	err = global.GVA_DB.Create(purchases).Error
	return err
}

// DeletePurchases 删除purchases记录
// Author [yourname](https://github.com/yourname)
func (purchasesService *PurchasesService)DeletePurchases(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&app.Purchases{},"id = ?",ID).Error
	return err
}

// DeletePurchasesByIds 批量删除purchases记录
// Author [yourname](https://github.com/yourname)
func (purchasesService *PurchasesService)DeletePurchasesByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]app.Purchases{},"id in ?",IDs).Error
	return err
}

// UpdatePurchases 更新purchases记录
// Author [yourname](https://github.com/yourname)
func (purchasesService *PurchasesService)UpdatePurchases(ctx context.Context, purchases app.Purchases) (err error) {
	err = global.GVA_DB.Model(&app.Purchases{}).Where("id = ?",purchases.ID).Updates(&purchases).Error
	return err
}

// GetPurchases 根据ID获取purchases记录
// Author [yourname](https://github.com/yourname)
func (purchasesService *PurchasesService)GetPurchases(ctx context.Context, ID string) (purchases app.Purchases, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&purchases).Error
	return
}
// GetPurchasesInfoList 分页获取purchases记录
// Author [yourname](https://github.com/yourname)
func (purchasesService *PurchasesService)GetPurchasesInfoList(ctx context.Context, info appReq.PurchasesSearch) (list []app.Purchases, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&app.Purchases{})
    var purchasess []app.Purchases
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.PurchaseID != nil && *info.PurchaseID != "" {
        db = db.Where("purchaseID = ?",*info.PurchaseID)
    }
    if info.ProductID != nil && *info.ProductID != "" {
        db = db.Where("productID = ?",*info.ProductID)
    }
    if info.SupplierID != nil && *info.SupplierID != "" {
        db = db.Where("supplierID = ?",*info.SupplierID)
    }
    if info.WarehouseID != nil && *info.WarehouseID != "" {
        db = db.Where("warehouseID = ?",*info.WarehouseID)
    }
    if info.WarehouseName != nil && *info.WarehouseName != "" {
        db = db.Where("warehouseName = ?",*info.WarehouseName)
    }
    if info.Status != nil {
        db = db.Where("status = ?",*info.Status)
    }
    if info.UserID != nil {
        db = db.Where("userID = ?",*info.UserID)
    }
    if info.UserName != nil && *info.UserName != "" {
        db = db.Where("userName = ?",*info.UserName)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&purchasess).Error
	return  purchasess, total, err
}
func (purchasesService *PurchasesService)GetPurchasesPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
