package app

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/app"
    appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type UserApi struct {}



// CreateUser 创建用户信息
// @Tags User
// @Summary 创建用户信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body app.User true "创建用户信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /US/createUser [post]
func (USApi *UserApi) CreateUser(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var US app.User
	err := c.ShouldBindJSON(&US)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = USService.CreateUser(ctx,&US)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteUser 删除用户信息
// @Tags User
// @Summary 删除用户信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body app.User true "删除用户信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /US/deleteUser [delete]
func (USApi *UserApi) DeleteUser(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := USService.DeleteUser(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteUserByIds 批量删除用户信息
// @Tags User
// @Summary 批量删除用户信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /US/deleteUserByIds [delete]
func (USApi *UserApi) DeleteUserByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := USService.DeleteUserByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateUser 更新用户信息
// @Tags User
// @Summary 更新用户信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body app.User true "更新用户信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /US/updateUser [put]
func (USApi *UserApi) UpdateUser(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var US app.User
	err := c.ShouldBindJSON(&US)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = USService.UpdateUser(ctx,US)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindUser 用id查询用户信息
// @Tags User
// @Summary 用id查询用户信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询用户信息"
// @Success 200 {object} response.Response{data=app.User,msg=string} "查询成功"
// @Router /US/findUser [get]
func (USApi *UserApi) FindUser(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reUS, err := USService.GetUser(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reUS, c)
}
// GetUserList 分页获取用户信息列表
// @Tags User
// @Summary 分页获取用户信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query appReq.UserSearch true "分页获取用户信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /US/getUserList [get]
func (USApi *UserApi) GetUserList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo appReq.UserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := USService.GetUserInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetUserPublic 不需要鉴权的用户信息接口
// @Tags User
// @Summary 不需要鉴权的用户信息接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /US/getUserPublic [get]
func (USApi *UserApi) GetUserPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    USService.GetUserPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的用户信息接口信息",
    }, "获取成功", c)
}
