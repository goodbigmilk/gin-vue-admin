package app

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	UserRouter
	SupplierRouter
	WarehouseRouter
	ProductRouter
	CategoryRouter
	PurchasesRouter
}

var (
	USApi        = api.ApiGroupApp.AppApiGroup.UserApi
	SPApi        = api.ApiGroupApp.AppApiGroup.SupplierApi
	warehouseApi = api.ApiGroupApp.AppApiGroup.WarehouseApi
	productApi   = api.ApiGroupApp.AppApiGroup.ProductApi
	categoryApi  = api.ApiGroupApp.AppApiGroup.CategoryApi
	purchasesApi = api.ApiGroupApp.AppApiGroup.PurchasesApi
)
