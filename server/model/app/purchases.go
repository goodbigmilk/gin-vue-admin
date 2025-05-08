
// 自动生成模板Purchases
package app
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// purchases 结构体  Purchases
type Purchases struct {
    global.GVA_MODEL
    PurchaseID  *string `json:"purchaseID" form:"purchaseID" gorm:"column:purchaseID;comment:purchaseID;"`  //purchaseID
    ProductID  *string `json:"productID" form:"productID" gorm:"column:productID;comment:productID;"`  //productID
    SupplierID  *string `json:"supplierID" form:"supplierID" gorm:"column:supplierID;comment:supplierID;"`  //supplierID
    WarehouseID  *string `json:"warehouseID" form:"warehouseID" gorm:"column:warehouseID;comment:warehouseID;"`  //warehouseID
    WarehouseName  *string `json:"warehouseName" form:"warehouseName" gorm:"column:warehouseName;comment:warehouseName;"`  //warehouseName
    Quantity  *int `json:"quantity" form:"quantity" gorm:"column:quantity;comment:quantity;"`  //quantity
    Status  *int `json:"status" form:"status" gorm:"column:status;comment:status;"`  //status
    UserID  *int `json:"userID" form:"userID" gorm:"column:userID;comment:userID;"`  //userID
    UserName  *string `json:"userName" form:"userName" gorm:"column:userName;comment:userName;"`  //userName
}


// TableName purchases Purchases自定义表名 purchases
func (Purchases) TableName() string {
    return "purchases"
}





