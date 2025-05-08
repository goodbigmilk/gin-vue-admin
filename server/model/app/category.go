
// 自动生成模板Category
package app
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// category 结构体  Category
type Category struct {
    global.GVA_MODEL
    CategoryID  *string `json:"categoryID" form:"categoryID" gorm:"column:categoryID;comment:categoryID;"`  //categoryID
    Name  *string `json:"name" form:"name" gorm:"column:name;comment:name;"`  //name
    Description  *string `json:"description" form:"description" gorm:"column:description;comment:description;type:text;"`  //description
}


// TableName category Category自定义表名 category
func (Category) TableName() string {
    return "category"
}





