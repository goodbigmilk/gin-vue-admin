
// 自动生成模板Supplier
package app
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 供应商 结构体  Supplier
type Supplier struct {
    global.GVA_MODEL
    SupplierID  *string `json:"supplierID" form:"supplierID" gorm:"column:supplierID;comment:供应商id;"`  //供应商id
    SupplierName  *string `json:"supplierName" form:"supplierName" gorm:"column:supplierName;comment:supplierName;"`  //supplierName
    Phone  *string `json:"phone" form:"phone" gorm:"column:phone;comment:电话号码;"`  //电话号码
    UserID  *string `json:"userID" form:"userID" gorm:"column:userID;comment:属于哪个用户;"`  //属于哪个用户
    Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态 0合作中，1暂停合作，2黑名单;"`  //状态 0合作中，1暂停合作，2黑名单
    CumulativeOrder  *int `json:"cumulativeOrder" form:"cumulativeOrder" gorm:"column:cumulativeOrder;comment:累计订单;"`  //累计订单
    Remark  *string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;type:text;"`  //备注
    IsDelete  *int `json:"isDelete" form:"isDelete" gorm:"column:isDelete;comment:是否软删;"`  //是否软删
}


// TableName 供应商 Supplier自定义表名 supplier
func (Supplier) TableName() string {
    return "supplier"
}





