package dao

import (
	"github.com/amingze/gochive/internal/app/model"
	"github.com/amingze/gochive/internal/pkg/utils/fileutil"
	"github.com/amingze/gochive/internal/pkg/utils/gormutil"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var gdb *gorm.DB

func Init(driver, dsn string) error {
	conf := gormutil.Config{
		Driver: driver,
		DSN:    dsn,
	}
	if driver == "" || driver == "sqlite" || driver == "sqlite3" {
		if err := fileutil.MkFileAll(dsn); err != nil {
			logrus.Error(err)
		}
	}
	db, err := gormutil.New(conf)
	if err != nil {
		return err
	}

	gdb = db.Debug()

	if err := gdb.AutoMigrate(model.Tables()...); err != nil {
		return err
	}

	return nil
}
