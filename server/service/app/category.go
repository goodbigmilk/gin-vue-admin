
package app

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
    appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
)

type CategoryService struct {}
// CreateCategory 创建category记录
// Author [yourname](https://github.com/yourname)
func (categoryService *CategoryService) CreateCategory(ctx context.Context, category *app.Category) (err error) {
	err = global.GVA_DB.Create(category).Error
	return err
}

// DeleteCategory 删除category记录
// Author [yourname](https://github.com/yourname)
func (categoryService *CategoryService)DeleteCategory(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&app.Category{},"id = ?",ID).Error
	return err
}

// DeleteCategoryByIds 批量删除category记录
// Author [yourname](https://github.com/yourname)
func (categoryService *CategoryService)DeleteCategoryByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]app.Category{},"id in ?",IDs).Error
	return err
}

// UpdateCategory 更新category记录
// Author [yourname](https://github.com/yourname)
func (categoryService *CategoryService)UpdateCategory(ctx context.Context, category app.Category) (err error) {
	err = global.GVA_DB.Model(&app.Category{}).Where("id = ?",category.ID).Updates(&category).Error
	return err
}

// GetCategory 根据ID获取category记录
// Author [yourname](https://github.com/yourname)
func (categoryService *CategoryService)GetCategory(ctx context.Context, ID string) (category app.Category, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&category).Error
	return
}
// GetCategoryInfoList 分页获取category记录
// Author [yourname](https://github.com/yourname)
func (categoryService *CategoryService)GetCategoryInfoList(ctx context.Context, info appReq.CategorySearch) (list []app.Category, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&app.Category{})
    var categorys []app.Category
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.CategoryID != nil && *info.CategoryID != "" {
        db = db.Where("categoryID = ?",*info.CategoryID)
    }
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name = ?",*info.Name)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&categorys).Error
	return  categorys, total, err
}
func (categoryService *CategoryService)GetCategoryPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
