package app

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ UserApi }

var USService = service.ServiceGroupApp.AppServiceGroup.UserService
