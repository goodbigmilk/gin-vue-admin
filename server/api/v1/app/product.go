package app

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/app"
    appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ProductApi struct {}



// CreateProduct 创建商品
// @Tags Product
// @Summary 创建商品
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body app.Product true "创建商品"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /product/createProduct [post]
func (productApi *ProductApi) CreateProduct(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var product app.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = productService.CreateProduct(ctx,&product)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteProduct 删除商品
// @Tags Product
// @Summary 删除商品
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body app.Product true "删除商品"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /product/deleteProduct [delete]
func (productApi *ProductApi) DeleteProduct(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := productService.DeleteProduct(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteProductByIds 批量删除商品
// @Tags Product
// @Summary 批量删除商品
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /product/deleteProductByIds [delete]
func (productApi *ProductApi) DeleteProductByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := productService.DeleteProductByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateProduct 更新商品
// @Tags Product
// @Summary 更新商品
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body app.Product true "更新商品"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /product/updateProduct [put]
func (productApi *ProductApi) UpdateProduct(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var product app.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = productService.UpdateProduct(ctx,product)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindProduct 用id查询商品
// @Tags Product
// @Summary 用id查询商品
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询商品"
// @Success 200 {object} response.Response{data=app.Product,msg=string} "查询成功"
// @Router /product/findProduct [get]
func (productApi *ProductApi) FindProduct(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reproduct, err := productService.GetProduct(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reproduct, c)
}
// GetProductList 分页获取商品列表
// @Tags Product
// @Summary 分页获取商品列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query appReq.ProductSearch true "分页获取商品列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /product/getProductList [get]
func (productApi *ProductApi) GetProductList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo appReq.ProductSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := productService.GetProductInfoList(ctx,pageInfo)
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

// GetProductPublic 不需要鉴权的商品接口
// @Tags Product
// @Summary 不需要鉴权的商品接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /product/getProductPublic [get]
func (productApi *ProductApi) GetProductPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    productService.GetProductPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的商品接口信息",
    }, "获取成功", c)
}
