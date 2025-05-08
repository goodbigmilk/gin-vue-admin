package app

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	UserApi
	SupplierApi
	WarehouseApi
	ProductApi
	CategoryApi
	PurchasesApi
}

var (
	USService        = service.ServiceGroupApp.AppServiceGroup.UserService
	SPService        = service.ServiceGroupApp.AppServiceGroup.SupplierService
	warehouseService = service.ServiceGroupApp.AppServiceGroup.WarehouseService
	productService   = service.ServiceGroupApp.AppServiceGroup.ProductService
	categoryService  = service.ServiceGroupApp.AppServiceGroup.CategoryService
	purchasesService = service.ServiceGroupApp.AppServiceGroup.PurchasesService
)
