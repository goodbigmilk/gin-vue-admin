
// 自动生成模板Warehouse
package app
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 仓库 结构体  Warehouse
type Warehouse struct {
    global.GVA_MODEL
    WarehouseID  *string `json:"warehouseID" form:"warehouseID" gorm:"column:warehouseID;comment:warehouseID;"`  //warehouseID
    Name  *string `json:"name" form:"name" gorm:"column:name;comment:name;"`  //name
    Status  *int `json:"status" form:"status" gorm:"column:status;comment:status;"`  //status
    UserID  *int `json:"userID" form:"userID" gorm:"column:userID;comment:userID;"`  //userID
}


// TableName 仓库 Warehouse自定义表名 warehouse
func (Warehouse) TableName() string {
    return "warehouse"
}





