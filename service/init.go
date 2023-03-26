package service

import (
	"my-admin/global"
	"my-admin/model"
	"os"

	"gorm.io/gorm"
)

func RegisterTable(db *gorm.DB) {
	err := db.AutoMigrate(
		model.Bucket{},
		model.User{},
		model.Photo{},
		model.OperationLog{},
	)
	if err != nil {
		global.Logger.Errorf("failed to register mysql table, err:%+v\n", err)
		os.Exit(1)
	}
	global.Logger.Infof("mysql register successfully")
}
