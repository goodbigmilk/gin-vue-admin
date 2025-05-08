
package app

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
    appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
)

type WarehouseService struct {}
// CreateWarehouse 创建仓库记录
// Author [yourname](https://github.com/yourname)
func (warehouseService *WarehouseService) CreateWarehouse(ctx context.Context, warehouse *app.Warehouse) (err error) {
	err = global.GVA_DB.Create(warehouse).Error
	return err
}

// DeleteWarehouse 删除仓库记录
// Author [yourname](https://github.com/yourname)
func (warehouseService *WarehouseService)DeleteWarehouse(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&app.Warehouse{},"id = ?",ID).Error
	return err
}

// DeleteWarehouseByIds 批量删除仓库记录
// Author [yourname](https://github.com/yourname)
func (warehouseService *WarehouseService)DeleteWarehouseByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]app.Warehouse{},"id in ?",IDs).Error
	return err
}

// UpdateWarehouse 更新仓库记录
// Author [yourname](https://github.com/yourname)
func (warehouseService *WarehouseService)UpdateWarehouse(ctx context.Context, warehouse app.Warehouse) (err error) {
	err = global.GVA_DB.Model(&app.Warehouse{}).Where("id = ?",warehouse.ID).Updates(&warehouse).Error
	return err
}

// GetWarehouse 根据ID获取仓库记录
// Author [yourname](https://github.com/yourname)
func (warehouseService *WarehouseService)GetWarehouse(ctx context.Context, ID string) (warehouse app.Warehouse, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&warehouse).Error
	return
}
// GetWarehouseInfoList 分页获取仓库记录
// Author [yourname](https://github.com/yourname)
func (warehouseService *WarehouseService)GetWarehouseInfoList(ctx context.Context, info appReq.WarehouseSearch) (list []app.Warehouse, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&app.Warehouse{})
    var warehouses []app.Warehouse
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.WarehouseID != nil && *info.WarehouseID != "" {
        db = db.Where("warehouseID = ?",*info.WarehouseID)
    }
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name = ?",*info.Name)
    }
    if info.UserID != nil {
        db = db.Where("userID = ?",*info.UserID)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&warehouses).Error
	return  warehouses, total, err
}
func (warehouseService *WarehouseService)GetWarehousePublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
