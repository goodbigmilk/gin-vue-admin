
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CategorySearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    CategoryID  *string `json:"categoryID" form:"categoryID" `
    Name  *string `json:"name" form:"name" `
    request.PageInfo
}
