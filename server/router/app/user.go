package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {}

// InitUserRouter 初始化 用户信息 路由信息
func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	USRouter := Router.Group("US").Use(middleware.OperationRecord())
	USRouterWithoutRecord := Router.Group("US")
	USRouterWithoutAuth := PublicRouter.Group("US")
	{
		USRouter.POST("createUser", USApi.CreateUser)   // 新建用户信息
		USRouter.DELETE("deleteUser", USApi.DeleteUser) // 删除用户信息
		USRouter.DELETE("deleteUserByIds", USApi.DeleteUserByIds) // 批量删除用户信息
		USRouter.PUT("updateUser", USApi.UpdateUser)    // 更新用户信息
	}
	{
		USRouterWithoutRecord.GET("findUser", USApi.FindUser)        // 根据ID获取用户信息
		USRouterWithoutRecord.GET("getUserList", USApi.GetUserList)  // 获取用户信息列表
	}
	{
	    USRouterWithoutAuth.GET("getUserPublic", USApi.GetUserPublic)  // 用户信息开放接口
	}
}
