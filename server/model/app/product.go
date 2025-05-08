
// 自动生成模板Product
package app
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// 商品 结构体  Product
type Product struct {
    global.GVA_MODEL
    ProductID  *string `json:"productID" form:"productID" gorm:"column:productID;comment:productID;"`  //productID
    Name  *string `json:"name" form:"name" gorm:"column:name;comment:name;"`  //name
    Price  *int `json:"price" form:"price" gorm:"column:price;comment:price;"`  //price
    SupplierID  *string `json:"supplierID" form:"supplierID" gorm:"column:supplierID;comment:supplierID;"`  //supplierID
    Description  *string `json:"description" form:"description" gorm:"column:description;comment:description;type:text;"`  //description
    Stock  *int `json:"stock" form:"stock" gorm:"column:stock;comment:stock;"`  //stock
    CategoryName  *string `json:"categoryName" form:"categoryName" gorm:"column:categoryName;comment:类别名称;"`  //类别名称
    CategoryID  *string `json:"categoryID" form:"categoryID" gorm:"column:categoryID;comment:categoryID;"`  //categoryID
    Imgs  datatypes.JSON `json:"imgs" form:"imgs" gorm:"column:imgs;comment:imgs;" swaggertype:"array,object"`  //imgs
    IsDelete  *int `json:"isDelete" form:"isDelete" gorm:"column:isDelete;comment:isDelete;"`  //isDelete
}


// TableName 商品 Product自定义表名 product
func (Product) TableName() string {
    return "product"
}





