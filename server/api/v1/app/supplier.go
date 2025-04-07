package app

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/app"
    appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type SupplierApi struct {}



// CreateSupplier 创建供应商
// @Tags Supplier
// @Summary 创建供应商
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body app.Supplier true "创建供应商"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /SP/createSupplier [post]
func (SPApi *SupplierApi) CreateSupplier(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var SP app.Supplier
	err := c.ShouldBindJSON(&SP)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = SPService.CreateSupplier(ctx,&SP)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteSupplier 删除供应商
// @Tags Supplier
// @Summary 删除供应商
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body app.Supplier true "删除供应商"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /SP/deleteSupplier [delete]
func (SPApi *SupplierApi) DeleteSupplier(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := SPService.DeleteSupplier(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSupplierByIds 批量删除供应商
// @Tags Supplier
// @Summary 批量删除供应商
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /SP/deleteSupplierByIds [delete]
func (SPApi *SupplierApi) DeleteSupplierByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := SPService.DeleteSupplierByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateSupplier 更新供应商
// @Tags Supplier
// @Summary 更新供应商
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body app.Supplier true "更新供应商"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /SP/updateSupplier [put]
func (SPApi *SupplierApi) UpdateSupplier(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var SP app.Supplier
	err := c.ShouldBindJSON(&SP)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = SPService.UpdateSupplier(ctx,SP)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSupplier 用id查询供应商
// @Tags Supplier
// @Summary 用id查询供应商
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询供应商"
// @Success 200 {object} response.Response{data=app.Supplier,msg=string} "查询成功"
// @Router /SP/findSupplier [get]
func (SPApi *SupplierApi) FindSupplier(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reSP, err := SPService.GetSupplier(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reSP, c)
}
// GetSupplierList 分页获取供应商列表
// @Tags Supplier
// @Summary 分页获取供应商列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query appReq.SupplierSearch true "分页获取供应商列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /SP/getSupplierList [get]
func (SPApi *SupplierApi) GetSupplierList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo appReq.SupplierSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := SPService.GetSupplierInfoList(ctx,pageInfo)
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

// GetSupplierPublic 不需要鉴权的供应商接口
// @Tags Supplier
// @Summary 不需要鉴权的供应商接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /SP/getSupplierPublic [get]
func (SPApi *SupplierApi) GetSupplierPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    SPService.GetSupplierPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的供应商接口信息",
    }, "获取成功", c)
}
