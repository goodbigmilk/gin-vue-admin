package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CategoryRouter struct {}

// InitCategoryRouter 初始化 category 路由信息
func (s *CategoryRouter) InitCategoryRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	categoryRouter := Router.Group("category").Use(middleware.OperationRecord())
	categoryRouterWithoutRecord := Router.Group("category")
	categoryRouterWithoutAuth := PublicRouter.Group("category")
	{
		categoryRouter.POST("createCategory", categoryApi.CreateCategory)   // 新建category
		categoryRouter.DELETE("deleteCategory", categoryApi.DeleteCategory) // 删除category
		categoryRouter.DELETE("deleteCategoryByIds", categoryApi.DeleteCategoryByIds) // 批量删除category
		categoryRouter.PUT("updateCategory", categoryApi.UpdateCategory)    // 更新category
	}
	{
		categoryRouterWithoutRecord.GET("findCategory", categoryApi.FindCategory)        // 根据ID获取category
		categoryRouterWithoutRecord.GET("getCategoryList", categoryApi.GetCategoryList)  // 获取category列表
	}
	{
	    categoryRouterWithoutAuth.GET("getCategoryPublic", categoryApi.GetCategoryPublic)  // category开放接口
	}
}
