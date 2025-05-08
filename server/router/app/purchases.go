package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PurchasesRouter struct {}

// InitPurchasesRouter 初始化 purchases 路由信息
func (s *PurchasesRouter) InitPurchasesRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	purchasesRouter := Router.Group("purchases").Use(middleware.OperationRecord())
	purchasesRouterWithoutRecord := Router.Group("purchases")
	purchasesRouterWithoutAuth := PublicRouter.Group("purchases")
	{
		purchasesRouter.POST("createPurchases", purchasesApi.CreatePurchases)   // 新建purchases
		purchasesRouter.DELETE("deletePurchases", purchasesApi.DeletePurchases) // 删除purchases
		purchasesRouter.DELETE("deletePurchasesByIds", purchasesApi.DeletePurchasesByIds) // 批量删除purchases
		purchasesRouter.PUT("updatePurchases", purchasesApi.UpdatePurchases)    // 更新purchases
	}
	{
		purchasesRouterWithoutRecord.GET("findPurchases", purchasesApi.FindPurchases)        // 根据ID获取purchases
		purchasesRouterWithoutRecord.GET("getPurchasesList", purchasesApi.GetPurchasesList)  // 获取purchases列表
	}
	{
	    purchasesRouterWithoutAuth.GET("getPurchasesPublic", purchasesApi.GetPurchasesPublic)  // purchases开放接口
	}
}
