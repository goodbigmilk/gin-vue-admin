package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(app.User{}, app.Supplier{}, app.Warehouse{}, app.Product{}, app.Category{}, app.Purchases{})
	if err != nil {
		return err
	}
	return nil
}
