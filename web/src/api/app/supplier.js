import service from '@/utils/request'
// @Tags Supplier
// @Summary 创建供应商
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Supplier true "创建供应商"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /SP/createSupplier [post]
export const createSupplier = (data) => {
  return service({
    url: '/SP/createSupplier',
    method: 'post',
    data
  })
}

// @Tags Supplier
// @Summary 删除供应商
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Supplier true "删除供应商"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /SP/deleteSupplier [delete]
export const deleteSupplier = (params) => {
  return service({
    url: '/SP/deleteSupplier',
    method: 'delete',
    params
  })
}

// @Tags Supplier
// @Summary 批量删除供应商
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除供应商"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /SP/deleteSupplier [delete]
export const deleteSupplierByIds = (params) => {
  return service({
    url: '/SP/deleteSupplierByIds',
    method: 'delete',
    params
  })
}

// @Tags Supplier
// @Summary 更新供应商
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Supplier true "更新供应商"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /SP/updateSupplier [put]
export const updateSupplier = (data) => {
  return service({
    url: '/SP/updateSupplier',
    method: 'put',
    data
  })
}

// @Tags Supplier
// @Summary 用id查询供应商
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Supplier true "用id查询供应商"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /SP/findSupplier [get]
export const findSupplier = (params) => {
  return service({
    url: '/SP/findSupplier',
    method: 'get',
    params
  })
}

// @Tags Supplier
// @Summary 分页获取供应商列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取供应商列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /SP/getSupplierList [get]
export const getSupplierList = (params) => {
  return service({
    url: '/SP/getSupplierList',
    method: 'get',
    params
  })
}

// @Tags Supplier
// @Summary 不需要鉴权的供应商接口
// @Accept application/json
// @Produce application/json
// @Param data query appReq.SupplierSearch true "分页获取供应商列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /SP/getSupplierPublic [get]
export const getSupplierPublic = () => {
  return service({
    url: '/SP/getSupplierPublic',
    method: 'get',
  })
}
