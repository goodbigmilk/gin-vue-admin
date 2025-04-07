
package app

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
    appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
)

type UserService struct {}
// CreateUser 创建用户信息记录
// Author [yourname](https://github.com/yourname)
func (USService *UserService) CreateUser(ctx context.Context, US *app.User) (err error) {
	err = global.GVA_DB.Create(US).Error
	return err
}

// DeleteUser 删除用户信息记录
// Author [yourname](https://github.com/yourname)
func (USService *UserService)DeleteUser(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&app.User{},"id = ?",ID).Error
	return err
}

// DeleteUserByIds 批量删除用户信息记录
// Author [yourname](https://github.com/yourname)
func (USService *UserService)DeleteUserByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]app.User{},"id in ?",IDs).Error
	return err
}

// UpdateUser 更新用户信息记录
// Author [yourname](https://github.com/yourname)
func (USService *UserService)UpdateUser(ctx context.Context, US app.User) (err error) {
	err = global.GVA_DB.Model(&app.User{}).Where("id = ?",US.ID).Updates(&US).Error
	return err
}

// GetUser 根据ID获取用户信息记录
// Author [yourname](https://github.com/yourname)
func (USService *UserService)GetUser(ctx context.Context, ID string) (US app.User, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&US).Error
	return
}
// GetUserInfoList 分页获取用户信息记录
// Author [yourname](https://github.com/yourname)
func (USService *UserService)GetUserInfoList(ctx context.Context, info appReq.UserSearch) (list []app.User, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&app.User{})
    var USs []app.User
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.UserId != nil && *info.UserId != "" {
        db = db.Where("userId = ?",*info.UserId)
    }
    if info.UserName != nil && *info.UserName != "" {
        db = db.Where("userName = ?",*info.UserName)
    }
    if info.Phone != nil && *info.Phone != "" {
        db = db.Where("phone = ?",*info.Phone)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&USs).Error
	return  USs, total, err
}
func (USService *UserService)GetUserPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
