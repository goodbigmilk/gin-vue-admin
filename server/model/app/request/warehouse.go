
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type WarehouseSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    WarehouseID  *string `json:"warehouseID" form:"warehouseID" `
    Name  *string `json:"name" form:"name" `
    UserID  *int `json:"userID" form:"userID" `
    request.PageInfo
}
