import service from '@/utils/request'
// @Tags User
// @Summary 创建用户信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.User true "创建用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /US/createUser [post]
export const createUser = (data) => {
  return service({
    url: '/US/createUser',
    method: 'post',
    data
  })
}

// @Tags User
// @Summary 删除用户信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.User true "删除用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /US/deleteUser [delete]
export const deleteUser = (params) => {
  return service({
    url: '/US/deleteUser',
    method: 'delete',
    params
  })
}

// @Tags User
// @Summary 批量删除用户信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /US/deleteUser [delete]
export const deleteUserByIds = (params) => {
  return service({
    url: '/US/deleteUserByIds',
    method: 'delete',
    params
  })
}

// @Tags User
// @Summary 更新用户信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.User true "更新用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /US/updateUser [put]
export const updateUser = (data) => {
  return service({
    url: '/US/updateUser',
    method: 'put',
    data
  })
}

// @Tags User
// @Summary 用id查询用户信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.User true "用id查询用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /US/findUser [get]
export const findUser = (params) => {
  return service({
    url: '/US/findUser',
    method: 'get',
    params
  })
}

// @Tags User
// @Summary 分页获取用户信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /US/getUserList [get]
export const getUserList = (params) => {
  return service({
    url: '/US/getUserList',
    method: 'get',
    params
  })
}

// @Tags User
// @Summary 不需要鉴权的用户信息接口
// @Accept application/json
// @Produce application/json
// @Param data query appReq.UserSearch true "分页获取用户信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /US/getUserPublic [get]
export const getUserPublic = () => {
  return service({
    url: '/US/getUserPublic',
    method: 'get',
  })
}
