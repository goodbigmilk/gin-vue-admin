
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type UserSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    UserId  *string `json:"userId" form:"userId" `
    UserName  *string `json:"userName" form:"userName" `
    Phone  *string `json:"phone" form:"phone" `
    request.PageInfo
}
