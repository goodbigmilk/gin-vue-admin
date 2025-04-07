package app

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ UserRouter }

var USApi = api.ApiGroupApp.AppApiGroup.UserApi
