
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ProductSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    ProductID  *string `json:"productID" form:"productID" `
    Name  *string `json:"name" form:"name" `
    Price  *int `json:"price" form:"price" `
    SupplierID  *string `json:"supplierID" form:"supplierID" `
    CategoryName  *string `json:"categoryName" form:"categoryName" `
    CategoryID  *string `json:"categoryID" form:"categoryID" `
    request.PageInfo
}
