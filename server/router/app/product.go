package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProductRouter struct {}

// InitProductRouter 初始化 商品 路由信息
func (s *ProductRouter) InitProductRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	productRouter := Router.Group("product").Use(middleware.OperationRecord())
	productRouterWithoutRecord := Router.Group("product")
	productRouterWithoutAuth := PublicRouter.Group("product")
	{
		productRouter.POST("createProduct", productApi.CreateProduct)   // 新建商品
		productRouter.DELETE("deleteProduct", productApi.DeleteProduct) // 删除商品
		productRouter.DELETE("deleteProductByIds", productApi.DeleteProductByIds) // 批量删除商品
		productRouter.PUT("updateProduct", productApi.UpdateProduct)    // 更新商品
	}
	{
		productRouterWithoutRecord.GET("findProduct", productApi.FindProduct)        // 根据ID获取商品
		productRouterWithoutRecord.GET("getProductList", productApi.GetProductList)  // 获取商品列表
	}
	{
	    productRouterWithoutAuth.GET("getProductPublic", productApi.GetProductPublic)  // 商品开放接口
	}
}
