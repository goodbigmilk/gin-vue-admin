package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WarehouseRouter struct {}

// InitWarehouseRouter 初始化 仓库 路由信息
func (s *WarehouseRouter) InitWarehouseRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	warehouseRouter := Router.Group("warehouse").Use(middleware.OperationRecord())
	warehouseRouterWithoutRecord := Router.Group("warehouse")
	warehouseRouterWithoutAuth := PublicRouter.Group("warehouse")
	{
		warehouseRouter.POST("createWarehouse", warehouseApi.CreateWarehouse)   // 新建仓库
		warehouseRouter.DELETE("deleteWarehouse", warehouseApi.DeleteWarehouse) // 删除仓库
		warehouseRouter.DELETE("deleteWarehouseByIds", warehouseApi.DeleteWarehouseByIds) // 批量删除仓库
		warehouseRouter.PUT("updateWarehouse", warehouseApi.UpdateWarehouse)    // 更新仓库
	}
	{
		warehouseRouterWithoutRecord.GET("findWarehouse", warehouseApi.FindWarehouse)        // 根据ID获取仓库
		warehouseRouterWithoutRecord.GET("getWarehouseList", warehouseApi.GetWarehouseList)  // 获取仓库列表
	}
	{
	    warehouseRouterWithoutAuth.GET("getWarehousePublic", warehouseApi.GetWarehousePublic)  // 仓库开放接口
	}
}
