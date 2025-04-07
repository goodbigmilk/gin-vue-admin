package app

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	UserApi
	SupplierApi
}

var (
	USService = service.ServiceGroupApp.AppServiceGroup.UserService
	SPService = service.ServiceGroupApp.AppServiceGroup.SupplierService
)
