
package app

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
    appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
)

type SupplierService struct {}
// CreateSupplier 创建供应商记录
// Author [yourname](https://github.com/yourname)
func (SPService *SupplierService) CreateSupplier(ctx context.Context, SP *app.Supplier) (err error) {
	err = global.GVA_DB.Create(SP).Error
	return err
}

// DeleteSupplier 删除供应商记录
// Author [yourname](https://github.com/yourname)
func (SPService *SupplierService)DeleteSupplier(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&app.Supplier{},"id = ?",ID).Error
	return err
}

// DeleteSupplierByIds 批量删除供应商记录
// Author [yourname](https://github.com/yourname)
func (SPService *SupplierService)DeleteSupplierByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]app.Supplier{},"id in ?",IDs).Error
	return err
}

// UpdateSupplier 更新供应商记录
// Author [yourname](https://github.com/yourname)
func (SPService *SupplierService)UpdateSupplier(ctx context.Context, SP app.Supplier) (err error) {
	err = global.GVA_DB.Model(&app.Supplier{}).Where("id = ?",SP.ID).Updates(&SP).Error
	return err
}

// GetSupplier 根据ID获取供应商记录
// Author [yourname](https://github.com/yourname)
func (SPService *SupplierService)GetSupplier(ctx context.Context, ID string) (SP app.Supplier, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&SP).Error
	return
}
// GetSupplierInfoList 分页获取供应商记录
// Author [yourname](https://github.com/yourname)
func (SPService *SupplierService)GetSupplierInfoList(ctx context.Context, info appReq.SupplierSearch) (list []app.Supplier, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&app.Supplier{})
    var SPs []app.Supplier
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.SupplierID != nil && *info.SupplierID != "" {
        db = db.Where("supplierID = ?",*info.SupplierID)
    }
    if info.SupplierName != nil && *info.SupplierName != "" {
        db = db.Where("supplierName = ?",*info.SupplierName)
    }
    if info.Phone != nil && *info.Phone != "" {
        db = db.Where("phone = ?",*info.Phone)
    }
    if info.UserID != nil && *info.UserID != "" {
        db = db.Where("userID = ?",*info.UserID)
    }
    if info.CumulativeOrder != nil {
        db = db.Where("cumulativeOrder <> ?",*info.CumulativeOrder)
    }
    if info.IsDelete != nil {
        db = db.Where("isDelete = ?",*info.IsDelete)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&SPs).Error
	return  SPs, total, err
}
func (SPService *SupplierService)GetSupplierPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
