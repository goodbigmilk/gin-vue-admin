import service from '@/utils/request'
// @Tags Purchases
// @Summary 创建purchases
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Purchases true "创建purchases"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /purchases/createPurchases [post]
export const createPurchases = (data) => {
  return service({
    url: '/purchases/createPurchases',
    method: 'post',
    data
  })
}

// @Tags Purchases
// @Summary 删除purchases
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Purchases true "删除purchases"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /purchases/deletePurchases [delete]
export const deletePurchases = (params) => {
  return service({
    url: '/purchases/deletePurchases',
    method: 'delete',
    params
  })
}

// @Tags Purchases
// @Summary 批量删除purchases
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除purchases"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /purchases/deletePurchases [delete]
export const deletePurchasesByIds = (params) => {
  return service({
    url: '/purchases/deletePurchasesByIds',
    method: 'delete',
    params
  })
}

// @Tags Purchases
// @Summary 更新purchases
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Purchases true "更新purchases"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /purchases/updatePurchases [put]
export const updatePurchases = (data) => {
  return service({
    url: '/purchases/updatePurchases',
    method: 'put',
    data
  })
}

// @Tags Purchases
// @Summary 用id查询purchases
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Purchases true "用id查询purchases"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /purchases/findPurchases [get]
export const findPurchases = (params) => {
  return service({
    url: '/purchases/findPurchases',
    method: 'get',
    params
  })
}

// @Tags Purchases
// @Summary 分页获取purchases列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取purchases列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /purchases/getPurchasesList [get]
export const getPurchasesList = (params) => {
  return service({
    url: '/purchases/getPurchasesList',
    method: 'get',
    params
  })
}

// @Tags Purchases
// @Summary 不需要鉴权的purchases接口
// @Accept application/json
// @Produce application/json
// @Param data query appReq.PurchasesSearch true "分页获取purchases列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /purchases/getPurchasesPublic [get]
export const getPurchasesPublic = () => {
  return service({
    url: '/purchases/getPurchasesPublic',
    method: 'get',
  })
}
