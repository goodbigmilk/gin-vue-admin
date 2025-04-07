
// 自动生成模板User
package app
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 用户信息 结构体  User
type User struct {
    global.GVA_MODEL
    UserId  *string `json:"userId" form:"userId" gorm:"column:userId;comment:用户id;"`  //用户id
    UserName  *string `json:"userName" form:"userName" gorm:"column:userName;comment:用户名;"`  //用户名
    Phone  *string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;"`  //手机号
    Password  *string `json:"password" form:"password" gorm:"column:password;comment:密码;"`  //密码
    HeadImg  string `json:"headImg" form:"headImg" gorm:"column:headImg;comment:头像;"`  //头像
    Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态 0 使用中， 1 注销;"`  //状态
}


// TableName 用户信息 User自定义表名 user
func (User) TableName() string {
    return "user"
}





