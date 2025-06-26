package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate()
	if err != nil {
		return err
	}

	// 初始化小程序相关数据表
	MiniProgramTables()

	return nil
}
