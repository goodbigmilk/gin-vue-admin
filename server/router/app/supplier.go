package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SupplierRouter struct {}

// InitSupplierRouter 初始化 供应商 路由信息
func (s *SupplierRouter) InitSupplierRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	SPRouter := Router.Group("SP").Use(middleware.OperationRecord())
	SPRouterWithoutRecord := Router.Group("SP")
	SPRouterWithoutAuth := PublicRouter.Group("SP")
	{
		SPRouter.POST("createSupplier", SPApi.CreateSupplier)   // 新建供应商
		SPRouter.DELETE("deleteSupplier", SPApi.DeleteSupplier) // 删除供应商
		SPRouter.DELETE("deleteSupplierByIds", SPApi.DeleteSupplierByIds) // 批量删除供应商
		SPRouter.PUT("updateSupplier", SPApi.UpdateSupplier)    // 更新供应商
	}
	{
		SPRouterWithoutRecord.GET("findSupplier", SPApi.FindSupplier)        // 根据ID获取供应商
		SPRouterWithoutRecord.GET("getSupplierList", SPApi.GetSupplierList)  // 获取供应商列表
	}
	{
	    SPRouterWithoutAuth.GET("getSupplierPublic", SPApi.GetSupplierPublic)  // 供应商开放接口
	}
}
