package app

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/app"
    appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type PurchasesApi struct {}



// CreatePurchases 创建purchases
// @Tags Purchases
// @Summary 创建purchases
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body app.Purchases true "创建purchases"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /purchases/createPurchases [post]
func (purchasesApi *PurchasesApi) CreatePurchases(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var purchases app.Purchases
	err := c.ShouldBindJSON(&purchases)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = purchasesService.CreatePurchases(ctx,&purchases)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeletePurchases 删除purchases
// @Tags Purchases
// @Summary 删除purchases
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body app.Purchases true "删除purchases"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /purchases/deletePurchases [delete]
func (purchasesApi *PurchasesApi) DeletePurchases(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := purchasesService.DeletePurchases(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePurchasesByIds 批量删除purchases
// @Tags Purchases
// @Summary 批量删除purchases
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /purchases/deletePurchasesByIds [delete]
func (purchasesApi *PurchasesApi) DeletePurchasesByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := purchasesService.DeletePurchasesByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePurchases 更新purchases
// @Tags Purchases
// @Summary 更新purchases
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body app.Purchases true "更新purchases"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /purchases/updatePurchases [put]
func (purchasesApi *PurchasesApi) UpdatePurchases(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var purchases app.Purchases
	err := c.ShouldBindJSON(&purchases)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = purchasesService.UpdatePurchases(ctx,purchases)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindPurchases 用id查询purchases
// @Tags Purchases
// @Summary 用id查询purchases
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询purchases"
// @Success 200 {object} response.Response{data=app.Purchases,msg=string} "查询成功"
// @Router /purchases/findPurchases [get]
func (purchasesApi *PurchasesApi) FindPurchases(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	repurchases, err := purchasesService.GetPurchases(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(repurchases, c)
}
// GetPurchasesList 分页获取purchases列表
// @Tags Purchases
// @Summary 分页获取purchases列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query appReq.PurchasesSearch true "分页获取purchases列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /purchases/getPurchasesList [get]
func (purchasesApi *PurchasesApi) GetPurchasesList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo appReq.PurchasesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := purchasesService.GetPurchasesInfoList(ctx,pageInfo)
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

// GetPurchasesPublic 不需要鉴权的purchases接口
// @Tags Purchases
// @Summary 不需要鉴权的purchases接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /purchases/getPurchasesPublic [get]
func (purchasesApi *PurchasesApi) GetPurchasesPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    purchasesService.GetPurchasesPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的purchases接口信息",
    }, "获取成功", c)
}
