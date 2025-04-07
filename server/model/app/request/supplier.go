
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type SupplierSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    SupplierID  *string `json:"supplierID" form:"supplierID" `
    SupplierName  *string `json:"supplierName" form:"supplierName" `
    Phone  *string `json:"phone" form:"phone" `
    UserID  *string `json:"userID" form:"userID" `
    CumulativeOrder  *int `json:"cumulativeOrder" form:"cumulativeOrder" `
    IsDelete  *int `json:"isDelete" form:"isDelete" `
    request.PageInfo
}
