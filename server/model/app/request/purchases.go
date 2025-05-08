
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type PurchasesSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    PurchaseID  *string `json:"purchaseID" form:"purchaseID" `
    ProductID  *string `json:"productID" form:"productID" `
    SupplierID  *string `json:"supplierID" form:"supplierID" `
    WarehouseID  *string `json:"warehouseID" form:"warehouseID" `
    WarehouseName  *string `json:"warehouseName" form:"warehouseName" `
    Status  *int `json:"status" form:"status" `
    UserID  *int `json:"userID" form:"userID" `
    UserName  *string `json:"userName" form:"userName" `
    request.PageInfo
}
